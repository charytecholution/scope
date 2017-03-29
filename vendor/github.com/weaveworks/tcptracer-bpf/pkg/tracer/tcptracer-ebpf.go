// Code generated by go-bindata.
// sources:
// ../dist/tcptracer-ebpf.o
// DO NOT EDIT!

package tracer

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
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

var _tcptracerEbpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5c\x01\x6c\x5b\xc7\x79\xbe\x47\x8a\x26\x25\xcf\x93\xe6\x96\x09\xcd\xa1\x98\xd2\x05\x89\x26\xa4\x29\x65\x4b\x32\xab\x36\x9d\xea\x36\xa9\xa6\x19\x95\x96\x89\xad\x60\x2c\xa5\x18\x96\xb6\x14\x3a\x36\x25\xd2\x96\x5e\x98\x6d\x2e\x8a\x64\x02\x91\x05\xb2\x93\x16\x42\x96\x66\xa2\x2c\xc7\xea\x9a\x0d\x1a\xb6\xc1\x01\xba\x81\x5a\x97\x61\xc2\x16\x0c\xc6\x5a\x0c\x1a\xea\x01\x42\x97\x75\xda\x86\xa2\xda\x9a\x79\x5a\x6a\x84\x03\xef\xff\x1e\xdf\x7b\xff\xbb\xf7\x48\xd9\x69\x4b\x20\x3d\xff\xdf\xbb\xff\xfe\xff\xee\xfe\xfb\xef\xbb\x7b\x4f\xfd\xed\x87\x8f\x3f\xe2\xd3\x34\x61\xfc\x34\xf1\xbf\xc2\x94\xcc\x5f\xe4\x13\xe6\xbf\x07\xf1\xbf\x3f\x2f\x34\x51\xb9\x8b\xb0\x67\x84\x10\x3f\x2b\x84\x28\xb6\xed\x54\x6b\xb2\x9e\xca\x49\xbc\x18\xdd\x95\x72\x65\x99\xea\x05\x7d\x42\xec\x54\xab\xd5\xca\x65\xc8\x7e\x21\x76\xab\xd5\x6a\x84\x19\xbd\xd6\x62\xb6\xeb\xab\xc9\xc0\xbf\x8c\x52\xbf\x7b\x94\xd9\x1d\x92\x76\xae\xa1\x9d\x62\x74\xd0\x61\x77\x48\x61\xe7\x19\xd9\x67\x21\xc2\x62\xbf\x7c\xa2\xef\x23\xbc\x19\x3d\xbf\x10\xe2\x54\x50\x88\x4e\x21\xc4\x2c\xca\x44\xd0\xa7\x71\xfd\x41\x0f\xbb\x95\x20\xc9\xe1\xe0\x8f\xa8\x5f\xd3\x90\xb5\xff\x23\xb9\x8c\x7e\xf9\xd0\xaf\xee\x18\x8d\x6f\x1e\xf5\xfc\x9f\x91\x7e\x24\xfc\x3f\x94\xed\xe9\x33\x21\x89\x3f\x76\xf0\xbf\x68\x1c\x5b\xa1\xf7\xd2\x76\x95\xca\x2d\x94\x9b\x28\xaf\xa3\xdc\x40\xb9\x8e\xf2\x75\x94\x6b\x28\x57\x51\x2e\xa1\x5c\xac\xfb\xe5\x93\x7e\x4d\x12\x1e\xcd\x91\x7f\xd3\x1d\x98\x87\x39\x92\x67\x22\xf0\xbf\x8b\xea\x75\x5f\x00\xde\x09\xbc\x13\xf8\x3c\xe1\xb9\x2e\xc2\xdb\x16\x48\x3e\x17\x23\xf9\x04\xd9\xd5\x0b\x71\x1a\x9f\x39\x8a\xc6\x62\x09\xfe\x0d\x93\x7f\xfa\xdc\x10\x3d\x9f\x19\x45\xfb\xe8\x4f\x69\xcd\xd6\xcf\xd4\x4c\x41\x3e\x4f\x77\xdf\x80\x3f\x73\x90\x69\x9c\x2a\xbf\x44\xe3\x37\xdb\x4a\xf3\x9b\x7d\xe9\x2d\x89\xcf\xfa\x84\xa8\x79\x94\xed\xfe\x1e\xda\xc9\x41\x8f\xc6\x75\x6a\xe6\x82\x94\xb3\xdd\xdb\x78\xfe\x05\x29\x9f\xf2\x51\x3b\x29\x7d\x42\xca\x23\x97\xa8\xfd\x94\x3e\x49\x65\xfe\x34\xd5\xf3\x53\xbd\x91\x17\x30\x7f\x03\x98\xa7\x6e\x9a\xa7\x8c\x96\x92\xe3\x1e\xd6\x9e\x46\xbc\x5c\xa2\x38\xd6\x34\x89\x07\xc4\x1f\xd3\xfc\x23\xee\xb2\xd1\x47\xa5\x5e\xfb\xb3\x24\x57\xca\x88\x4f\x4d\x88\x47\xab\xd5\xaa\x31\x8f\x95\xb4\x19\x9f\xb5\x25\x38\x55\x26\xbd\x6c\x94\xfa\x11\x10\xaf\xa0\x9f\xe3\x55\x2a\x6f\xd1\xf3\xee\x31\x94\x37\xe1\xe7\x24\x9e\xef\x22\x1e\x68\xfe\x8d\xf6\xdb\xdf\xe7\xf4\x63\x1c\x7e\xb4\x58\xec\xeb\xd3\x9d\x0d\xf5\x6e\x29\xf5\xba\x1a\xea\x8d\x59\xfa\x6d\xea\xc5\x1a\xea\xdd\x54\xea\x51\x3c\xb6\xdf\xe5\xac\x3f\x89\xfa\x21\xc5\xf8\xea\xd3\x14\xbf\xfa\x65\x1a\x2f\x95\xbd\x5d\x45\xff\x52\x2b\x18\xff\x61\xca\x07\xa9\xcb\x34\x0f\xe9\x81\x5e\x8a\xbf\x65\x63\x3e\xe2\x24\x97\x6f\x62\x1e\x07\xa8\xfe\x15\x9a\x97\xf4\x89\x8f\xc9\x32\xdc\x92\xb5\xc5\xdd\x2c\xca\xb0\x9f\xe2\xf2\x73\x3e\x21\xaa\x55\x21\xc2\xbe\x93\x24\x6b\x90\x35\x8a\x6b\x23\x4f\x15\xa3\x93\x2c\x0f\xcf\xd9\xf2\x82\xb5\x5f\x39\xf4\xab\x03\xfd\xea\x60\x79\x33\xc6\xf6\x89\x49\x45\x1e\x0d\x88\x47\x65\x19\xd6\xee\x97\x79\x30\xac\x3d\x24\xc7\xa7\xb6\x0e\x02\xf2\x39\xad\x7f\xc3\x9f\x74\x74\x0c\xfe\xd0\xf8\xb5\x3f\xe2\x1e\x87\xde\xf3\x35\xee\x3a\x5f\x63\x8a\xf9\x3a\x25\x30\xae\x28\x6b\xfd\xf8\xef\x6a\xb5\x6a\xf4\xe3\xb1\x8e\x9e\xba\x9f\x9a\xf4\x6f\x1d\xeb\x2d\xc6\xfc\xa7\x79\xd6\xcb\x94\x47\x55\xf1\x66\x5d\x0f\xdc\xff\x54\x19\x71\x12\xbd\x81\xf6\xef\xb5\xb5\x9f\x8d\x52\x9c\xb4\xef\x73\xb6\x7b\xd3\x33\x4f\x18\xf1\xf5\x16\xda\x3d\xc4\xfc\x46\x1e\x28\x77\xba\xfa\xbd\xeb\xe9\x37\xe2\x35\xba\x85\xf6\x0f\x30\xbf\x4f\x52\xbb\x9f\x75\xb6\x7b\xd2\xd3\xef\x93\xf0\xfb\x7b\x68\x97\x36\xda\x54\x79\x0c\xf6\x36\x1d\x3c\xa2\x99\xb8\x6c\xc4\x5f\x7e\x6a\xbc\x29\xc2\x79\xd3\x28\x5b\xaf\x43\x0e\xbb\xa3\x5e\xfc\x05\x71\x12\xde\x77\x4b\x34\xab\x57\x5b\xa7\xfa\x49\xe1\xa8\xef\xc5\xcf\x4c\x9e\xf4\x03\xea\x47\x9d\x27\x7d\x9f\xe4\x65\xde\x8f\xce\xaa\xbd\xde\xa3\xc4\x93\xb4\x7f\x23\x9e\x34\x0d\x9e\x74\xcf\xbf\x0a\xbb\x1e\x78\x51\x14\xfb\x6d\x14\xbc\x08\xeb\xb1\x18\x05\x8f\x88\x82\x47\x44\xc1\x3b\xa2\xe0\x45\xd1\x45\x94\x0b\x28\x89\xcf\x5c\xf3\x83\x27\x0d\x8c\x63\x7f\x9c\x00\xdf\x20\x9e\xa4\xe7\xc1\x8f\x06\x72\xb6\xfd\x53\xcf\xd3\xfe\xa7\x17\xc0\x87\x86\x89\x37\x15\x07\xc0\xab\xf2\x31\x3c\x8f\xe3\x39\xec\x0e\x80\x47\xe5\x91\xaf\x0a\x43\x78\x0e\x3f\x07\xc0\xa3\xf2\xe0\x47\x03\xab\xb6\x7e\xa5\x0a\xe0\x45\xc3\x6f\x82\x6f\x81\x17\x0d\xd3\xb8\x54\x1e\xa6\xf1\x9a\x0d\x80\x17\x9d\xf8\x36\xf1\xa2\x16\xf0\xa2\xe1\x7f\x44\x3b\xe0\x45\xc3\x34\x8e\x53\x05\xf0\xa2\xe1\x4d\x3c\x07\xdf\x69\x01\x2f\x3a\x47\xfb\xcd\xc8\x18\xe6\x63\x18\x3c\xad\x40\x7c\x29\x75\x0e\x3c\x0a\x76\x47\xe0\x47\xf1\x84\x31\x2f\x71\x94\xb4\x2f\x26\x7c\x5f\x13\xb5\x10\xad\x5c\x45\x9c\x05\x85\x78\xbd\x5a\xad\xb6\xf7\x23\xfe\x2c\x79\x22\x66\xd9\x8f\x1c\x79\xbf\x8c\x76\xef\xa6\xb8\xd1\xcb\x31\xc8\x68\xc7\x12\xc7\x9d\x6c\x9d\x8e\x2b\xf7\xad\x17\xa9\x1f\x21\xec\x0b\x28\x13\xa1\x17\x58\x3c\xc2\x0e\xe2\xb9\x18\xdd\x41\xb9\x8d\xbc\xf5\x0e\xca\x5d\xe0\x5b\xc8\x5b\xd8\x27\xa6\x69\x7e\x39\x9f\xa9\xf9\xd5\xe1\xd2\xff\x90\x6d\xbf\x73\xd7\x8f\x30\xfd\xce\x3d\xe8\x5b\xf5\x76\xf6\x68\x37\xc4\xf4\xb7\x95\xfa\x8d\xf9\xdf\x3b\xb7\xc9\xff\x76\xf7\xc8\xff\xb6\x9a\xe2\x13\x5b\xae\x7c\x42\xc9\x6f\xbf\x86\xb8\x78\x89\xf2\x89\xfe\x2a\xe2\xe3\x09\xca\x2b\x95\x14\xe9\x8f\x5c\xa4\x52\xbf\x8a\xb8\x79\x0a\x79\x65\x15\xe7\xc1\x12\xe5\x9b\xa9\x15\xc4\xd1\xf0\x1c\xf8\xe3\x2e\xf8\xe3\x53\x8c\x5f\x3e\x4d\x7c\x51\xfb\x38\xf9\xf9\x45\xd8\xa9\xe7\xd9\x8f\x12\x8e\xfb\x02\x93\x27\xd2\xf8\x54\xee\xe1\x78\x9f\x94\x8b\xc3\x21\xe4\xa5\x0e\xe4\xbf\x08\xcb\xe7\x83\xb6\xf5\xad\x97\x47\x6d\xeb\xdc\x3a\x5e\x71\x05\xaf\x4c\xb4\xd2\x46\x65\x8c\xcb\x29\x0d\xeb\x0e\x65\x44\x6e\xfb\x26\x1f\xfb\xc0\x3d\xb4\xe1\x24\x82\x87\x6c\xfd\x6c\xa4\xf7\xd8\x3d\x34\x81\xb3\xfb\xe8\x79\xfa\x09\x9a\x17\x63\x7d\xa7\x4b\xf7\xc2\x7f\xea\xef\x0d\x8d\xfc\xcb\x46\xbb\x80\x47\x18\xde\x0d\xbc\x03\xeb\xfa\x01\xc7\xfe\xcc\xf3\x4d\x4c\x99\x6f\xd0\x7f\x8b\xde\x78\x13\x7a\x7b\xe1\x31\x21\xc9\x27\x06\xd9\xbc\x51\x9e\xd7\x4f\x51\x5e\xd7\x27\x73\x0e\x3f\x36\xbd\xf6\x7d\xdc\x5f\x84\x5b\xaf\x90\x3e\xe2\x39\xa1\x5d\x96\xfb\x79\x65\x5a\x3d\x2f\x35\xfe\x1f\x94\xf5\x5e\x91\xed\x70\x7e\x93\xc5\x3e\xd7\xfc\xfd\x10\xf1\x42\x1d\xcf\xc3\x5a\x9b\xb4\x7f\xe7\xed\x96\x30\x3f\xbf\x83\xf6\x69\x9f\xae\x7c\x89\xea\xa9\xf2\xc1\x9a\x22\xff\x4c\x95\xc9\x6e\x42\xfb\x7e\xd5\x67\xe3\xdd\x6b\x88\x9b\x1d\xf8\xb7\x85\xf2\x56\xd5\xea\x77\x1a\xfb\x89\x7e\x31\xd2\xb4\x7d\x5b\x3e\xba\xd8\xd9\x50\x6f\x47\xa9\xd7\xd5\x50\x6f\x4b\x95\x6f\x2f\xc6\x1a\xea\xdd\x52\xea\x21\x4f\x47\x9c\xf5\x87\xac\x79\xfa\x4b\xe6\x3c\x91\x9e\x91\xa7\x87\x5c\xf3\xf4\xb6\xf2\x9c\x8e\xf1\x1f\x5e\x44\x1e\xdd\x41\x1e\x7d\x19\x79\x16\xf3\xd1\x4d\xbc\x6c\x0a\xe7\xb3\x6c\x74\x99\xea\x5f\xa1\x79\x49\x9f\x58\xa1\xbc\xdb\xf2\x29\x69\xcf\x79\x4e\xa7\x84\x6b\x9e\xd3\x7f\x99\xe4\x7a\x9e\x7d\x88\xf4\x1c\xf7\x95\xb4\x4e\x02\x82\xce\x9f\xd6\xb8\x5d\xf4\x8c\xdb\x23\x52\x36\xe3\x5f\xcd\x8f\x55\x71\x23\x79\x83\xdf\x3e\xbe\xcf\x48\x1f\xcc\x7b\x5c\xfd\x32\xe5\x11\xde\xcf\xe2\x80\xc9\xf3\x35\x19\xb7\xb4\xff\xa5\xca\x7f\x8a\x38\xfe\x0d\xc8\x8b\x90\x73\x90\x5f\x86\x5c\x80\x4c\xe3\x5d\x39\x48\x7e\xdc\xc0\x7e\x91\xed\xa6\xfd\x6f\x16\xf9\xd7\xcc\xc3\xd8\x07\xcb\x2b\x68\xe7\x69\xe4\x65\xa3\xdf\xd8\x57\xcb\x46\xff\xb1\x0f\xaf\xa0\xff\x2d\x14\x5f\xc6\xfe\x13\xf1\xd9\xc7\xf5\x5a\x40\x88\x38\xc6\xe1\x50\x93\xf3\xe0\xb7\x9c\x57\xcd\xbc\xe4\xa3\xbc\x84\x7a\xf5\x73\xdf\x1d\xe7\x95\xff\xa9\xb6\x29\xf2\x9d\x79\xde\xd9\x23\x2f\xbd\x38\xaa\xf4\x47\xc5\x4b\x87\x54\xfc\xce\x43\x9f\xf3\xd2\xc5\x3d\xe8\x37\xe4\xa5\x1e\x76\x9b\xe2\xa5\x4d\xe4\x3b\x25\x2f\x6d\x22\xdf\x29\x79\xa9\x47\xbe\xdb\x6a\x2a\xdf\xed\x91\x97\xd6\xef\x01\x29\x6e\xf4\x2b\x88\x0f\x9c\xcf\x46\xc2\x88\x47\xe4\xc1\xe2\x00\xad\x1f\x7d\x05\x7c\xd4\x38\x27\x2e\x23\x7e\xba\xaf\x23\x2f\x82\x8f\x46\xff\x81\xd6\xe1\x12\xf8\x68\x1b\x9d\x3b\xc3\x81\x32\xb5\xdf\x4b\xed\x87\x5b\x7e\x5f\xca\x66\x3e\xfc\x2a\xc9\xf5\x7c\xf8\x7b\x94\x5f\x1c\xf9\x70\x9f\xcc\x87\xd6\xf5\xb7\xe6\xb9\xfe\x2e\x39\xf6\x7b\xb7\xfa\x94\x37\x9f\x63\xfc\x48\x7d\x7f\xa0\x8a\xdf\x1f\x4f\xde\x7c\xd9\x96\x37\x75\xac\x77\x7d\x19\x79\xac\xdb\x7e\x7f\x6b\xe4\x3b\x7d\x19\xf7\x22\x78\x5f\x53\x8c\xe2\x7d\x4e\xf9\xba\x4b\x5e\xa5\x7b\x09\xce\x97\x2b\xef\xa7\x32\x49\xd7\x90\xf2\xfe\x40\xe6\x5d\x94\xd9\xe1\x4b\xf5\x7c\x2c\xf9\x73\xf4\x45\xd8\x59\x65\x79\x77\xe9\xb6\xf3\xee\xa8\x25\xef\x06\x04\xf5\xcf\xe0\xb5\x9c\xc7\x9a\xf7\x66\xf1\xba\x2c\xd7\x4f\x08\x71\x3d\xb5\x47\x7e\x5b\xbf\xd7\x7a\x80\xf4\xc1\x67\x13\x5a\xb7\x8c\xaf\xe2\x13\x38\xdf\xe0\xbd\xcd\x35\xd8\x29\x96\x68\xfc\xd3\x25\x8a\x7f\xe3\xfc\x59\xc1\xfb\x22\xd5\x7a\xdd\x50\x9e\x93\x8d\x75\xbe\xe1\xba\xce\xbf\xad\x58\xe7\xc5\x12\xc5\x59\xb6\xb4\xe4\x88\x7f\xaf\xf7\x9d\x4e\x1e\xdd\x4e\x3c\x7e\x12\xe3\xf7\x6a\x4c\x79\xee\xcb\x46\xef\xd4\xce\xaf\x7b\xef\x8b\x1e\xe3\xb6\xa4\xdc\x17\xe1\x0f\xf3\x3b\xa1\x6d\x33\xfe\xbd\xd0\x80\x7f\x0f\xda\xf9\xb7\xf1\x9e\xcc\xc3\x9f\x05\xaf\xf7\x64\x1e\x7a\x4a\xfe\x6d\xdc\x93\x78\xe8\x29\xf9\xb7\x71\x4f\xe2\xa1\xa7\xe4\xdf\xc6\x3d\xc9\x21\x67\xfd\x41\xeb\x7e\xf4\x82\x39\x9f\xf6\x38\x1d\xdc\x23\xff\xc6\xf8\xe3\xde\xd3\xe4\xdf\x2b\x8c\x7f\xaf\x32\xfe\xfd\x07\x8c\x7f\xbf\x06\xfe\x4d\x0b\xd6\xc9\xbf\xe9\xa0\x6c\xee\x37\x14\x7f\xc6\x7e\x93\xd0\xfe\x9e\xfc\xb5\xc4\x6f\x2d\xae\x02\xe2\x2f\xdf\x93\x38\x4c\x68\x2b\x8c\x9f\x19\xf7\x25\x0b\xb7\xc7\xcf\x8c\xfb\x37\xe6\x8f\x8a\x9f\x0d\x7a\xdd\xdf\x29\xf4\x39\x3f\x5b\xd8\x83\x7e\xd3\xf7\x86\x0a\xbb\x7b\xba\x37\xf4\xb0\xeb\x79\x6f\xe8\xa1\xe7\x79\x6f\xa8\x58\x0f\x5b\x4d\xad\x87\x3d\xf2\xb3\x15\xc4\x85\xf1\x9d\xc4\x65\xe3\x7d\xc1\x2a\xf6\x75\xc4\x49\x37\x78\x40\x79\xdb\x76\xae\x9b\x5a\x42\xdc\xb4\xd1\x3e\x39\x85\xf7\xc8\xd9\x13\xdf\xa4\xf5\x72\x15\xbc\xec\xa9\x37\x28\x3f\x1e\x25\x7f\x46\x3e\x4d\x65\x38\x78\x8a\xfc\xc4\xbd\xf9\xc8\xfd\x06\x4e\xf7\xfa\x9f\x03\x3b\x08\x8b\xc7\x49\x0e\x40\x0e\xd0\xfd\x7f\xa2\x25\x20\x4b\xbe\xfe\x22\x72\xf9\x59\xee\xef\x3e\x49\xf7\x5d\x09\x1f\x6d\x9c\x8d\xef\xed\xbe\x58\x1f\xdf\x80\xe5\x1e\x8b\xdf\xcf\x17\x4b\xe6\x77\x46\x7e\xb9\x5e\x46\x31\x4e\x71\xc6\x53\xc6\xd8\x3e\x36\x81\x71\x34\x78\xd7\x6b\x58\x6f\x93\x90\xff\x1c\xf2\x69\x1b\xcf\x72\xf2\x28\xe2\x61\x91\xa0\xdd\x7f\xa3\x5e\x72\x86\xca\x59\xd4\x37\xf5\x66\x98\x7f\x05\x07\xcf\x1c\x57\xdc\x9b\xea\x65\x9c\x9f\x97\xc1\xff\xf0\x9d\x42\x05\x3c\x7b\x16\xef\x6d\x2a\x14\xfe\x8e\xf7\xdc\x15\xf0\xfd\xb0\xc0\xbe\x66\xc8\x2d\x44\xb0\x2a\xf4\xfa\xab\x3e\xef\x66\xfd\x83\xb6\xe7\xc9\x9c\xb0\xeb\x4b\x16\x66\xe7\x79\x83\x1e\x3c\xcf\x38\xdf\xf0\x73\xb6\xd1\x3f\xd7\x7e\x1d\xb4\xdb\x4d\x0f\x18\xfb\x09\xf1\xd0\x1b\x01\x8c\x2f\xbe\x4b\x72\x1f\xff\x97\xd8\xf8\xbf\xec\xe0\x33\x0b\x9e\xe7\x8c\xed\x77\x79\xfd\x25\xcf\xfa\xdf\x7d\x57\xfc\x44\xe3\x79\x89\xc5\xf3\x8a\x2d\x9e\x53\xe0\xeb\x6e\xf1\xec\xbc\x6f\x99\x61\xeb\xa4\xb9\x78\xbd\xd3\x78\x08\x88\x6f\xfc\x44\xc7\xcd\x38\xb7\xe8\xcb\x38\xcf\x60\x3c\xcc\x73\x0d\x3f\x8f\xe1\xfd\xaf\xf1\x5d\x9d\x6b\x9e\x98\x57\x9f\xb7\xb0\x4f\x24\x7f\x11\x71\x8a\xfd\xfc\x06\xca\xec\x40\x89\xf9\xff\xdc\x8f\x6d\xdc\x47\xe5\x18\xd3\x79\xd1\xfc\xbe\x71\xd7\x3e\x3e\x87\xec\xdf\x2b\x18\xdf\x1b\xf0\x73\x57\xb8\x95\xee\x3f\xad\xeb\x43\xf5\x7d\x84\xf3\xdc\x35\x20\x4b\xf3\xdc\x15\xa7\x73\x57\x09\xe7\x65\xb4\x9f\x2d\xe1\xfb\x96\xd2\xdb\xe0\x47\x8d\xf7\xfb\xb7\xf7\xc0\x13\x82\xf8\x3e\xa0\xd1\x77\x66\xc5\x12\xf8\x68\x69\xa7\xe9\x73\x9f\x9a\xe7\x18\xfc\x61\xdb\x95\x3f\xa8\xce\x0b\xc5\xd2\x22\xc6\x61\xc2\x91\xbf\xbc\xef\x8b\xf9\x79\xec\x00\x7b\x7f\x82\x76\xa3\x77\xda\xee\xaf\xde\xf6\x39\x6f\xa2\xd6\xdf\x10\xde\x73\x39\x78\xf6\x04\x78\xf6\xbf\x57\x7d\xb6\xfc\x80\xef\x28\xf0\x3e\x92\xf3\x8c\x62\xd4\x78\x4f\x89\xfc\x88\xef\x54\xcd\xf3\xa1\xf7\xf7\x25\xa9\x32\xce\x2b\x51\xfa\x0e\xc3\x7c\xaf\xf8\x5a\xd3\xe7\xc5\x75\x8f\xf3\xa2\x8a\x77\xbe\x61\xfd\x0e\x8b\xf1\x4e\xf3\x3b\x2f\xdc\xd3\x95\x69\x3d\x98\x79\xfb\x6f\x1d\xf9\x62\x4d\xc9\x2b\xec\x79\xcb\xf8\x2e\x76\x16\x65\xd8\xf7\x37\x96\xfc\x45\xbf\xda\x90\x7d\xe8\xe0\x1b\x18\x97\x0d\x9b\xbe\x79\xee\x5a\x57\xce\x03\xe5\x17\x4d\x7c\xe8\x20\xf1\x2c\x6b\xbe\xaa\xc5\x97\x5e\x1e\x77\xd8\x93\xe3\x7b\x99\xe6\x2f\x20\x9e\x7d\x8f\xe3\xe8\xeb\xf2\xbc\xf6\xde\xc5\x11\xbe\xeb\xde\xe3\xf7\x4b\x66\x7c\xbd\xe5\x12\x5f\x8d\xcf\x61\xd7\xef\xf0\x1c\xb6\xb9\x47\x7d\x7e\xfe\x5c\xf7\xd0\x37\xe2\x3b\xd8\xea\x3c\x77\x6e\x78\x9c\xaf\xcc\x38\xdf\x72\x89\xf3\xef\x36\x19\xe7\x1b\xd8\xbf\x69\xfc\x13\x3e\xbc\x87\x47\x7f\x1a\x9f\x63\x5a\x11\x87\x78\x8f\xbe\x42\xf3\x9a\x68\xa1\x83\x55\xa5\x0f\xed\xe0\x7b\x28\xe3\x7b\x2c\xca\x0a\x96\xf3\x4d\xa0\xa7\xde\x3f\x8a\x37\x7c\x7f\x76\x15\x71\xf5\x94\xf9\xdd\xae\xcf\xf2\xde\x2c\xfd\x44\x81\xf1\x17\x3b\xff\x48\x95\xdf\x44\xfc\xcc\x43\xbe\x01\xf9\x39\xc8\xc6\x7d\xd7\x82\x8d\xe7\x65\xa3\x97\x58\xfc\xbd\xe8\x18\xcf\x39\xe5\x78\xce\x37\xc8\x1b\x83\x2e\x79\xe3\x21\xd8\x5b\x70\xc9\x1b\x1f\x69\x90\x37\xfa\x1c\x79\x63\xc2\x23\x6f\x54\xf0\xf7\x19\x4e\x7e\x49\xdf\x6f\x8c\xe0\x3b\x98\xb0\x46\xdf\x77\x8c\x7c\x12\xb2\xbf\xcb\x36\x6e\x4e\x3f\xef\x6d\xe0\x67\xe7\x6d\xe5\x37\x93\x8f\xfd\x87\x94\x5b\x2c\xff\x09\x7c\x03\xa2\xc2\x34\xf0\xb0\x50\x13\x58\x87\x11\x97\x16\xac\x4b\x81\xd5\xf4\xe6\xf1\xef\xda\xef\xd3\xa3\xc7\xc5\xbb\x96\x6f\xa5\x55\xbf\xaf\xc8\xf7\x22\x01\xd1\xb1\xcf\x8e\x7f\x1c\xf8\x0e\xc3\xbf\xe5\x23\x7c\x35\x68\xc7\xbf\x0c\xbc\xcb\x6f\xc7\xa3\xc0\x47\x59\xfd\x1f\x6a\x84\x87\x42\x76\xfc\x2f\x80\xc7\x98\xdd\x67\xe1\xcf\x04\xab\x7f\x12\x78\x8e\xe1\x9f\x01\x7e\x81\xf9\xf3\x2d\xb4\xbf\xc4\xea\x5f\x01\xbe\xc6\xf0\xbf\x93\xfe\xb7\x8a\x18\xf3\xff\x20\xf0\x49\x86\xff\xa7\x46\xb8\x60\xed\xfc\x19\xf0\x18\xc3\x5f\xf5\x13\x3e\xc8\xf0\x37\x51\x7f\x81\xf9\xdf\x2f\xe5\xfd\xe2\x3a\xc3\x7f\xd7\x47\xf8\x78\xc0\x8e\xb7\x01\x9f\x67\xf8\xbf\x68\x84\x2f\x32\xfc\x0f\x81\xaf\x32\xfc\x4f\xd0\xce\x26\xf3\xf3\x37\x81\x87\x5a\xed\xf8\x2f\xc0\xcf\x55\x3b\x2c\x4e\xa2\x3e\xff\x63\xb9\x1f\xc1\xee\x28\xc3\x9f\x97\xed\x1c\x10\xe3\x2d\x76\xfc\x0c\xf0\xb9\x36\x3b\xfe\x59\xe0\x1b\x0c\xff\x2b\x1f\xe1\x9b\x9c\x0f\x03\xcf\xb1\xf6\x7f\x0b\xed\x2c\xb2\x76\x1e\x34\xda\x67\xf3\x7e\x0e\xed\xec\x32\xfc\x3c\xea\x47\x58\x3b\x9f\x07\x2e\x7e\xc6\x8e\x3f\x02\x7c\x67\xbf\x1d\xbf\xcf\xaf\xf6\x3f\x02\x3c\xc7\xf0\x15\xf8\x73\x81\xe1\x49\xe0\xdb\x0c\xff\x86\xc4\xdb\xc5\x2a\x1b\x87\xef\x68\x84\x77\xb2\x7e\xbd\xe2\x27\x9c\xc7\xff\xd7\xd1\x0e\x8f\x93\x27\x81\x6f\x31\xdc\x07\x7c\x87\xe1\x5f\x05\xbe\xcb\xfc\xfc\x6b\xf8\x33\xca\xf2\xf1\x2b\xc0\x27\x18\xfe\x8c\xdc\x37\x23\x82\xff\xee\x93\xf8\x21\x07\xfe\xb6\xdc\x37\xdf\xe7\xc0\x3f\x2a\xf1\xf7\x3b\xf0\xcf\xcb\x76\x0e\x38\xf0\x0f\x4b\x7c\xbf\x03\xff\xa6\xcc\xe7\x7e\x07\x7e\x42\xe2\x41\x07\x1e\x94\x78\xbb\x03\x5f\x95\x78\xc0\x81\x0f\x4a\xbc\xd5\x81\x6f\x49\xff\x7f\xce\x81\xdf\x2f\xf1\x83\x0e\x3c\x2c\xfd\xbf\xdb\x81\xb7\x49\x1e\x19\x76\xe0\xbf\x26\xdb\xb9\xcb\x81\x0f\xa1\xac\x4d\x4f\x8d\x79\xc5\x99\x9c\x63\xf2\x9a\x45\x3e\x26\x84\xd8\x0d\x98\x72\x6d\xab\x5f\x08\xd9\x9f\x5b\xdb\x8f\xb3\xf6\xe3\xac\xfd\x9a\xbc\xc4\xda\x9f\xf3\xdb\xe5\xed\xa0\xdd\x5e\xa7\xc5\xde\xa7\x98\xbd\x5a\xfd\xeb\x4c\xee\xd2\xec\xf2\x96\xdf\xde\x5e\xdc\xd2\x9f\xda\x5c\x4d\xb2\xfe\x6d\x33\x7b\x91\x56\xbb\xbc\xd6\x66\xca\x9f\x10\x02\x7f\xa5\x65\xda\x1b\x64\xf6\x57\x99\xbc\xde\x6a\xb7\x3f\xd7\x66\xb7\xbf\xda\x66\xb7\x77\x7d\xbf\xbd\xfe\xd0\x01\xbb\xfd\x71\x66\x7f\x83\xd9\xeb\xf4\xd9\xe5\xf9\x16\x7b\x7b\xeb\x21\x36\x1e\xad\xf6\xf6\xc5\x83\x85\xcc\x5c\x41\x64\x67\x32\x85\xdc\xcc\xd9\xc7\x33\xc9\xe4\xd4\x99\x4c\x21\x99\xce\x67\x93\xa9\x74\x3a\x93\x2b\x88\x07\x67\x32\xa7\xeb\x8f\x3f\xcc\x9f\x5a\x14\x0b\xe9\x5c\xf2\x7c\x7f\x32\x7d\xf6\xcc\x99\x4c\xba\x20\xb2\x6a\xd8\xde\x9c\xea\xa1\xf2\x09\xb7\xd3\xab\xb6\xd3\xeb\x65\xa7\xd7\xd5\x8e\xf9\xe4\xc9\x54\x2e\x5f\xc3\x0a\x33\xa9\x74\x66\x26\x99\x2f\xa4\x0a\xe7\xf2\x22\x79\x3e\x33\x93\x9f\x3a\x7b\xc6\x66\x2c\x9f\x29\xc8\xe7\x19\xde\x9c\xf9\xc0\x5a\x3d\x7d\xfa\x6c\xde\x51\x95\xc0\xe4\xe9\xa9\x74\xe6\x4c\xed\x69\xbe\x30\x53\x48\x3d\x2e\x1e\xcc\xeb\x4f\xd6\xca\xe3\xc7\x8e\xf5\x27\x3f\x52\x2b\x7a\x92\x87\x65\xd9\x9b\xec\x91\xe5\x11\x94\x3d\x28\xfb\x93\x71\x82\xe3\xa8\x1d\x27\xb4\x27\x0e\x2d\x3c\xee\xc1\x73\x03\x3f\x4a\xd5\x0e\x1f\x45\xb5\xa3\xa8\x76\x14\xd5\x8e\xd6\x47\x24\x99\x39\x9f\x39\x53\x48\x4e\xe5\xce\xf7\x13\x86\x21\xcb\x9f\x4d\x67\x2d\x68\xe1\x5c\xee\x74\x26\x37\xf5\x05\x82\x8e\x1f\x3b\xd6\x97\xec\xa7\xb6\xfb\x61\x0a\xe8\x61\xc0\x86\xdc\x03\x99\xca\x1e\x94\xfd\xc9\x3e\x7a\xdc\x07\xed\x3e\x68\x41\xee\xc1\x63\x2a\x8f\xa0\xec\xa9\x95\x4e\xc7\x7b\x95\x8e\xf7\x3a\x1d\xef\xa5\x46\x7b\xc9\x16\xa4\x23\x28\x0f\x03\xee\xe9\x85\x29\xe3\xf9\x11\x0c\xfd\x11\xb8\x84\xf2\xc8\x61\xe0\x87\x81\xa3\x3c\xd2\x83\xf6\x7a\xf0\x1c\x72\x0f\x64\x2a\xfb\x93\x87\x63\xd4\xb5\xc3\x31\x71\xe7\xbf\xe7\xf1\x5e\x80\xff\x26\xe9\x33\x5c\xf1\x47\xec\x21\xa3\x07\x52\xae\xfd\xc7\x8e\x0f\xf8\xff\x5e\x70\xfe\x18\xdd\x11\xef\x34\xd0\x9f\x63\x38\xa3\x2d\xe2\x26\xb1\x7f\xc7\xaf\x8b\xfe\xbc\x43\xe0\xb5\x9c\x88\xe2\x7d\x93\xa1\x6f\xe0\xff\xec\x62\xff\x75\x94\xeb\x21\x6f\xfb\xff\xe4\x62\x3f\x0e\xfb\xa3\x16\xfb\x01\x85\xfd\x4b\x2e\xf6\xe7\xd0\x68\xa3\xfe\x3f\xef\x62\x7f\x5e\xd1\xff\xa0\xc2\xfe\x39\x17\xfb\x3b\x68\x74\xb4\x41\xff\xa7\x5d\xec\x2f\xc2\xfe\x84\xc5\x7e\xab\xc2\xfe\xaf\x68\x2e\xf3\x0f\x8a\xb4\xd9\xe6\x6d\xff\x61\x4d\x6d\x7f\x17\xf6\x17\x2c\xf6\xf7\x2b\xec\x17\x5c\xec\x47\x3e\x40\x65\xee\x80\xb7\xfd\x9c\x8b\xfd\xf9\x8f\x51\x79\xc1\x62\xff\x80\xc2\xfe\x7d\x2e\xe3\x3f\xdf\x0d\xff\x5b\xbd\xed\x7f\xd0\x65\xfc\x43\x58\xbf\xd6\xf1\x6f\x57\xd8\xff\x8e\x8f\xec\xf3\x1c\x10\xc1\xfb\x65\xce\xc0\xf9\xfa\xfd\xa0\x8b\xfe\x03\x4d\xea\xff\xc0\x45\x7f\xb0\x49\xfd\x7e\x17\xfd\xb1\x26\xf5\x03\x7e\xb5\xfe\x64\x93\xfa\xc7\x5d\xec\x3f\xdd\xa4\x7e\x87\xa6\xd6\x5f\x68\x52\xff\x59\x17\xfd\x95\x5e\x75\x7d\x9e\xbf\xef\x77\xd1\x5f\x75\xd1\xe7\xf2\x57\x70\x0f\xcb\x7f\x6b\xd0\x9f\xdb\x67\xda\xed\xb3\xc4\x9f\x71\x42\xfc\xff\x00\x00\x00\xff\xff\x53\xce\x24\x51\xd8\x47\x00\x00")

func tcptracerEbpfOBytes() ([]byte, error) {
	return bindataRead(
		_tcptracerEbpfO,
		"tcptracer-ebpf.o",
	)
}

func tcptracerEbpfO() (*asset, error) {
	bytes, err := tcptracerEbpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tcptracer-ebpf.o", size: 18392, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"tcptracer-ebpf.o": tcptracerEbpfO,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"tcptracer-ebpf.o": &bintree{tcptracerEbpfO, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
