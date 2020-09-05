// Code generated by go-bindata.
// sources:
// tmpl/monitor.tmpl
// tmpl/screenboard.tmpl
// tmpl/timeboard.tmpl
// DO NOT EDIT!

package main

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

var _tmplMonitorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xcf\x6e\xe2\x30\x10\xc6\xef\x79\x8a\x51\xce\xbb\x3c\x41\x39\xac\x5a\x56\xe5\xb0\xa0\xad\x90\x7a\x58\xad\x2c\x2b\x1e\x88\xd5\x60\x53\xc7\x21\x42\x5e\xbf\xfb\xca\xff\x20\x4e\x4d\xcb\x29\xfe\xbe\x6f\x7e\x99\x4c\x26\x28\xec\xe5\xa0\x1a\x84\x9a\x51\x4d\x99\x3c\x90\xa3\x14\x5c\x4b\x55\x43\xcd\x18\x31\x06\x16\x6b\x06\xd6\xd6\x60\x2a\x00\x41\x8f\x08\xf9\x6f\x09\xb5\x0b\x6d\x9c\x63\x6d\x5d\x01\xe8\xcb\xe9\x4e\x68\xe7\x9c\x10\x32\xe6\x3b\xf0\x3d\x2c\x76\xf4\xd0\x83\xb5\xae\xcc\x5d\xcd\xcb\xfe\x18\xa3\xa8\x38\x60\x08\x5a\x5b\x1b\xb3\xb0\xb6\xfe\x66\x0c\x0a\x66\xed\xdf\x48\x42\xc1\x02\xe4\x88\x7d\x4f\x0f\x98\x43\x1e\x1e\x56\xdb\x5d\xe5\x1a\xf8\x15\x6d\x6b\x2b\x27\x01\x60\xdf\xd0\x8e\x6a\x2e\x05\x49\xa5\xd3\xfc\xf6\xe4\xac\x7e\xb1\xba\xc6\x66\x84\x0a\xe0\x7d\x40\x75\x81\x25\xb8\xfc\x6f\x7f\xfd\xcf\x63\x4f\xf8\xd8\x52\x45\x1b\x8d\xca\x3f\x61\x6c\x75\xe4\xba\xbd\x82\x43\xd3\x69\x16\x1b\xa9\xf9\xfe\xb2\x91\x4f\x54\xd3\xe0\x08\xaf\x10\x21\x89\x7b\x3b\xe1\x69\xfc\xb4\x3f\x24\xb3\x29\x24\xe0\x0b\x06\xc0\x5a\x68\x54\x67\xda\x05\x57\x45\x95\xf0\x24\x07\x68\x39\x5d\x04\x87\xfb\xff\x18\x18\xd7\x59\xa3\xd4\x2b\xd3\x1e\x27\x99\x22\x69\xc7\x8f\x28\x07\xfd\x1c\x77\x20\x9c\x48\x1b\x19\xb9\x5b\x04\xac\x45\xd3\x0d\x0c\x6f\x7b\xc4\x83\x40\xfc\x3e\x05\xcc\x34\xf3\xe9\xb8\xde\x07\xae\xf0\xe7\xd0\x75\xaf\x5c\x30\x39\xa6\x79\x79\x99\xec\x87\xae\x23\x63\x30\xd2\xc4\x8a\x05\xe5\x91\xe1\xf8\x2c\x7b\xfd\x84\x1d\xbd\xc4\x99\xe1\x48\x5a\xd9\x6b\xc2\xbc\x16\xa7\xf6\x21\x56\xa4\xad\xce\xb4\x1b\xfc\x42\x4e\x92\x78\x15\x33\x64\x31\x3b\xa1\x4e\xdf\x46\xab\xb0\x6f\x65\xc7\x26\x9b\x19\x16\x76\xee\xe8\xdb\x79\xe9\xff\x1a\x6e\x90\xed\x5b\x88\x00\xc8\xb7\xd8\xc2\x4d\xca\x9e\xe6\x56\xf3\x4a\x95\xe0\xe2\x90\xe4\x31\x1e\x43\xf5\xcc\xfc\x1c\xf1\x82\x8d\x3c\xbb\xaf\x30\x47\x11\x95\xf4\x8c\x39\x4f\xdf\x61\x3f\x2a\xae\x79\x93\x3e\x09\x80\x26\x9d\x03\x6c\x6e\x7f\x41\x99\xdf\x34\xd1\xe6\x3d\xde\xcb\x67\xf8\xe2\x8a\x94\x0e\x95\xad\xfe\x07\x00\x00\xff\xff\x41\xd6\x1b\xf0\xed\x05\x00\x00")

func tmplMonitorTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMonitorTmpl,
		"tmpl/monitor.tmpl",
	)
}

func tmplMonitorTmpl() (*asset, error) {
	bytes, err := tmplMonitorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/monitor.tmpl", size: 1517, mode: os.FileMode(420), modTime: time.Unix(1599283620, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScreenboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xd9\x52\xe3\x3c\x1a\xbd\xe7\x29\x54\xb9\x9e\xe6\x7f\x02\x2e\x68\x96\xbf\xa9\x82\x6e\x86\xd0\xf4\x2c\x35\xe5\x12\xf6\x17\x47\x85\x6d\xa5\x25\x39\x24\x64\xf2\xee\x53\xda\x25\xc7\x96\x45\xd5\xfc\xdc\xc4\x3a\xe7\xe8\x68\xdf\x61\xc0\x69\xcf\x4a\x40\x8b\x0a\x0b\x5c\xd1\xba\xe0\x25\x03\xe8\x5e\x29\x66\xd5\x02\x2d\xaa\xaa\x38\x1c\xd0\xf9\x5d\x85\x8e\xc7\x05\x3a\x9c\x21\x24\x88\x68\x00\x5d\xa0\x85\xc4\x9f\x55\xe0\x78\x5c\x9c\x21\x74\x38\x7c\x41\x64\x85\xce\x9f\x00\x57\x3f\xba\x66\x8f\x8e\xc7\x33\x84\x18\xe0\xaa\xa0\x32\x78\x81\x64\x8c\x98\x95\x71\xa0\xab\x7c\x40\x1a\x2c\xd7\x98\x81\xc1\xb8\xfe\xd6\x71\x43\x62\x34\xe6\x33\xb4\x9b\x06\x0b\x78\xc1\x8c\xe0\xd7\x06\xb8\xa7\xdf\x89\x58\x27\x05\x0c\x77\x35\xa0\x73\x0d\x08\xa3\x2b\xb6\x46\xa8\x8a\x8e\x50\x87\x5b\x90\xbf\xa6\xf8\xdf\x65\x50\x97\x1e\xa1\x0d\x83\x15\xd9\x39\xee\x51\x07\x2d\x5b\xc1\x0a\xf7\x8d\xb0\xec\xb5\x09\x6a\x7a\xb4\x40\x63\x81\xa8\x30\xbf\x48\x55\x83\x98\x2a\xc2\xbb\x62\x4d\xbe\xc5\x7e\xe3\xdb\x4c\x7e\xdb\x6c\xed\x2c\xfa\x0f\x07\xed\x2d\xf4\x4f\x07\xb9\xfa\x35\xcd\xad\x3d\x4f\x3a\xc2\x33\xec\x44\x14\xc7\x15\x61\x60\x71\xd9\x90\xba\x8b\x7c\x0a\xac\xa0\xd0\xcd\x8a\xe6\xed\x96\xe4\x23\xce\x55\xc1\x25\x12\x9a\x19\x49\xd2\xeb\x1b\x90\x7a\x2d\x2c\xba\xd6\x21\x63\xe2\xb8\xa4\xc3\x2f\x52\x89\xb5\x05\xdf\x55\xc0\xc4\xb7\x4c\xba\x30\xba\xfe\x74\x39\xe4\xb7\x2d\x41\x46\xbd\x5e\xd1\x86\x32\x0b\x96\x2a\x60\x62\x5b\x66\xa6\x1e\x5b\x08\x31\x33\x5c\x02\x54\xc8\xef\x0b\xd3\x9f\x10\x6a\xc8\x16\x0a\xbe\xc1\xae\xcd\xee\xc9\x16\x96\x32\x6c\x53\x3a\x4e\xa5\x37\x9a\x7c\x03\xd7\xb0\x1a\xcb\x41\x44\x08\xd2\x40\x51\xc1\xca\x65\xc3\x1a\xbc\x90\x0f\xab\x41\x68\x4b\x3e\x6c\xae\x34\xbe\x08\xd4\x41\xea\x41\xed\xf5\x5c\xd0\xf6\x67\x47\x84\xe7\x4a\x85\x15\xbd\x04\x6d\x55\x86\xb2\x59\xd3\xcb\x5e\x50\x5e\x62\x3f\x62\x10\xc2\x0e\x32\x8e\xa1\x66\xd6\x50\xf6\x83\x68\xe8\xe8\x6e\x32\x18\x3b\x81\x68\xd6\x71\x29\xf6\x61\xf6\x7c\xbd\x0f\x08\xae\x82\xbe\xf5\xbd\xc3\x23\x6e\x40\x88\x40\x8a\xd0\xc6\x40\x76\x22\x74\x8a\x45\x14\x3b\xca\xd0\x89\xe1\x6d\x43\x36\x23\xa6\xc5\x4a\xe2\x7a\x39\x98\x50\x26\xbc\x6f\x49\xd3\x3c\x90\x2e\xa4\x56\xa4\x69\x8a\x96\xb8\xea\xf3\x92\xac\xdc\x2a\x39\xde\x9d\x3a\xe2\x5d\xe4\x88\x77\x69\xc7\x23\xfa\xe3\x0f\x55\xc9\xa9\x16\x1b\x6d\xc2\x27\xf8\xdd\x03\xb7\xcb\x40\xdc\x8a\xe3\x5c\xb4\x4c\xc8\x3f\xa6\x65\x23\xad\xfb\xf7\x1e\xd8\x3e\x2c\xdc\x6f\x5b\x2a\xcb\x64\xd5\x92\x59\x72\x1c\x3e\xb9\x1c\xcd\xf8\xa8\x44\x87\x66\xbf\x25\x58\x84\x96\xa1\x2c\xcb\xf7\x01\x04\x23\x65\xc8\xb4\x1a\x31\x86\x8e\xcf\x2b\x2d\xec\xc4\x2d\x69\x04\xb0\xa8\xcc\x72\xa4\xae\x34\x1c\x0c\x55\x27\xcc\xb2\xbe\x27\x6d\x38\x45\xc9\x69\xb8\xf5\xd3\x93\x65\xb3\xac\x2e\xeb\x9a\x41\x8d\x05\x8d\x72\x89\x3d\x6a\x67\xa8\x50\x97\xe5\x7c\x45\xdb\x0d\x66\xf0\x4c\x43\xb2\xd4\x60\x21\xa8\x5f\x97\xbc\x2c\xcf\x77\x2d\x3b\xee\xb0\xf9\x4b\x85\x46\xed\x1f\x09\xb3\xac\x7f\xb0\x0a\xd8\xd7\xa8\xa7\x53\x09\x15\xaf\x6e\x37\xe4\x25\xf9\x8e\xd7\x84\x9d\x5a\x56\x84\x45\x9e\x5a\x94\x65\x7a\xb3\x13\x0c\x5f\xd1\x26\xe4\x40\x62\x45\x49\x1b\x6b\x1a\x88\xb2\x4c\xef\xba\x92\x01\xe6\xf0\x27\xa5\x11\x4f\x0c\x5e\xd4\x92\x30\xe6\x03\x71\x66\x7f\xe8\x2a\x22\x08\xed\x70\x73\x4b\x59\x8b\xc3\xf9\x28\x9c\xad\x46\x74\x5f\x06\xc2\x93\xa9\x4b\x76\x2c\x17\xab\x58\xa9\x68\xc1\x3c\x36\xb1\x3f\xb2\x31\x27\x77\x49\x93\x65\x4a\x2e\x7c\x59\x4b\xdf\x8c\xad\x1e\x15\xc3\x61\x69\xc7\x4f\x38\x30\x23\x65\xb6\xff\x5d\xb7\x05\x26\x62\x8e\x68\xcc\x35\xb1\x51\x64\x7b\xbe\xe0\xa6\x1f\x54\xc4\x56\x41\x76\x13\x66\xf8\xfc\x4c\xb6\xb8\x86\x9f\x4f\xf7\x83\x6c\x4a\xb4\xe8\x99\xeb\xe9\x81\x6c\xc6\x7a\x76\x77\x90\x07\x8d\x6c\x99\x12\x9b\xa6\xb1\x6d\xd3\x5f\xde\x7f\x86\xd3\xe3\xcc\x62\x3b\xe3\x16\x1d\x68\xf4\x5f\xe2\x58\x33\x5d\xfd\x83\xdd\xcd\x67\xeb\xfc\x19\xd7\x7a\x9d\x8c\x66\x0e\x81\x6b\xb3\xa0\x72\x74\x81\xfe\x7d\x38\x98\xd9\xc1\xab\x8f\xc7\xc5\xe1\x70\x7e\x3c\x2e\xfe\x76\x38\x40\x57\x1d\x8f\xff\x49\xef\xc0\xcc\x46\xe8\x53\x7b\xb0\xd1\x4d\xd9\xcd\x16\xba\x68\xfa\xf2\x9d\x64\x94\x3a\x99\xd7\x40\xaa\x82\x7e\x33\xb9\xef\x92\xd9\x56\xe2\x41\x96\xbe\x9c\x64\x72\x1c\x19\x64\xfc\x01\xb3\xb7\xa8\x9a\x7d\xc6\x47\xa9\x93\x8c\xb7\x4a\x35\xb2\x95\xfc\x7f\xed\x01\xef\xf1\x2b\x44\x0b\x60\xa3\x00\xbb\x05\x32\x6c\x96\xd5\xc9\xbc\x35\x33\x6b\x8d\xf6\x99\xb1\xca\xcf\xee\x31\xd2\xc0\x9e\x69\xcf\x46\xf8\xc4\x45\x41\x74\xe9\x21\xf7\x96\xd1\x9d\x87\x17\x24\x8f\xfd\xe1\x89\x37\x3c\xe8\x46\x47\xdc\x89\xb8\x8f\x0c\x4a\xc2\x09\x75\xc7\xa9\x8d\x03\xfc\x15\x98\x53\xcc\x5e\x7c\xc4\x77\x42\x19\xc7\xda\x09\xaf\x5b\xda\x45\x75\xb3\xa2\x5d\x5c\x37\x81\x20\x69\x74\xd9\x00\x13\x77\xd7\x16\xc6\x32\x58\x10\xb7\x13\xf2\x74\xda\xa5\x17\xf4\x09\x56\x0c\xb8\x9b\x47\x71\x2f\x68\xc1\x0c\x16\xdc\x03\x78\x59\xd2\xf1\x1e\xea\x00\x6d\x74\xc8\x76\x7f\xcb\x65\x38\x84\x95\xa4\x5d\xa2\x6a\x8a\x44\x49\xbb\xe8\x90\xa8\x4e\x63\xe3\x93\xd5\x54\x1f\xf4\xeb\x7c\xb0\xb6\x87\xcb\xfa\x44\xc4\xab\x35\x94\x6f\xee\xca\x4b\x05\xdc\x01\x40\x33\xc9\xe8\x7f\x32\xda\x6f\x48\x57\x5b\xbc\xb6\x61\x63\x12\xf0\xf3\x3e\x91\x49\xe4\x30\x7f\xf3\x56\xbe\x3d\x52\xee\x6f\xb8\xca\xb7\x62\x43\xb9\xbf\xbd\xb4\xf4\xac\xcb\x4d\x55\x43\x64\x03\x12\x08\x7c\x8c\x20\x7d\x09\xfa\xfc\xe0\x9a\x63\x2d\x5a\xd7\x1e\x0a\xff\x2f\x02\x5e\xe2\x0d\x7c\x87\xf7\x86\x74\x90\x95\xab\x30\x47\x61\x66\xe6\xe2\x7e\xad\xcb\x70\xcb\xfe\x5a\x47\xdb\x75\xcf\x26\x4d\xd4\x62\x1b\x76\x74\x35\x5b\x47\xfd\x3c\x94\x24\xbd\x96\xe4\x23\xe8\x2d\x5c\x87\x8c\x89\xe3\x92\x0e\x0f\x98\xd5\xfe\xf6\xa9\xd5\x21\x7b\xc5\x60\xb9\x74\x79\xba\xad\x2b\x49\xb7\x75\x45\x50\x68\x3a\xf3\xc0\xb6\xa4\x04\xf3\xe3\x0a\xa1\x83\x85\xf9\x75\xa5\x19\x8a\x73\xac\xcd\x5b\x48\xe4\xab\x9e\x4b\x62\xd3\xe8\xc9\x64\xba\xa6\xe1\x05\x58\xb8\xb6\xc8\x16\x2b\xb6\x06\xf3\x95\x1e\xc8\xd2\x33\x1e\xde\xd3\x5e\x0c\x3c\x1b\x05\x0e\x5d\x87\xd2\x74\x8b\xf6\x5c\x2c\xd7\xf4\xfd\x1b\xf1\x87\xda\xb6\xe7\xa2\xe0\x6b\xfa\x5e\xac\x25\x6a\xdb\x37\x56\x66\xb9\xde\x30\x46\xd9\x88\x2f\x68\x7c\xe0\xec\xd4\x59\xde\xf7\x58\x40\x57\xee\x4f\xcd\x1b\x43\x0c\xdc\xbd\x3e\xcb\xfe\x2b\x03\xfc\x56\xd1\xf7\xee\x34\x81\x57\x47\x0d\x92\x08\xe3\x64\x25\x72\x4d\xb8\x60\xe4\xb5\x17\x41\xa3\xfa\x74\xaa\x90\x1d\x24\x35\x88\x99\x95\xda\x93\x79\x2a\xbd\x27\x5c\x9c\xa6\x66\x1f\x52\x8b\x46\xd2\x83\xe4\x06\x51\x93\xc9\x5d\x13\xbe\x69\xf0\x5e\xdf\x80\x58\xb2\xd2\xa0\xbd\xe0\xb0\x8f\x88\x03\xe9\xfc\x03\xd1\x23\x83\x15\x30\xe8\xfc\x14\xa0\xa6\xd1\x62\xe3\xf1\xf0\x3e\x24\x92\xa7\x97\x0d\x52\xc1\xbf\x80\xd1\x2b\xda\x77\x7e\x28\xac\x49\x05\xc5\x07\x30\x5a\x94\x1a\xb7\x8b\xc9\x50\x3d\x33\x6d\x76\xb8\x86\xa5\xc0\xa2\xe7\xb2\x36\xa3\x57\xc8\x56\x91\x05\x57\xac\x6e\x8a\xe8\x61\x72\x2a\x72\x76\x8a\xe1\xcb\xe6\x48\x8a\xfa\xbd\x31\x7c\xad\x9b\x8a\xfc\xb9\x14\xc3\x45\x6b\x2c\xc5\x70\x09\x9b\x8a\xfc\xb9\x14\xa3\x0d\xf8\x58\x92\xd1\x7e\x7c\x32\x7a\xfa\xcc\x80\x19\x6e\x79\x88\xea\xb3\x65\x8c\x6f\x74\xe8\xe2\xe4\x95\x6f\x49\xc3\xab\x2b\x4e\xfd\xa5\x95\x61\xb2\x9e\xd0\x3c\x3a\xf9\xc8\x9a\x70\x50\x7d\xd6\xc3\xaa\x63\xfb\x41\xa3\xb9\x8c\x77\x37\x1c\x95\x44\x05\x6d\x51\x0c\x37\x61\x32\x79\x30\x1c\x1f\xf3\x7d\xdb\xf1\x60\xac\xab\xa0\x7e\x3e\xb3\xa4\xdd\xd4\x5d\xad\x31\xc3\x65\x78\xc9\x32\xb5\x98\xd2\x9a\x83\xcb\x7d\xa3\x43\x76\xf5\xb4\xdc\x48\x37\x90\x87\x5d\xfd\x5f\x09\x89\xff\x70\x38\x9e\xfd\x2f\x00\x00\xff\xff\x3d\x13\x19\x14\x93\x22\x00\x00")

func tmplScreenboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScreenboardTmpl,
		"tmpl/screenboard.tmpl",
	)
}

func tmplScreenboardTmpl() (*asset, error) {
	bytes, err := tmplScreenboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/screenboard.tmpl", size: 8851, mode: os.FileMode(420), modTime: time.Unix(1599289036, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xc9\x6e\xdb\x3c\x10\xbe\xfb\x29\x06\x42\x0e\xff\x0f\xc4\x7e\x80\x02\x3e\xa4\x09\x1c\x14\xe8\x92\x26\x41\x7a\x28\x0a\x81\x96\x46\x0a\x51\x6a\x09\x45\x25\x71\x08\xbe\x7b\xc1\xe1\x26\xdb\x8a\x7b\xa8\x4f\x9c\xf9\xbe\x59\x39\x43\x4b\xe2\xd0\x8d\xb2\x40\xc8\x4a\xa6\x58\xd9\xd5\xb9\xe2\x0d\x6e\x3b\x26\xcb\x0c\xb2\xb2\xcc\xb5\x86\xd5\xa7\x12\x8c\xc9\x40\x2f\x00\x14\x57\x02\xc1\xfd\xd6\x90\x59\xf4\x9e\x54\xc6\x64\x0b\x80\x12\x87\x42\xf2\x5e\xf1\xae\x0d\xf0\xd5\x44\xe5\x48\x12\x59\x99\x77\xad\xd8\x91\x0f\xcb\xb9\x45\x56\x7e\xb3\x8a\xa5\x31\x0b\x00\xad\x5f\xb8\x7a\x84\xd5\xb5\x64\xfd\xe3\x10\x95\x92\xb5\x35\xc2\x0a\x48\xac\x2d\x46\x29\x85\xa4\x66\xd2\x49\xae\xae\xb0\xe2\x2d\xa7\x24\x9c\x3b\x80\x67\xfe\x96\x8a\x78\xe0\x6f\x16\x08\x46\x4b\xe0\x15\xac\x2e\x46\xd5\x0d\x05\x13\x68\x21\x16\x05\x6f\x92\x50\x63\x32\x6b\x82\x6d\xe9\x5d\x07\x07\x37\x12\x0b\x3e\xf8\xa0\x7d\x14\xbc\x83\x84\xbe\xeb\xe0\x5a\x76\x63\x4f\x1d\xa8\xed\x09\xd6\xf0\x53\xeb\xb3\xda\x69\x3f\xac\x03\xc1\x98\xd0\x9d\x33\xde\x96\xf8\x7a\x0e\x67\x28\xb0\x39\x60\xf0\xca\xc3\xc6\x9c\x6b\x4d\xc1\x32\xad\x89\x49\x27\xd2\xfc\x9a\x4f\xe4\xae\xe8\x7a\xa4\x44\x06\x7b\xf2\x89\x0c\x4e\x6b\xc3\x38\xc2\xa9\x44\x12\xe3\x9f\x12\x51\x3b\x77\x1f\xa4\x1d\x48\x5a\xfb\x39\x00\xe8\x99\x40\xa5\x70\x6f\x3e\xc9\x62\x75\xe3\x91\x70\xc7\x91\x9b\x57\x82\xf7\xb3\xdc\x8d\x05\x02\xdf\xcc\xa7\xf3\x85\xc9\xdf\x28\xa9\x31\x56\xe5\x86\x6d\x4f\xe9\xa7\xd6\xdb\x35\x04\xc5\x7c\xd5\xae\xc7\xb4\x48\x56\x48\xf9\x3d\x33\x31\xc6\x61\x7b\x20\x21\x81\x5a\x53\xf8\xcf\x6c\x8b\xc2\xc6\x11\x74\xf0\x64\xa7\x3d\x1a\xaa\x54\xc2\xf1\x21\x16\xe5\x2a\xb8\xc5\xa7\x11\x07\x35\x5b\x82\x74\x58\xac\xe1\x69\xd2\xeb\xef\x23\xca\x5d\x5a\xa4\x98\x27\x95\xb6\x34\x86\xea\x3d\x28\x57\x6b\x9b\x02\x78\xef\xd1\xe4\xa2\xae\x25\xd6\x4c\x75\xd2\x25\x61\x95\x2d\x42\x96\xc1\x7f\x57\x78\x8b\xd5\x9d\x92\xbc\xad\xa7\xbc\xff\x69\x4d\x93\x59\xd8\xd3\xa4\x89\xd1\xc8\xa1\x8d\xba\x3c\x0c\x1b\xe7\x4b\x6b\xdf\x0b\xa7\x89\xbc\xc3\x91\xb3\x96\xb4\xeb\x7e\xc0\xec\xa6\xfb\x63\xd8\xf3\x34\x7a\xfb\xfd\x4e\xc6\x3f\x78\xa9\x1e\xad\xe9\x0b\x1d\xbc\xa1\xd3\x9e\x30\x3b\xdd\xd6\x7d\x9b\x49\x9d\xbe\xf0\xd0\x8a\x08\xb8\x91\xbe\xec\xda\x92\x5e\x4a\x26\x36\x9d\x6c\x98\x1a\x60\x3a\xdc\xef\xc2\xe1\x79\x4e\x2d\x2d\x12\x35\xaf\x88\xbb\xd7\x35\x98\xb6\xed\x74\xd7\xf6\x27\xc4\x7a\x6e\x7a\x26\xa7\xb7\x7c\x99\x34\x69\xfa\x52\x51\x61\x7f\xe6\x97\xea\x38\x40\x6c\xc6\x38\xa8\xae\xf9\x58\x5f\x76\x82\x3c\x17\x24\xe7\xdb\x3a\x2f\x48\x13\xa2\x1f\xd0\xfe\xea\x71\x73\xe8\xb1\x9a\xf5\xb8\x79\xdf\xe3\xf4\xda\x1c\x30\x77\x5a\x24\xea\xf1\xae\xcf\x4b\xc6\x3f\xc0\x53\x5d\xba\xfe\x7b\x6c\x7a\xc1\x14\x3e\x30\xc9\xd9\x56\x60\x7c\xfa\x26\xff\xce\xf6\x3b\xc1\xd3\xf2\x67\xcf\xf3\x37\xdf\xb2\x06\x27\xef\xc5\x57\x2b\x86\xfb\xea\x25\x56\xfc\x15\x26\x7f\x8f\x56\x0c\x68\x89\x15\x1b\x85\x4a\x9f\x15\x4e\x74\xf0\x71\xce\x66\xf1\x27\x00\x00\xff\xff\x85\x76\x0d\xbc\xdb\x08\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 2267, mode: os.FileMode(420), modTime: time.Unix(1599283620, 0)}
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
	"tmpl/monitor.tmpl": tmplMonitorTmpl,
	"tmpl/screenboard.tmpl": tmplScreenboardTmpl,
	"tmpl/timeboard.tmpl": tmplTimeboardTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"monitor.tmpl": &bintree{tmplMonitorTmpl, map[string]*bintree{}},
		"screenboard.tmpl": &bintree{tmplScreenboardTmpl, map[string]*bintree{}},
		"timeboard.tmpl": &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
	}},
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

