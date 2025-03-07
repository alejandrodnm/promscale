// This file and its contents are licensed under the Apache License 2.0.
// Please see the included NOTICE for copyright information and
// LICENSE for a copy of the license.

package end_to_end_tests

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/timescale/promscale/pkg/internal/testhelpers"
)

type leaseState struct {
	cluster, leader        string
	leaseStart, leaseUntil time.Time
}

func (l leaseState) String() string {
	return fmt.Sprintf(
		"c: %s, l: %s; start: %s; until: %s",
		l.cluster, l.leader, l.leaseStart, l.leaseUntil,
	)
}

func callUpdateLease(db *pgxpool.Pool, cluster, writer string, minT, maxT time.Time) (*leaseState, error) {
	row := db.QueryRow(context.Background(), "SELECT * FROM _prom_catalog.update_lease($1,$2,$3,$4)", cluster, writer, minT, maxT)
	lock := leaseState{}
	if err := row.Scan(&lock.cluster, &lock.leader, &lock.leaseStart, &lock.leaseUntil); err != nil {
		return nil, err
	}
	return &lock, nil
}

func checkLease(db *pgxpool.Pool, lock *leaseState, wantedLockState *leaseState) bool {
	row := db.QueryRow(
		context.Background(),
		"SELECT lease_start, lease_until FROM _prom_catalog.ha_leases WHERE cluster_name = $1",
		wantedLockState.cluster,
	)
	var start, stop time.Time
	if err := row.Scan(&start, &stop); err != nil {
		return false
	}
	if start != wantedLockState.leaseStart || stop != wantedLockState.leaseUntil {
		return false
	}

	return reflect.DeepEqual(lock, wantedLockState)
}

func TestUpdateLease(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	const leaseTime = time.Minute
	const refreshTime = 10 * time.Second
	withDB(t, "ha_check_insert", func(dbOwner *pgxpool.Pool, t testing.TB) {
		db := testhelpers.PgxPoolWithRole(t, "ha_check_insert", "prom_writer")
		defer db.Close()
		// first check
		cluster := "c"
		writer := "w1"
		minT := time.Unix(1, 0)
		maxT := time.Unix(3, 0)
		// w1 becomes leader
		lock, err := callUpdateLease(db, cluster, writer, minT, maxT)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expectedLeaseUntil := maxT.Add(leaseTime)
		if lock == nil {
			t.Fatal("error calling check_insert")
		}
		if !checkLease(db, lock, &leaseState{cluster, writer, minT, expectedLeaseUntil}) {
			t.Fatal("first call to check insert didn't set lock properly")
		}

		// wrong leader
		writer = "w2"
		lock, err = callUpdateLease(db, cluster, writer, minT, maxT)
		leaderHasChanged := "ERROR: LEADER_HAS_CHANGED (SQLSTATE PS010)"
		if lock != nil || err.Error() != leaderHasChanged {
			t.Fatalf("expected leader changed error, got: %v", err)
		}

		// lease_start > min time -> no change to lock
		writer = "w1"
		originalMinT := minT
		minT = time.Unix(0, 0)
		lock, err = callUpdateLease(db, cluster, writer, minT, maxT)
		if lock == nil || err != nil {
			t.Fatalf("expected err: %v", err)
		}
		if !checkLease(db, lock, &leaseState{cluster, writer, originalMinT, expectedLeaseUntil}) {
			t.Fatalf("expected lock details not to change")
		}

		// no update lease_until
		minT = time.Unix(1, 0)
		maxT = maxTime.Add(refreshTime)
		lock, err = callUpdateLease(db, cluster, writer, minT, maxT)
		if lock == nil || err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !checkLease(db, lock, &leaseState{cluster, writer, minT, expectedLeaseUntil}) {
			t.Fatal("expected lock details to not change")
		}

		// update lease_until
		maxT = expectedLeaseUntil.Add(time.Second)
		expectedLeaseUntil = maxT.Add(leaseTime)
		lock, err = callUpdateLease(db, cluster, writer, minT, maxT)
		if lock == nil || err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !checkLease(db, lock, &leaseState{cluster, writer, minT, expectedLeaseUntil}) {
			t.Fatal("expected lock details to change")
		}
	})

}

