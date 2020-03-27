// Code generated by vfsgen; DO NOT EDIT.

package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// SqlFiles statically implements the virtual filesystem provided to vfsgen.
var SqlFiles = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 3, 26, 16, 4, 17, 604080077, time.UTC),
		},
		"/1_base_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "1_base_schema.down.sql",
			modTime:          time.Date(2020, 3, 27, 12, 47, 12, 395328333, time.UTC),
			uncompressedSize: 88,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x76\xf6\x70\xf5\x75\x54\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x86\x0a\xc4\x3b\x3b\x86\x38\xfa\xf8\xbb\x2b\x38\x3b\x06\x3b\x3b\xba\xb8\x5a\x73\xe1\x55\x1d\x10\xe4\xef\x0b\x57\x0a\x08\x00\x00\xff\xff\xac\xa9\x98\xf1\x58\x00\x00\x00"),
		},
		"/1_base_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "1_base_schema.up.sql",
			modTime:          time.Date(2020, 3, 27, 13, 2, 41, 586671425, time.UTC),
			uncompressedSize: 20080,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\xfb\x6f\xe3\x38\x92\xff\xef\xfe\x2b\x0a\x8b\xcc\xda\x9a\xb1\x3c\x9d\x9d\xef\x2e\xbe\x97\x4c\x1a\x70\xdb\x4a\x5a\xb7\x8e\x9d\xb5\x95\xe9\x9e\x1b\x0c\x04\x5a\xa2\x6d\x4e\x64\x49\x2d\xd2\x49\x7c\xd8\x3f\xfe\xc0\x22\x25\x51\x0f\x3b\x4e\x3f\xb0\x77\xc0\xfa\x87\xee\x58\x0f\xb2\x1e\x9f\x2a\x56\x15\x8b\xb6\xed\xe9\xcc\x73\x16\x1d\xdb\xf6\x36\x8c\x43\x90\x84\x14\x08\xe7\xbb\x2d\xe5\x20\x36\x44\x80\x20\xcb\x88\x42\x4c\xe4\x85\x80\xc4\x90\xc4\xd1\x1e\x96\x14\xfe\xf6\x13\x04\x1b\x92\x71\x88\x92\x78\xdd\xe9\x8c\xe6\xce\xd0\x73\x60\x31\x7a\xef\xdc\x0e\xc1\xbd\x86\xe9\xcc\x03\xe7\xa3\xbb\xf0\x16\xfa\xa2\x3f\x1a\x7a\xc3\xc9\xec\xe6\x12\x6c\x1b\x02\x22\x48\x94\xac\xd5\xe8\x1c\x7e\x00\x16\x0b\x9a\xc5\x24\x82\xd5\x2e\x0e\x04\x4b\x62\x7e\xca\x90\x77\xf3\xd9\x2d\x8e\x17\x12\x41\xca\xc1\xd2\xdd\x32\x62\x81\x31\x54\x3e\x96\xf3\xd1\x73\xa6\x0b\x77\x36\xad\x0d\x27\xd8\x96\xf2\x80\x44\x34\x5c\xc2\x07\xd7\x7b\x9f\x4f\xaa\x06\xba\xec\xd8\xed\x9f\x8e\x6d\x83\x87\xf2\x09\xe9\x8a\xc5\x0c\x27\x03\xbc\xde\xfe\x7c\x4e\x87\x37\x7c\x37\x71\x6a\x72\x19\x70\x9a\x31\xca\xa1\xd7\x01\x00\x60\x21\x2c\xd9\x5a\x5e\x22\x11\xdc\xcd\xdd\xdb\xe1\xfc\x57\xf8\xbb\xf3\x6b\x1f\xef\x6e\xa9\xc8\x58\xe0\xb3\x50\xca\x4d\x5d\x8a\xc8\x92\x46\x5c\x7e\xff\xed\x77\x75\xe5\x7e\xea\xfe\xe3\xde\xe9\xa9\x1b\x16\xb8\xd3\xd1\xe4\x7e\xec\x40\x8f\x85\x56\xc7\xba\xcc\x49\x71\xa7\x63\xe7\x23\xa8\xb9\x7d\xf5\xac\x1c\x77\x36\x3d\x40\xde\xfd\xc2\x9d\xde\xc0\x8d\x3b\x85\x7c\xe4\xcb\xe3\x6c\xe1\x53\x25\x57\x8a\x25\x45\xe1\x03\xdd\x83\xe7\x7c\xf4\xd4\xb7\x47\x12\xed\x28\x08\xfa\xac\x39\x32\xb8\x46\xa2\x4b\x0e\x1e\xe8\xbe\xaf\x1e\xb7\x4c\x56\x2b\x37\x1a\xfc\xbe\x4c\xa4\xff\x40\xf7\x7e\x9a\x70\xd4\xa3\xa6\x58\x49\xda\xa0\xaa\x4a\x73\x9a\xf0\x52\x05\x39\x15\xea\x9d\xbe\x7c\xd4\x20\x23\x4d\x78\x53\xee\x4d\x31\x37\x29\x39\x32\xde\x0b\x5c\x69\xe2\x0b\xd9\x2f\x9c\xb9\x3b\x9c\x1c\x84\x93\x34\x72\x83\x53\x34\x27\x75\x51\xfe\xa3\x2e\x86\x74\x45\x76\x91\xf0\x83\xcd\x2e\x7e\xf0\xd1\x6a\x1f\x49\x04\xef\x66\xb3\x89\x33\x9c\xc2\xd8\xb9\x1e\xde\x4f\x3c\x10\xd9\x8e\xb6\x09\x05\x87\x33\x98\x28\xe7\xa8\x68\xd2\xbc\xfe\xa2\xee\x34\x49\x9a\xcd\x5c\x3f\x4d\x2e\x15\xbe\xe4\x3d\x1c\xd2\x9d\x2e\x9c\xb9\x07\xee\xd4\x9b\x1d\x18\x11\xe1\xa4\xd1\xf4\xcb\x70\x72\xef\x2c\xa0\xd7\xad\x32\xde\xed\x43\xcf\x9d\x7a\xce\xfc\x97\xe1\x04\xba\xff\x1f\x36\xc9\x2e\xe3\x5d\xeb\xe2\x42\x8a\xd1\x20\x7c\x36\x87\xb9\x73\x37\x19\x8e\x1c\xb8\xbe\x9f\x8e\x3c\xb7\xa9\xf8\x35\x15\x7e\xbb\x74\x7b\x16\xd2\x3f\x77\xbc\xfb\xf9\x74\x01\xf9\x7c\x9d\xe1\x02\xce\xa4\x8f\x3b\xc3\xdb\x0b\x67\xe2\x8c\x3c\xc5\xe5\xc5\x45\x41\xd4\xf5\x7c\x76\x7b\x48\x62\x1f\xde\x3b\x73\x47\x4a\xec\xaa\xce\xd6\x65\x47\x8f\x3c\x19\x4e\x6f\xee\x87\x37\x0e\xf0\x4f\x11\x2c\x94\xfc\xef\x86\xf3\xe1\x64\xe2\x4c\x60\x31\xbc\x76\x2e\x3b\xaf\x61\x72\x4b\x1e\xa8\xaf\xa1\x80\x3a\xae\xf1\x26\x32\xb6\x5e\xd3\x0c\xaf\x95\xec\x8d\x9d\xd1\x64\x38\x77\x3a\xef\x9c\x1b\x77\x2a\xef\x39\x1f\x9d\xd1\xbd\xe7\xc0\x2a\xc9\xb6\x44\xf4\xba\x6d\xf0\x90\x6b\xc3\xe0\x3b\xb7\x27\x7d\x3b\x78\xee\xad\xb3\xf0\x86\xb7\x77\xde\x7f\x69\x07\x01\xe3\xd9\x3d\x32\x33\x77\x46\xae\x5c\x12\xfa\xb9\x13\x64\xa1\x14\xb1\xd5\x55\xa8\xa9\x7f\xa6\xce\x87\x81\x81\xce\xcb\x23\xe4\xd4\x4d\x5c\x53\x04\xbd\x62\xa2\x3e\xae\x3c\x86\x3d\x28\xb4\xbd\x62\xea\x3b\x67\x7e\x3d\x9b\xdf\x42\x90\x51\x22\xa8\xbf\xd9\xa7\x34\x53\x82\xcd\x89\xa9\x4e\xde\xed\xd7\x87\xe9\x43\x57\x12\x71\x60\xce\xe2\xa3\x00\x22\x9f\x2c\x50\x72\xf5\xf6\x15\x10\x6e\x95\xd4\x99\x82\x2e\x0c\x27\x9e\x33\x6f\x55\x1f\x2c\x1c\x4f\x5b\x36\xfa\xa4\x72\xa1\x1e\x04\xc9\x36\xcd\x28\xe7\xfd\xa3\x77\x7d\x4e\xd7\x5b\x1a\x8b\xe5\x1e\xae\xa0\x5b\x48\xbe\xfb\xc2\x5b\x49\x16\xd2\x4c\xbd\x83\xd2\xc1\xa7\xad\x4b\x38\x3b\x6b\x08\xf0\xb2\x23\x6f\xda\x36\x72\xcc\xe1\x69\x43\x33\x0a\x62\x43\x81\xc6\x21\x0e\x0e\x8c\xc3\x92\xae\x92\x8c\x42\x9c\x3c\xf5\x2c\xfb\xfc\x0d\x6c\x59\xbc\x13\x94\xc3\x13\x8b\x22\x19\x4e\xe5\x13\xd3\xd0\xd4\x2a\x09\x43\xbf\x20\x49\x8d\xef\xa7\x49\xc4\x82\xfd\x2b\xd4\x5b\xfa\xa7\x72\xde\xae\x52\x87\x32\x3b\xf9\xca\x65\xc7\x99\x8e\x1b\x46\x9f\x46\xe9\x9a\x7f\x8a\x0c\x07\x3c\x77\x6f\x6e\x9c\x39\x34\xec\xd8\xaf\x58\xee\xb5\x54\xa7\x76\xb0\x2d\x4e\x00\xdf\xc3\x27\xaf\x67\x73\x70\x86\xa3\xf7\x30\x9f\x7d\xc0\x0b\x39\x40\xee\xe6\xb3\x91\x33\xbe\x9f\x37\xd7\xb4\xa6\x07\x91\x2e\xa8\x73\x28\xe6\xc2\x20\xcd\x6d\x44\x97\x47\xa2\x34\x5b\x8e\x05\x73\x2a\x76\x59\x0c\xc4\x08\x80\x61\xb9\x63\x91\x80\x55\x96\x6c\x81\x54\x16\x4d\x12\x87\x40\x80\xef\x56\x2b\xf6\x3c\xc0\xa8\x70\x43\xf3\xd0\x01\x1f\x60\x5c\x2e\x89\x71\x40\x04\x0d\x81\x27\x3a\xb4\xde\x50\xfd\x0e\x04\xc9\x2e\x0a\x61\xc5\x04\xb0\x18\x56\xbb\x28\x1a\xbc\xc6\xa9\x9a\x7a\x90\xd3\xf9\x4f\x4c\x6c\x7c\x35\x74\x69\x3b\x8d\x55\x3e\x9f\x1c\xd7\xaa\x8a\x1b\x96\xcf\xb4\x2f\x2f\x3d\xbe\x5b\x72\x91\xb1\x78\x6d\xae\xe8\xd2\x9e\xe1\x6f\x3f\xd9\x3d\x99\x12\xf8\x11\x8d\xd7\x62\xd3\x53\xa3\x5b\x3f\x9c\x5b\x16\xfc\xf3\x9f\xd0\xf5\xbb\xf2\x3f\x7d\xf5\xe2\x02\xe7\x68\x5b\x65\xdc\xdb\xdb\xfb\xf6\x85\xc6\x54\x4b\x4c\x9f\xa0\x98\x9a\x14\x71\x9a\x94\x81\xd6\x00\xca\x5c\x86\xac\x22\x81\x1d\x57\x26\x69\x3e\x56\x88\x1a\xde\xed\x04\xb0\x95\x7c\x40\xbe\x69\xea\x2d\x4c\x28\x8f\xbb\x42\x6a\xa6\x0f\x6b\x1a\xd3\x8c\x48\xab\x55\xd3\xef\x62\xf6\x69\xa7\xa0\x81\x53\x4e\x13\x41\x73\xcd\x32\x95\x2d\xc9\x79\x77\x29\x4e\x1d\xd3\x67\x21\x97\x1e\x48\x56\x07\x14\x88\xaa\x93\xf2\x79\x96\xa3\xf1\x04\x98\x00\xbe\x41\x64\x48\x0f\x41\xa2\x88\x86\x2a\xff\x62\xab\x02\x98\x92\x42\x88\x13\x01\x7b\x2a\x80\x3e\x33\x2e\x5e\x03\x9d\x98\x3e\xf9\x0d\xf8\xb4\x42\xc6\x27\xd9\x5a\xc3\xa6\x92\x7d\x9c\x0a\x9c\xd1\x70\xe1\x14\xe3\x7e\x78\xef\x4c\xc1\xc4\x4a\x6d\x22\x0b\x7e\x96\x09\xa6\xf7\xde\x99\x56\x56\xa5\xda\x63\x1a\x44\xf9\x5d\x67\x62\x4c\x81\x53\x7f\x96\x9d\x1c\x98\xcc\xe0\x5b\x85\x77\x95\x17\xac\x92\x8a\x16\x3f\x2a\x61\xfd\xcb\x6c\x32\xf4\xdc\x56\x54\x8f\x70\xf1\xce\x61\xa5\xf4\xaa\x60\xbd\x66\x8f\x34\x36\x11\x39\xc8\x33\xf5\x1d\xa7\x5c\x42\x8b\x27\x5b\x0a\x9c\x7e\xda\xd1\x38\xa0\x5c\xa2\x46\x43\x26\x4f\xd4\x15\x6e\x3a\xb6\xed\x22\xc2\xbf\x0a\x6c\x74\xb4\x51\x71\xc3\x2f\x80\x66\x76\xef\x81\xce\x57\xf1\xef\x5a\x76\x61\x75\x9a\xe1\x1e\x48\x71\x68\x98\x5d\xea\xe0\x4f\x5f\xb9\x42\x73\x92\x81\x45\xba\xf6\x65\xd8\xa1\xd2\x4a\x3f\x17\x44\xb1\x2e\x56\x75\xdf\xed\x77\x59\xd8\xb5\xac\x8b\x0b\x1c\x72\x32\x9b\xdd\x21\xd9\x47\x72\x81\x3c\x87\x92\xc1\x9a\xc1\x59\x1f\xcc\x14\xa5\x00\x9c\xc2\xba\xa6\xbb\x19\x51\xd5\x21\xd5\x78\xe0\x24\xf3\x6c\x00\x53\x4d\xa7\xc8\x98\x4d\x61\x34\x9b\x5e\x4f\xdc\x91\x07\xe3\x19\x4c\x67\xde\x7b\x77\x7a\x63\x18\xa9\x4c\xdf\x5b\x79\x1c\x48\x16\xdb\xef\x94\x93\x6b\x71\x79\x33\xc0\xe8\xb5\xb8\x8e\xe1\x04\xd8\x36\xec\xe2\x90\x66\xb0\x61\xeb\x0d\x04\x49\x1c\xec\xb2\x8c\xc6\xc1\x1e\x91\xc7\x62\x4e\x33\x01\x5b\xb2\x47\xe4\x65\xda\x95\xc7\x7b\xb1\x61\xf1\xba\x8f\xeb\x62\xb6\x97\x0b\x29\x8d\x68\x20\x70\x55\x8d\x92\x24\xcd\x87\xde\x08\x91\xf2\x8b\x1f\x7f\xe4\x82\x04\x0f\xc9\x23\xcd\x56\x51\xf2\x24\x43\xb6\x1f\xc9\x8f\xe7\x7f\xfd\x8f\xbf\xbe\xf9\xe9\x2f\xff\x4f\x07\x11\xae\xa7\x7c\xcc\xf5\xec\x7e\x3a\x56\x01\x5a\xae\x9c\x2d\xf2\xb9\x3d\x81\x27\x15\xa1\xb4\xe4\x4f\x1a\x13\xdb\x8e\xf6\x65\x73\xa7\xb2\xb8\x5e\xd5\xf5\xac\x09\x68\x90\xe5\x4c\xc7\x20\x31\xd8\x1e\x7c\xdd\x4d\xee\x6e\x16\xff\x98\x1c\x73\x1c\x70\x43\x85\xf6\x1a\xaa\xcc\x42\xb2\x8c\xec\xa1\xa8\x61\x28\x27\xa2\x6e\x3d\xd0\xfd\x00\xae\xe5\x85\x78\xaf\x09\xec\xcb\x21\x9e\x28\x3c\x91\x58\x45\x23\xf9\x8b\xb8\x66\x2e\x29\x10\x8e\x81\x29\x91\xca\xe0\xf2\x2e\x67\x95\xf5\x15\x9d\x10\x7a\xa0\x34\xa3\x42\xec\x61\x43\xc9\xe3\x1e\xa2\x24\x78\x40\x57\x24\x97\x3f\x9e\x12\x19\x2e\x44\xfb\xd7\x38\x18\x69\xd2\x12\xd5\x69\xc2\xfd\x55\x92\xf9\x0f\x74\x7f\x2c\x90\x79\xa0\xfb\xf2\x6b\x75\x45\x62\xb1\x68\x75\x2d\x50\x4a\x09\x7d\x81\xbc\x22\x7d\x8a\xdf\xbc\x6c\x5a\x20\x4c\x87\xb7\xce\x65\x99\x89\x82\x6d\x4b\x26\xc3\x64\x27\x6f\x06\x1b\x1a\x3c\x20\xfb\x2c\x5e\x83\x4c\x04\xf4\x33\x2b\xc6\x05\x24\xa9\x60\x5b\xc6\x05\x0b\xd4\x83\x17\x06\x2c\x0b\xe6\xd2\x84\x17\xb8\xeb\x1c\x70\x0b\xcd\x1a\x51\x89\xc4\x9a\x94\xaa\x60\x2c\x6e\x0e\xa7\x63\xac\xab\x5d\x15\xa2\x2b\x8d\x20\x1f\x53\xa3\xd6\xbd\x56\x70\xad\xae\xc6\x3a\x73\x28\x9f\xd5\x4b\x1f\xb8\xd7\x55\x73\x3b\x62\x4e\x98\xb3\x48\x4d\x27\x99\xdf\xb2\xa2\x34\x9c\x9d\x55\xd2\x68\x3e\x97\xbb\x1e\x29\x76\x09\x53\x09\x7a\x33\xa1\x5c\xd2\x80\x48\x1d\x3d\x51\x20\x19\x46\x7d\x74\xb5\x92\xfe\x25\xd8\x90\x78\x2d\x15\x85\x81\x78\xb0\xa1\x5b\x62\xea\x8c\x44\x3c\xc1\x44\x87\x03\xdf\xe9\x94\xae\x8a\x90\x25\x8d\x92\x27\xc0\xb2\x79\x96\xc9\x11\x59\x0c\x82\x66\x5b\x2e\x63\x3c\xc3\xfb\x55\x32\x9b\x3c\x71\x9b\xcc\x46\x7f\x6f\x4f\x79\xdd\x29\x2c\xde\x0f\xe7\x0e\xdc\xdf\x8d\x55\xd5\x7a\x34\xb9\x5f\xb8\xbf\x38\x70\x3b\x1b\x3b\xdd\x7e\x85\x7b\x2b\x67\x9f\xd3\x20\x89\x43\x0d\x41\xb2\x12\x34\x43\x20\xfe\x6f\xc2\xd8\x37\xc0\x57\x49\x0a\x79\xc6\x92\x28\xfc\x00\xe7\xdf\x8e\x39\xc5\x41\x05\x05\x25\x1b\x55\x70\xb8\x0b\x98\xde\x4f\x26\x55\xb6\xaa\x8f\x5c\x5c\xc1\x39\x6e\x5f\x9c\xdb\x2c\x0e\xe9\x33\x0d\x95\x03\xe7\x4d\x6e\x8f\x44\x28\x07\x98\x91\x9f\xbc\x6a\x59\x09\x5c\x72\x6d\xf4\xab\xd4\x34\xc2\x87\x62\x94\x43\x61\x84\x89\xa7\x43\xaa\x9d\xce\xbc\x56\xf5\x0e\xdd\x85\x03\xdd\x11\x46\xaa\x32\x18\x58\x31\x4c\xa1\xe5\x4a\x96\x0f\xd2\xad\x2a\xbd\x15\x15\xc7\x8b\x16\xb6\xdd\x1a\x0b\x6b\xeb\x20\xda\x56\x74\xc6\xa6\x13\x72\xb5\x58\xe6\x79\x1f\x46\xc6\x9f\xb1\x6e\x29\x9d\xb0\xb0\x57\x59\x99\x74\xe9\xd1\xbc\xa0\xc3\x62\x77\xea\x99\x11\xb0\x5a\x5e\xda\xc2\xd3\xa3\x88\x36\xf7\x3f\x3a\xa5\xfe\x8b\x77\x0a\x6a\xfa\x25\x1d\x9f\x1f\x35\xe2\x94\x03\x16\x9a\xe1\xd3\xa1\x48\xa7\xcd\x62\xeb\xef\x1e\x0c\xb7\x94\x4a\xa2\x16\x13\x7d\xa0\x7b\xd3\xbd\x0c\xa7\xe3\xe2\x96\x2a\xf2\x5e\x19\x12\xff\xa2\x28\x0c\xd1\xf4\x94\x91\x34\x95\xc8\xc9\x92\x5d\x1c\xc2\x1f\x3c\x89\x97\x3e\x25\xc1\xc6\x97\xca\x94\x31\x93\xcc\xd7\x80\xc0\x92\x0a\x89\xb0\x2c\x79\xf2\x29\x17\x6c\x4b\x04\xed\xd8\xb6\x5c\x98\xf4\x36\x5d\xef\xfc\x0d\xc2\xfe\xfc\xcd\x1b\xeb\x15\xf0\x52\xb0\xaa\xcd\xdb\xfb\x83\x2b\x52\xfa\x80\x70\x92\x42\x29\xc1\x55\xee\xaa\x59\x9d\x22\x28\x5a\x38\xde\xec\x1a\x32\x1a\x24\x59\xd8\x81\x82\xd7\x7c\x13\xb6\x73\xa8\x06\x03\x0b\x6f\x2e\x21\x32\x9f\x7d\x58\xc0\xf9\x9b\x02\xb1\xd2\x18\xcf\x6a\x64\x95\x37\xda\x64\xb7\x8b\x63\xca\x4b\x91\x95\x02\x83\x5c\x60\x5f\x26\x23\x35\xbe\xda\xa2\xf4\x55\x54\x4c\xe2\x3d\xfe\xd1\x90\x03\x89\xf7\x34\xa2\x5b\x1a\x8b\xaf\x26\x0b\x9c\x48\x13\x51\x11\xc4\xc1\xa2\xe3\x91\x4f\xdb\x3b\x70\xa7\xf6\xb8\x87\x77\x2e\x87\x13\xdf\x79\x71\x9e\x4e\x39\x6e\x4e\x33\x6a\x88\x2a\xd7\x88\x36\x96\x98\x55\x84\x66\x91\xa2\x48\x0d\xb0\xb0\xad\xa2\xba\x96\xca\xdb\x0a\x98\xf8\xcc\x1a\xc4\x49\x51\xe3\xc1\x5c\xa1\x9a\x1c\x28\xb5\xf6\xf2\xba\xc4\x91\x9a\x44\x19\xd0\x56\x73\xc4\xb2\x04\xf5\x52\xa2\xa8\xf3\xc4\x41\x35\x53\x7c\x81\x91\x41\x2d\xb0\xba\x9f\x4a\x49\x0c\x27\x13\x83\xa0\xef\x0f\xcd\xde\x56\xa5\x79\xc5\x7c\x28\xaa\x89\x7b\xeb\x7a\x70\xfe\xda\x8a\x56\x7a\x18\x45\xc7\x53\xd4\xd7\x00\xa0\x12\xf8\x7c\x66\x7a\xe8\x4e\x3d\xd4\xf2\x99\x8e\xa2\x31\x52\xa0\xcf\x34\xc0\xdd\x19\x04\x6e\x92\x51\xa0\xcf\x29\x8d\xb9\xf4\x54\x79\x56\x5e\xb0\xa6\x0a\xc7\xad\x71\x43\xcb\xa2\xf7\x2f\x4d\xef\x2a\xe8\xa9\x53\x76\x42\x06\xde\x1a\x46\x2a\x79\x16\x38\x31\x30\x92\x2f\x9b\x37\xd4\x08\xad\x7c\x16\x6a\x9d\x97\xc1\x0a\xa4\x84\x65\xaf\xd7\x3c\x0b\x7b\x66\x24\x70\x24\xce\x3a\xae\xf3\x15\xcb\xb8\xc8\x8b\x4e\x22\x81\x34\xa3\x8f\x34\x16\x45\x55\x55\xed\xf8\x2c\xa9\x4c\x15\x77\x9c\x86\xb0\xcb\x4b\x52\xd2\xc1\x07\x94\x73\x92\xb1\x68\xdf\x26\xd4\x97\xa2\x9a\x2f\x8d\x69\x3e\x5b\xad\x8d\x00\xd5\x94\xd9\xcb\x2a\x45\x17\x6f\xee\x87\x94\x39\x31\xe1\x79\x1a\xa3\xe4\x26\x35\x8f\xa1\x41\xc7\xb6\xdf\x70\xc8\x68\x9a\x51\x2e\xc5\xfb\x40\xf7\xba\xc7\x8c\xe0\xce\xa8\x14\xb8\x80\xde\x13\x85\x30\x91\x36\x24\x73\x76\x99\x41\xc9\x30\x9e\x49\x35\xb0\x58\xa8\x71\x8b\x85\x83\xef\xd2\x34\xc9\x04\x30\x61\x15\x85\x71\x56\xdc\xa2\x19\xa4\x34\xc3\x7c\x5b\xbe\x1e\x64\x4c\xb0\x80\x44\xa0\x46\x13\x1b\xc6\x3b\xb6\xcd\xb8\xca\x0e\x50\xb1\xd2\x53\x95\x85\xca\x20\x62\x92\x4e\x12\x87\x98\xe3\x93\x60\x43\x43\x79\x3f\xa3\x27\xaf\x53\x2a\x26\x12\x89\x6f\x04\x22\x45\xbc\x66\x75\x0c\x44\xfe\xf6\x3b\x94\x98\xc4\x1e\x34\x16\x3e\xfb\x8f\x24\x92\x97\x7b\xb5\x0a\x73\xa5\x6e\x6c\xdb\x8a\x03\x99\xdc\x20\xf9\xe5\x26\xa7\x48\xf2\x25\x58\x26\x56\x12\x59\x45\x85\xaf\x3e\x04\x16\x5b\xd1\x87\xb1\x90\x6b\xa7\xb6\xd7\x9a\x40\x6f\x06\x3d\x69\xb6\xa6\x40\x33\x4a\x78\x12\x73\xab\x32\x54\x90\x90\x88\xf2\x80\xf6\xa2\x87\x74\x90\x26\xbc\x5e\xe3\x3e\xee\xc4\xff\xe0\xf6\xdb\xb7\x5d\x5f\x55\x4e\xfd\x6e\x1f\xe8\xe0\x81\xee\x2d\x4b\x0a\xa3\x7f\x60\x9e\x41\xb3\xd2\x7e\xd0\x5f\xe0\x70\x72\x54\x95\x24\x59\x12\xf4\xc5\xbb\x07\xad\xb4\x25\xe2\xb6\x80\x56\xe7\x9c\x38\xd7\x1e\xfc\xe7\xcc\x6d\x8f\x47\x21\xaa\x51\x28\x33\xae\x5e\x34\x50\xc6\x8e\x54\xa1\xd3\x8e\x06\xb9\x8d\xe7\x24\x76\x4e\x9f\xa4\xda\xbf\x16\x3d\xa4\xcd\x39\xeb\x57\x9a\xbb\x5e\x20\x5f\x1c\x14\xab\x4b\x4d\x21\x15\x77\x54\x7d\xc5\x60\xa5\xfe\x44\xc9\x84\x6d\xc7\x94\x86\x08\x4c\xec\xb0\x80\xe5\x5e\xe5\x2b\xa5\xd7\x0d\x29\x09\x55\x01\x99\xad\xc0\x54\x1e\x1a\xa1\x44\xb3\xf4\xc3\x2a\x8f\x2a\xc6\x9d\xcd\xc7\xce\x1c\xde\xfd\x0a\x51\x31\xbf\x65\x56\x21\x87\xf3\xf9\xf0\xd7\xba\x15\x95\x18\xd2\xa6\x26\x45\xde\x87\x37\x56\x05\x11\x15\x66\x72\x97\xe7\xab\x46\x92\x36\xf1\x01\x9c\xb7\xb7\xd2\xf4\xf2\x0d\x08\xf2\x2c\x27\xb4\x14\xde\xf4\xd4\x55\x3d\x5b\xb0\x3e\xa0\xf7\xdc\x29\x48\xf8\xe4\x54\xb3\xf0\x59\x46\x92\x96\x66\xbb\xe6\xaf\x0f\x87\x69\x27\xfa\x30\x89\x2a\xb5\x36\xa8\xac\xa6\xe6\xcd\xcc\x70\x0b\x1b\x23\xa1\x00\x23\xc7\xc5\xf7\xb7\xdf\xf3\x4b\x38\x4a\x7e\xf1\xdf\xde\xaf\xe1\xfd\xaa\xf1\xd5\xe3\xd7\x75\x7d\x6a\x3c\x1c\xf7\xa0\xf3\xc3\x5c\x56\xfe\xd5\xab\xe4\xcf\x52\x95\x56\x1f\xee\xa7\x53\x67\xe1\xf5\x4c\x5d\x5a\x96\x54\xd0\xc3\x63\xa3\xf2\xd4\x44\xee\xeb\xdd\xa2\xa2\xb8\xe6\x17\x0b\xf2\xff\xc5\x8e\xd1\x84\xfd\x4b\x4e\x51\x31\x72\xd8\x2b\x16\xde\xcb\x78\xf0\xdf\xee\xeb\x25\xf7\x65\xdb\xaa\x19\x88\x97\x21\xa6\xce\x2c\x74\xab\x3a\x36\xe5\xd3\x50\xae\x31\x3a\xc3\x54\x2e\xab\x63\xdb\xb5\x56\x9d\x22\x99\x2b\xdb\x6d\xd4\x66\xdd\x7f\x53\x95\xad\x18\x86\x7f\xaa\xdb\x34\x26\x94\x2e\x13\xe9\xeb\x99\x5d\xfa\xa5\x03\x54\xb4\x97\x0e\xb0\xe6\xe6\x54\xf4\x91\x2c\xff\xa0\x81\xe8\xa9\x01\xc9\x7a\xad\xcc\xc4\xea\x83\x79\x45\x5b\x76\x3d\xc5\x7c\xb1\x44\xc6\xad\x22\x3b\xd3\xaf\xb8\xd3\xa9\x33\x3f\x66\xb9\xda\x54\xb1\x19\x23\x7f\xb7\xa9\xbd\x03\x2d\xc3\xb6\x3d\x4e\x30\x4c\x47\xb7\xad\x9b\x23\xb1\x1c\xaf\xf6\xe3\xf2\xee\x38\xad\xc8\x46\xa1\xe8\xf5\xbd\x2a\x75\xc4\xb7\x1c\xa1\x80\xfc\x18\x85\x2e\x18\x1a\x67\x29\xd0\x4a\xef\x3d\xa3\x45\xf8\x9d\x7b\x73\x72\xd1\xbe\xfd\x28\x45\xaf\x20\x41\xaf\x8f\xbc\x62\xf3\xf5\xbb\x1a\xbb\x70\x6a\xb5\xde\x2c\xae\x17\x64\x9f\x50\x9f\x6f\x7f\xf1\x60\x1a\xab\x9e\x28\xf3\xd8\xbc\xc2\x7d\x65\x52\xfd\x8d\x1a\x20\x9a\x20\x38\x5c\x32\x28\x38\xc1\x6a\x86\xda\x2e\x51\x48\xae\xe5\x62\x4a\xb1\xa5\x35\x62\x34\x32\xf2\x9c\x4a\x24\xa2\x45\xf5\x62\xaa\x87\x7f\xa3\x5a\xad\x4e\x5d\xc2\xc7\xa5\x59\x13\x66\xee\x7d\xbf\x57\xaf\x05\x42\x2d\x80\x6d\xf5\xc0\xa3\xf8\xef\xbd\xa6\x8a\x8a\xb3\x57\xe2\x7f\xcb\xc2\xee\x99\x16\x6a\xaa\xd5\xc2\xcf\x0f\x41\x4f\xd5\x60\x2d\x2a\xfd\x0a\x61\xe8\xe7\x2a\xfe\xf4\xf8\xd8\xa4\xa9\x4a\xcc\xff\x59\x88\x98\xf5\xe2\xaf\x8c\x0d\xd5\xba\x7d\x47\x32\xb2\xa5\x82\x66\xb0\x25\x31\x4b\x77\x11\x51\x65\xe3\xe2\x6c\x9f\x71\xac\xee\x84\xd5\x80\xd3\xfa\xd1\x04\x3f\x89\xab\x25\xf2\x26\x92\xb0\xbf\x2e\x3f\xea\x94\xb7\xd3\x97\xc0\x79\x4c\x58\xd8\xde\xf4\x5a\xce\x56\x39\x38\x51\xae\x42\x87\xba\xf8\x7b\x8d\xee\x99\x2f\xeb\x9c\xb1\xac\x8b\x8b\x8c\xae\x83\x88\x98\xe7\x25\x4c\xbe\xac\xe6\x91\x9f\xaf\x91\x3c\xf2\xc3\x47\x42\x6a\x47\xc8\x9a\x72\xd5\x87\xca\x8c\x5a\xed\x89\x67\xb6\xa0\x72\x68\xab\xe5\xcc\x56\xf5\x82\x3e\xad\xd5\x58\x5a\x7b\x78\xe0\x6e\x3c\xcb\xfb\x70\x16\x8e\x57\x94\x60\xb1\x27\x67\xec\x8c\x55\xc4\x55\x5d\x40\xbf\x08\x72\x75\xe2\xac\x83\x2b\xaf\x71\x8c\x42\x79\x82\x76\x39\xd7\xda\xb1\x32\x49\xed\xd9\xa9\x3b\x39\xaf\xd0\xb2\xe6\xa1\xa6\xe4\xa6\x31\xbd\x4a\xed\xa5\x45\xe5\x3a\x68\xdd\x60\x93\x9a\x39\x70\x32\xf1\x0a\x56\x24\xe2\xd4\x10\x13\xb6\x5b\x14\x6e\x8a\x85\xaf\xb7\xac\x83\xec\x56\xb6\xcf\xac\x6f\x89\x8a\x56\xad\x7e\x03\xfb\xcd\xe8\xe9\xba\xfd\x76\x2a\x94\xfc\x55\x35\x78\xf5\x45\x0a\xfc\x66\x6a\x3a\xb6\xed\x72\xe4\x64\xdc\x17\xea\xb2\x63\xdb\x61\x96\xa4\xa0\x4f\xa0\xe1\xb6\x8b\xa2\x8a\xe7\x87\xe2\x49\x1c\x42\x48\x23\xaa\x77\xe3\x49\x9a\x66\x49\x9a\x31\x22\xf2\x24\xeb\x35\xdd\xb9\x72\xb2\x0a\x26\x78\x8b\x9d\x27\x51\x48\x33\x5f\x6c\x48\x6c\x9e\xbe\xac\xee\xc4\xe5\x38\x81\xd6\xe3\x9e\xd0\xde\x7b\x0b\x78\x1a\x91\x06\x6a\x51\x35\x07\x57\xf7\xca\x89\x15\x71\xcd\x27\x70\x31\x0e\xd9\x96\xc6\x9c\x25\xb1\x3e\xf0\xa9\x6e\xe5\x99\xac\xce\xd1\xcd\x4e\xdf\xf6\x86\x56\xb5\x16\xa9\x56\x10\x93\xd8\x86\xdb\x7e\x35\x30\xab\xb0\x30\xc4\xf9\x83\x79\xa4\x0f\x8f\x1c\x77\x4b\x52\x4a\xc9\xe8\xf7\xcb\xfd\x57\x14\x57\xc1\x36\xe8\x8d\xd8\xe6\x1d\x73\xda\xb0\xd2\xa4\xa5\xf9\x6c\xc8\xaf\x64\xd6\x37\x4e\x53\xfa\xfa\x57\x1e\x06\xe5\xa9\x54\xd8\x74\x6a\xc5\x85\xd6\x17\x4a\x22\xf1\x97\x08\x7a\xa1\x31\x84\x3a\xfe\xb1\x19\xe4\x07\x1e\x94\x4b\xd8\x0c\x54\x03\x6f\xde\x57\x61\x46\x54\xb8\x6b\x01\x1b\xa3\xe3\xbf\xac\xdf\x95\xba\x2a\x8a\x70\x92\x65\x18\x2e\x46\x78\x55\x47\xad\x15\x59\x12\x10\x6c\xbd\x11\xa6\x4a\x7a\x45\x8f\xae\x65\xf6\x1b\xab\xc6\xfa\x04\x1e\xe2\xe4\x09\x8f\x97\xa9\x41\xe8\x33\x09\x04\x04\x3b\x61\x27\xab\x55\x71\x60\x94\xc5\xeb\xf2\x3c\xa8\x34\xb1\x54\x1d\x06\xcd\x55\x51\x91\x54\xde\x9e\x34\x10\x89\xba\x2e\xc8\x36\xed\x65\x24\x5e\x53\x9f\xc6\xa1\xd1\x2a\x5d\x52\xf9\x82\x96\x94\xb1\x04\x27\x29\x48\xf9\xb0\x20\x89\xb9\xc8\x08\x8b\x05\x04\x01\x2a\x2a\x50\xf5\xa0\x20\xd0\x4f\xb0\x82\x92\x13\x15\xee\xf3\x88\x05\x14\x42\xae\xf4\xce\x8b\xf1\x6a\x4f\x14\x23\xdb\x76\xc1\x34\x30\x0e\xf4\x39\x88\x76\xd8\xa8\x81\xa7\x2c\xd5\x7e\x30\x7d\xa4\x99\x3a\x64\x02\x3f\x57\xb4\xf6\xb4\x61\xc1\x46\x3e\x81\xbd\xde\xc5\xbb\x26\xb0\x42\x3e\xa8\x78\x8a\xab\x16\xef\x21\xe1\x15\xf2\x41\x49\xc8\xcf\x57\x87\xb5\xb5\x8b\xd9\xb3\xbf\x65\x41\x96\xa8\x8e\x6d\xde\x2b\x29\xb2\xaa\x48\x2c\x07\x1c\x3b\xad\x78\x74\xaf\x4d\x76\x5a\x3b\x8e\x75\xcb\x2c\x06\x41\x2d\x5d\xd4\xb6\x1d\x6c\x08\x9e\x1c\x23\x19\x2d\x0b\x6f\xd2\xa9\xe8\x36\x59\xf9\x15\x57\x97\x34\x91\x8a\x46\x80\x6e\xc8\xa3\xee\xea\x4a\xb8\x00\xce\xb6\x2c\x22\x99\x1e\x4f\xd7\xec\x44\x02\x4f\x72\x34\xc6\x73\x2c\xe3\x01\x1f\xd5\x6b\xb1\x62\x91\x50\x9b\x7f\x24\x8a\xf2\x2a\x1f\x4e\x8e\x23\x2f\x29\x8d\x2b\x16\x60\xdb\xcb\x9d\x28\x7a\x05\xe2\xae\xea\xb4\x97\x5f\xd5\x78\x8a\xdc\x18\x77\xe7\x63\xec\xd9\x2f\x5a\xf6\xf7\x95\x37\x28\xfe\x52\x0c\xa7\x42\xd7\xa3\x2a\x2d\xf9\x78\xed\xec\xd3\x8e\x66\xfb\xb3\x42\x7e\x58\x05\x48\x13\x41\x63\xc1\x48\x14\xed\x7d\x5c\xfc\x34\xc9\x95\xcd\x29\xd3\x6b\x32\x2e\x58\x1c\x88\x5a\xfd\x2c\xff\x34\x96\x85\xef\xce\xcf\xdc\xca\x13\x0a\x7b\xe8\x96\x7f\x86\xef\xfe\x72\x36\xa9\xdc\x75\x3e\x8e\x9c\x3b\xef\x5b\x4f\xfc\xf6\x0a\x67\x46\x74\xe7\x94\xfc\x64\x50\x62\xf5\x21\x48\xe2\x15\xcb\xb6\x34\x3c\x49\x2a\x47\x68\x3a\x20\xe0\x16\xd2\x8c\x5f\xed\x69\xd9\xbe\xd0\x33\x9d\x37\xef\xe0\x34\x0d\xde\x01\xf1\xe0\x63\x4d\x99\x37\x5f\xd2\x2e\xa0\x7c\x64\x50\x16\x80\xaf\x0e\x11\x6d\x3c\x53\x88\x4e\xca\xf2\xa7\x9a\x16\xf1\xa3\x8e\x0f\x29\xd7\x4b\xd2\x54\xda\xfa\x0f\xaa\xe5\x2c\x62\x0f\x34\xc2\x9d\x71\xec\x8d\xe7\xc9\x96\x2a\x17\xc6\x05\xc9\x70\x0f\x9c\x08\xa0\x24\x8b\x18\x36\xd2\xb2\x2d\x6d\x8e\x5e\x78\x12\x24\x22\x5f\xd3\x2a\x9f\xbc\x2c\x63\x5e\xb3\x4c\x1d\xab\xa8\x31\x3c\xa0\xdc\xb1\x33\x71\xa4\x05\xc9\x90\xf3\x70\x65\xd8\x94\x66\x35\xfd\x2a\x65\xa5\xaa\x45\x6d\x80\x32\x37\x97\xcc\xfa\x76\xbf\xde\x05\xd0\x38\x88\xa9\xf6\xcd\xf4\x97\xb1\xbb\xf0\xdc\xe9\xc8\x83\xda\xc6\x07\xe1\xf5\xbd\x0f\x8d\x96\x2a\xe7\x96\xe9\x1e\xaa\x67\x70\xcc\x60\xb7\x6f\x44\x60\x96\x5a\x81\x8b\x98\xb2\xf0\xb9\x45\x13\xc4\x92\xe2\x76\x55\x4a\x32\x19\x89\xe3\xd8\xe8\xc7\xe2\x44\x86\x19\x23\xcf\x31\x9a\x95\xa0\x78\xeb\x4f\x9c\xd2\x3f\xe9\xa1\x8c\x6d\x92\x2c\x79\xe2\x39\xd1\x40\x96\xc9\x23\x1e\xe0\xd3\x17\x06\x15\x8f\x67\x7a\x39\xf4\x70\x35\xc1\xeb\xa2\xe2\x21\x4b\x6e\xc8\xab\x90\x99\x96\xed\xd9\x79\x29\x57\xde\x2b\xb7\x87\x9a\xc6\xf5\x55\xec\xb9\xf6\x1b\x55\x1a\x54\xc7\xad\xba\xf2\xd0\x40\x33\xfc\xe7\x3f\x2b\xcc\xfc\xa6\xbe\x0f\x72\xca\x7f\x7f\xb5\xe1\x14\x7f\x69\x0b\x39\xde\x02\x08\x07\xcc\xe3\xfb\x56\xb3\x28\x7e\x78\xcb\x40\xa4\xfe\x61\xae\x1a\xd6\xf2\x1f\x12\xc1\xd7\x74\xaa\x56\x86\xc1\x57\x6f\xab\x28\x36\x42\xe8\xab\xb7\xd5\x10\xda\x84\xf8\xd5\x5b\x23\x62\x31\x7f\x47\x44\xa5\xae\x47\xcf\xe4\xfc\x4f\x00\x00\x00\xff\xff\x23\x78\x2e\xae\x70\x4e\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/1_base_schema.down.sql"].(os.FileInfo),
		fs["/1_base_schema.up.sql"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
