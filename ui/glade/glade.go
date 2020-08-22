// Code generated by go-bindata. DO NOT EDIT.
// sources:
// contest.glade (30.079kB)

package glade

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _contestGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x5f\x73\xdb\x36\x12\x7f\xcf\xa7\xc0\xf1\x66\x3a\xb9\xe9\xd8\x8e\xe5\xa4\x77\x73\x95\xd8\xb1\xd5\x38\xed\x34\x4e\xa6\x91\xef\x92\x37\x0e\x44\xae\x48\xc4\x20\xc0\x02\xa0\x65\xdd\xa7\xbf\x01\x29\x59\xff\x48\x8a\xa0\x28\x89\x94\xf5\xd2\xa9\x18\x2c\x84\xdd\xfd\xed\x1f\xec\x2e\xad\xee\x2f\x4f\x21\x45\x8f\x20\x24\xe1\xac\x67\x5d\x9e\xbf\xb1\x10\x30\x97\x7b\x84\xf9\x3d\xeb\x3f\xf7\xb7\x67\xff\xb2\x7e\xb1\x5f\x75\xff\x76\x76\x86\x3e\x00\x03\x81\x15\x78\x68\x4c\x54\x80\x7c\x8a\x3d\x40\x57\xe7\x9d\xce\x79\x07\x9d\x9d\xd9\xaf\xba\x84\x29\x10\x23\xec\x82\xfd\x0a\xa1\xae\x80\xbf\x62\x22\x40\x22\x4a\x86\x3d\xcb\x57\x0f\x3f\x5a\xf3\x2f\xba\x3a\xef\xbc\xb1\x2e\x92\x75\x7c\xf8\x1d\x5c\x85\x5c\x8a\xa5\xec\x59\x1f\xd4\xc3\x57\xc2\x3c\x3e\xb6\x10\xf1\x7a\x96\x8b\x29\x25\x6c\xc4\xa7\xcf\x34\x01\x42\xdd\x48\xf0\x08\x84\x9a\x20\x86\x43\xd0\x8b\x98\x33\xe2\x6e\x2c\x2d\xfb\x16\x53\x09\xdd\x8b\xd9\x82\xe9\x7a\x37\x20\xd4\x43\x6a\x12\x41\xcf\x52\x44\x51\x18\x62\x31\xdd\x4b\xef\x46\xb1\x0b\x01\xa7\x1e\x88\x8b\x29\xc1\x45\x42\xb1\x48\x5d\x6e\x75\xf7\x22\xe5\x26\x9b\xb1\x6b\xef\x7b\x2c\x55\x08\x4c\xa5\xcc\xc9\x08\xc0\x5b\x78\x98\xcd\x1d\xe5\x63\x10\x96\xfd\x6e\x8d\xab\x95\x75\x71\x14\xe9\x75\x3f\xbd\xd9\xb4\xf0\x11\xd3\x18\x2c\xbb\xf3\x76\xd3\x42\xa9\x20\x72\x08\x73\x05\xa4\xc7\xbb\xdc\x44\x10\x61\x1f\x96\x08\x56\xce\xb2\x41\x3c\x51\x44\x89\x8b\x15\xe1\x6c\x11\x02\x21\x26\x6c\x3b\xf5\xaf\xac\x0f\xe0\x29\xc2\xcc\xb3\xec\x7b\x11\x6f\x5c\x9c\xc0\xc5\x42\x4a\x60\x26\x29\x56\x78\x48\xa1\x67\x4d\x40\x5a\xf6\x6f\x40\x29\x47\x7d\xce\x14\x48\xb5\x53\xc8\xad\x49\xea\x86\x3f\xa5\xb2\x11\x9c\x2b\xfd\x61\xb6\x34\x43\xd9\x44\x92\x21\x85\x6c\x66\xab\x48\x33\x8b\x86\x0b\x02\x4c\x25\x9a\xb3\xec\x47\x10\x8a\xb8\x98\x66\x12\x2e\x31\x96\xcd\xdc\x1d\xb0\xf8\x06\x8b\xb9\xf2\xf5\x03\x6b\x91\xa6\x02\x97\x55\x39\xcd\x3e\x74\xfe\xc1\x7f\x57\x10\x4e\x4f\x0e\x2c\xbe\x25\xfa\x4c\x2b\x84\x15\x4f\xbf\x0d\x07\x99\x6e\x05\x0f\x81\x66\x02\xdb\xd1\xc7\x36\xd9\x2a\x96\xe0\xc4\xcc\x03\x41\x09\x2b\xc3\xc5\xa2\x6d\xc8\x78\x18\xae\xeb\xb7\x48\xc4\x53\xdf\x99\xd2\xe5\x48\x78\x2b\x29\x6f\x2b\x69\x94\x8b\x99\x62\xc6\xd6\xb1\xf3\x09\xc6\x39\xcc\x6d\xcd\x60\x1d\x4c\x66\xed\x91\xc2\xca\xfe\x04\xe3\xf3\xf3\xf3\x2a\x1b\x18\x82\x69\xba\xc9\x42\x60\xc9\xfa\xd7\x02\x6d\xd4\xa3\xa9\xcf\x11\xb0\x96\xaa\x4a\x1f\xfd\x45\xe9\x6a\x80\x1f\xe1\x5a\xb6\x54\x5b\xfa\xf0\x08\xcb\x23\x55\xd8\x00\x22\x2c\xb0\xe2\x62\x59\x73\x72\xf6\x58\xab\xef\xb2\xd9\x9a\x3b\x38\xbc\xdf\x3f\x45\x5c\xa8\x3e\x1e\x0a\x42\x29\x6f\xb6\xb0\x4c\xd2\x92\x94\x2f\x34\x63\xec\x48\x2d\xa0\x48\xa7\xd7\xbf\xfe\x7e\x7b\x74\xfa\xd4\x4c\x1d\xa9\x2e\xcb\x78\xb3\x4e\xb3\x15\x7a\x70\xe4\xff\x19\x13\xd5\x6c\x11\x99\x60\x5e\x73\xd3\x7c\xa4\xe7\x13\xe6\x10\x65\x13\x64\x2e\xae\x7e\x99\x5e\xaa\x01\x15\x09\xab\x69\xd7\xe9\xf4\xe0\x2d\xb9\x50\xe7\x4a\x79\x2b\x49\x6f\x2b\x6d\x54\x8b\x3f\x49\x79\xeb\x4f\x4b\xca\xc7\xe3\x55\x66\x1c\xed\xcd\xb3\x20\xd4\xc5\xae\x0b\x14\x92\x30\x86\x1e\x60\xd2\xb3\x88\x85\x24\xf1\x19\xa6\x3d\x0b\xbb\x8a\x3c\x62\x05\x16\x0a\xb9\x47\x46\x04\x84\xd6\xc8\xaf\x7f\x38\xfd\xcf\x9f\xee\xbf\x7c\xfe\xe8\xdc\x5d\x0f\xfe\xb0\x2e\x5e\x88\xdf\xfa\x0d\x68\xd4\x42\xaf\xa5\x8f\xdd\x12\x9f\x95\x23\xe1\xad\xa4\xbc\xad\xa4\x51\x2d\x1e\x4b\x73\x76\x3d\xe4\x71\x4b\x53\x20\x3b\x39\xfb\x4b\x4e\x78\xb2\x16\x76\x23\xec\x3e\x10\xe6\x17\x77\x27\x66\x4d\xaa\xe2\xd6\xc4\x0a\xd1\x88\x50\x6a\xd6\x04\x89\xb8\x24\x69\xdf\x66\xad\x6f\x38\xe3\x60\xed\xb8\x6b\x7c\x96\x69\xeb\x0c\x5c\xc1\x29\x05\x6f\xb1\xb5\x47\xb9\xff\x5f\x02\xe3\x3e\x67\x0a\x13\x06\x62\x43\x97\x67\x4c\x3c\x15\x38\x02\xfe\x8a\x41\x2a\xcb\x7e\xf7\x6e\xad\x27\x9a\x47\xb9\x6d\x7f\xc8\x84\x4c\x06\xd8\xe3\x63\x47\x7b\x36\xcb\x26\x6c\xdb\xae\xd2\xbd\x00\xd0\x42\x5a\x92\x58\x89\x78\x12\x00\xf1\x03\x35\x97\xd6\xe5\x9b\x1c\x05\x6f\x21\xb1\x0d\x52\x33\x0e\x47\x85\x9d\xd9\xc2\xd3\x56\x25\x0c\x64\x02\x4b\x27\xe2\x94\xb8\x13\xcb\x66\x58\xc5\x22\xbb\x77\x99\xfb\xdd\xdb\x6f\x11\x00\xf6\x40\x48\xc7\xa5\xc4\x7d\xc0\x89\xd8\x8d\x65\x07\x4c\x13\x3a\x12\xb0\x70\x83\x0a\xf4\x32\xe0\x63\x27\x15\x23\x88\x2a\xca\x9b\x1e\xc0\x17\xc4\x73\xb4\xd7\x96\xc5\xad\xe0\x0d\xdb\x28\x01\x30\xdb\xa6\x64\x2e\x91\x8c\x9f\x30\x4c\xcf\x92\x8f\x3d\x4b\x02\x05\x37\xf5\x6d\xa5\xf2\x0a\x6d\x69\x83\x67\x9a\x99\xb9\xcd\x9f\x64\x24\xac\x87\x0f\x0f\x26\x9e\x69\x9b\xe8\xb0\x36\xfb\x31\x63\xa0\xa6\xe8\xa0\xb0\x8a\xe5\x70\xd6\xf6\x97\xcf\x1f\x77\xeb\xd7\x8d\x82\x6b\x88\x85\x4f\x98\x43\x61\x94\x31\xda\xb2\x91\x4c\x68\x77\x5c\x81\x4e\x2a\x2c\xaa\xd0\x81\x46\x87\x31\x95\xe2\x91\x65\xff\x64\x48\x34\xe4\x4a\xf1\xd0\x80\xae\xf4\xbc\x48\x16\xb1\x8c\xb0\x4b\x98\x6f\xd9\x9d\x3c\x48\x36\x3d\xe5\xc2\xee\xc3\x34\x3d\x00\xe6\x99\x9b\x62\x2e\xdf\xf5\x98\xe2\x07\x41\xbc\xd4\x0a\x81\x29\x31\x49\x3e\x36\xd5\x0a\xd7\x46\xd8\x36\x50\x4d\x8d\xd0\x94\x2c\xb1\x8a\xd2\xc9\xe6\x8a\x55\x94\xa6\x13\x7c\xec\x3c\x83\xbb\x34\x95\xcb\x69\x1c\xb2\x32\x84\x25\x13\xce\x8f\x69\x4d\x20\x09\x7f\xfa\x7f\xb3\x3a\x14\x8d\xc8\x16\xf3\x8b\x17\xf7\x01\x10\xf1\x6f\x93\xbd\xd2\xab\x85\x1b\x60\x9d\xfa\xe4\xfa\xb1\x2c\xca\x27\x4c\x89\x9f\x7f\x7d\x42\x05\xd7\xcb\x6c\xaf\x94\xc9\x2a\x8c\x94\x83\x95\xc2\x3a\xb5\x33\xca\xe3\x15\x8f\x9e\x09\x73\x62\x38\xca\x76\x1e\x68\xdb\x52\xd8\x7b\xed\x3e\xe6\x63\xbc\x92\xf8\x2c\x7d\xb4\xf9\xd4\xfa\xbf\x96\xbd\x44\xb5\xef\xcb\x8b\x29\xa9\xe2\x9c\x2a\x12\x39\x0a\x9e\x54\x3e\x28\x51\x7f\xca\x93\xc9\xd6\x0b\xd3\xa2\xf9\xdb\xeb\x8d\xf7\x06\xc1\x02\x24\x55\x86\x60\x8e\x51\xe6\xc6\x3b\xb4\x7b\xd8\x2a\xad\xb2\x2f\x10\x71\xa1\xcc\x90\xbb\x4a\x78\x24\xe0\x4d\x39\x32\xd9\x38\xc4\x4f\x0e\x05\xe6\x6b\x3d\x5e\x55\x76\xc8\x46\x94\xfa\x2b\xab\x53\x97\xb2\xb5\x2f\x83\x7b\x93\x3d\x09\x8b\x62\xe5\x44\xb1\x88\xb8\x04\xcb\xf6\x88\x4f\x94\xdc\x9b\xad\x1a\xb1\xdf\x8c\x70\x91\x98\xcf\xa7\x38\x1c\x82\xa8\x60\x77\x0b\x84\x47\x62\x77\x29\x47\x95\xcd\x27\x37\x25\xcc\xa2\x2c\x65\x00\x7f\x6f\x0f\xfc\x73\x73\xfd\x6c\x7d\x34\x01\xfe\xe1\xa4\x4a\xcc\x59\xa2\x6a\x3d\xf0\xef\x26\xd3\x68\x83\x5e\xf7\x95\xa0\x67\xe2\x1f\xa7\xb0\xb3\xdb\xb0\x93\xd5\xd4\x17\xf3\xa6\xbe\x2f\xf0\xf0\x2c\xd5\xb6\x59\x5b\xbf\x29\xe1\xac\xf0\x82\xb6\x63\x7b\xae\x12\xcb\x96\xa8\xf6\x62\xcf\x12\x98\x24\x8a\x3c\x56\x69\x3f\x6c\x73\x91\x2f\xe3\x0c\x4e\x21\xf0\x00\x21\x70\xcf\x26\xb3\x5a\x79\xca\x9a\xf4\x6f\x78\xe5\xe9\x6e\x62\x54\x76\x02\x4a\x49\x24\xc9\xff\xc0\xb2\x93\x3a\xff\xa9\x64\xd5\x18\x04\x86\x20\x25\xf6\x21\x7d\xd2\x3e\x1c\xa6\xa7\x3f\x52\x54\x14\x54\x84\x72\xab\x48\xff\xdc\x2f\x92\xfa\x3c\x1c\xf2\x1b\xfe\x74\x9f\xc4\x0f\x0d\xa8\x21\x66\x5e\xf2\xb4\xa9\x68\xda\x18\x86\x6f\x30\xf3\xd2\x74\xfc\xc7\x61\x71\x3a\xbe\x96\x48\x0e\xe7\x89\x64\xc4\xa3\x38\x3a\x6c\x0e\x59\xb9\x7c\xb9\x67\x77\xb4\x0e\xa2\x90\x7b\xd0\x6e\x10\xdd\x71\x0f\xa6\x77\xba\xd0\x10\x44\x61\xa3\x40\x64\xe4\x84\x0e\x08\xa2\x9b\x58\xa9\xd9\x3c\x8b\x00\x09\x6a\xfa\x60\x9b\xf0\xf2\x45\xef\xd3\xf8\x71\x32\x01\x2e\x90\x47\x90\x8e\x07\x23\x1c\x53\xb5\x83\x32\x45\x22\x08\xf4\xfa\x07\xaa\x7e\x7e\x3f\xe8\xff\xe0\xab\x9f\x0b\x30\x5d\x33\x00\x8d\x72\xbe\x86\x00\x90\x72\xbf\x06\xf8\x7d\xe4\xfe\x09\x7c\x89\x18\x90\x0a\x00\xfd\x39\xf8\x3c\x85\x20\x53\x20\x5a\x02\xc2\x43\x96\x57\xbf\xb9\x01\x66\x3e\x98\xd6\x63\x16\xc9\x8e\xa1\xc0\xfa\xfe\x29\x65\x68\x3f\xb7\xce\x52\x55\x95\x6f\xfd\xa0\xc0\xb4\x6b\xc6\xae\x51\x51\xa8\x19\xa5\xc4\xa4\xc1\x55\x09\xbe\x6b\x94\xad\x47\x70\xda\x1b\x3b\x81\xf8\x90\x0e\xb8\xf9\xc3\x96\xcf\x63\x93\x39\x8d\x83\xfa\xc7\x26\x1f\x60\x02\xe2\x34\x36\x79\x90\xb1\xc9\xdc\x4b\xd9\xa6\xb1\xc9\xd2\x84\xfa\xeb\x02\x1e\x72\x1f\x18\xf0\x8d\xef\x06\x55\x48\x91\x47\x97\x35\x64\xc8\xb7\x46\x25\x8e\x63\x4d\x90\x07\xc0\x3c\x74\x87\x85\x2b\x38\x7a\x7d\x7b\x69\x58\x6e\xb8\xbd\x9c\xd7\x1b\x92\x37\x73\xc0\xdb\x79\x51\xa1\x2d\x85\xf2\x25\xc0\x76\xea\x00\xac\x51\x39\xe5\x65\x00\xb6\x63\x0a\xd8\x4e\x8b\x00\xbb\xe7\xfb\xdf\x12\x60\xaf\xea\x00\xac\xd1\x20\xc2\xcb\x00\xec\x95\x29\x60\xaf\x5a\x04\xd8\x3d\x8f\x21\x2f\x01\xf6\x6d\x1d\x80\x35\x1a\x03\x78\x19\x80\x7d\x67\x0a\xd8\xb7\x2d\x02\x6c\x81\x83\xda\x71\x95\x62\x74\x59\xb6\x36\x71\x88\xc2\x42\xc9\xb7\xcc\x5b\xd0\x73\xcc\x6d\x76\x17\x98\xfa\xae\x55\xdf\x39\xa9\xfe\xa0\x6f\xcb\x1c\x50\xf5\x57\x27\xd5\xd7\xa6\xfa\x2a\x23\x2e\x07\x54\xfd\xdb\x93\xea\x6b\x53\xbd\xf9\xc0\xfa\xde\x55\xbf\x98\x9c\x4a\xc5\xa3\x3a\xd2\x53\xa3\x42\x76\xda\xd3\x98\x7f\x75\xd3\x9b\x19\xfb\xc8\x6c\x15\x8f\x92\x76\x70\xff\x6b\xfa\xcf\x21\x91\x92\x70\x56\x25\xcd\x7d\xd7\xa2\x34\x77\xcf\xd0\x6f\xf4\x40\x68\x30\x9d\xe1\x34\x1e\x26\x0e\x4a\xb5\x63\x32\x95\x98\x6b\xd1\xe3\x28\xdc\x9b\xb3\xad\x3c\xe2\x7e\x42\xcf\x3a\x7a\x0a\xfe\x2a\x49\x26\x59\xd5\x3f\xbe\x95\x0f\x9d\x41\x04\xe0\x19\x4d\xd2\x7f\x8f\xa5\x22\xa3\x89\x65\x27\xbd\xaa\xbd\xe1\xae\xf2\x10\xe0\x9e\x71\x37\x88\x08\x5b\x0a\xda\x5a\xc0\x4d\xce\xd9\x36\x86\xbb\xfe\x57\x94\xa0\x64\x3a\x4b\x2a\x8d\xde\x0f\x9c\x21\xdd\x05\xa6\xcc\xde\x2a\xaa\xee\x28\x17\xdf\x49\x34\xcf\xee\x67\xd3\x0a\x46\x94\x6b\xef\x24\x1a\x51\x9b\xbf\x3f\xb8\xb2\x01\x5e\xf8\x69\xbb\x95\xdf\xba\x33\xca\xf5\xe2\x10\x04\x71\x4b\x60\x68\x2d\x91\x91\x47\xf4\xfe\xe2\x29\x52\xd5\x90\xe7\x2c\x8d\x41\x18\x59\xc3\x2e\xa2\x9c\xce\xd8\xef\xbf\x35\xff\x46\x5a\x0b\xf4\x5a\x34\x2a\x94\xc3\xef\xe6\x51\xa1\x65\x1e\x33\x7f\x18\xb4\x7b\xb1\xf0\x33\xa9\xff\x0f\x00\x00\xff\xff\xbe\x77\xfa\xed\x7f\x75\x00\x00")

func contestGladeBytes() ([]byte, error) {
	return bindataRead(
		_contestGlade,
		"contest.glade",
	)
}

func contestGlade() (*asset, error) {
	bytes, err := contestGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contest.glade", size: 30079, mode: os.FileMode(0664), modTime: time.Unix(1598084446, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x45, 0xeb, 0x3e, 0x19, 0xf9, 0xaa, 0x45, 0x28, 0x51, 0xc9, 0x16, 0xa9, 0xb2, 0x52, 0xf6, 0x5b, 0xb8, 0xf8, 0xfb, 0xd3, 0x7d, 0x3b, 0xff, 0x17, 0x36, 0x5c, 0x2c, 0xb1, 0x7, 0x55, 0x84, 0xa5}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"contest.glade": contestGlade,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"contest.glade": &bintree{contestGlade, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