func TestCheckLeaseMultiCluster(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	const leaseTime = time.Minute
	withDB(t, "ha_check_insert_multicluster", func(dbOwner *pgxpool.Pool, t testing.TB) {
		db := testhelpers.PgxPoolWithRole(t, "ha_check_insert_multicluster", "prom_writer")
		defer db.Close()
		// first check
		cluster1 := "c1"
		cluster2 := "c2"
		writer := "w1"
		minT := time.Unix(1, 0)
		maxT := time.Unix(3, 0)
		lock, err := callUpdateLease(db, cluster1, writer, minT, maxT)
		expectedLeaseUntil := time.Unix(3, 0).Add(leaseTime)
		if lock == nil {
			t.Fatalf("error calling check_insert for first cluster: %v", err)
		}
		if !checkLease(db, lock, &leaseState{cluster1, writer, minT, expectedLeaseUntil}) {
			t.Fatal("call to check insert didn't set lock properly")
		}

		// same writer instance different cluster
		writer = "w2"
		lock, err = callUpdateLease(db, cluster2, writer, minT, maxT)
		if lock == nil {
			t.Fatalf("error calling check_insert for second cluster: %v", err)
		}
		if !checkLease(db, lock, &leaseState{cluster2, writer, minT, expectedLeaseUntil}) {
			t.Fatal("call to check insert didn't set lock properly")
		}
	})
}
func TestConcurrentUpdateLease(t *testing.T) {
	// 3 clusters with 3 writer instances
	// each writer calls check_insert from a separate routine
	// per cluster 1 routine will become leader, the others will
	// always get a LEADER_HAS_CHANGED
	clusters := []string{"c1", "c2", "c3"}
	writers := []string{"w1", "w2", "w3"}
	numCallsToCheck := 10
	timeBetweenCheck := 10 * time.Millisecond

	withDB(t, "ha_check_insert_concurrent", func(dbOwner *pgxpool.Pool, t testing.TB) {
		db := testhelpers.PgxPoolWithRole(t, "ha_check_insert_concurrent", "prom_writer")
		defer db.Close()
		wg := sync.WaitGroup{}
		wg.Add(len(clusters) * len(writers))
		firstLockCounters := make(map[string]*int32)
		firstLockCountersArr := make([]int32, len(clusters))
		for i, c := range clusters {
			firstLockCountersArr[i] = 0
			firstLockCounters[c] = &(firstLockCountersArr[i])
			for _, w := range writers {
				go func(cluster, writer string, counterForCluster *int32) {
					minT := time.Unix(0, 0)
					maxT := time.Unix(2, 0)
					firstLockRes, firstErr := callUpdateLease(db, cluster, writer, minT, maxT)
					if firstLockRes != nil {
						atomic.AddInt32(counterForCluster, 1)
					}
					for i := 0; i < numCallsToCheck; i++ {
						time.Sleep(timeBetweenCheck)
						minT = minT.Add(time.Second)
						maxT = maxT.Add(time.Second)
						lockRes, err := callUpdateLease(db, cluster, writer, minT, maxT)
						// this routine got an error on first try -> other routine is leader
						if firstErr != nil && err == nil {
							wg.Done()
							t.Fatal("non-leader routine acquired lock")
							return
						} else if lockRes == nil && firstLockRes != nil {
							wg.Done()
							t.Fatal("leader failed to acquire lock")
							return
						}
					}
					wg.Done()
				}(c, w, firstLockCounters[c])
			}
		}
		wg.Wait()
		for cluster, numTimesFirstLockWasNotNil := range firstLockCounters {
			if *numTimesFirstLockWasNotNil != 1 {
				t.Fatalf(
					"Expected first lease to be set only once per cluster. It was set [%d] times for cluster [%s]",
					*numTimesFirstLockWasNotNil,
					cluster,
				)
			}
		}
	})

}
