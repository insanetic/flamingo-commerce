// Package graphql Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// schema.graphql
package graphql

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
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x14\xec\x25\x0b\x14\xfd\x00\xdd\x12\xa7\x29\x82\x8d\x17\xe9\xca\xed\x65\x11\x04\x63\x6a\x2c\x13\xa6\x48\x95\x1c\x65\x23\x14\xfd\xf7\x42\x24\x25\x93\xa2\x64\xfb\xd8\xc3\xee\x69\x33\xef\x71\x66\x38\x9a\x79\x43\x73\x49\xa8\xf7\xc0\x30\x5b\xab\xba\x46\xcd\xf0\xed\x45\xab\xb2\x65\x94\xfd\xb3\xca\xb2\x2c\xdb\x81\xc1\x07\x20\xc8\x4f\x84\x7b\x30\x9c\x79\x56\x0f\xdd\x58\x22\x21\x18\xd4\x13\xaa\x67\x6d\x47\xcc\x71\x4d\x83\x8c\xef\x39\x03\xe2\x4a\x9a\x94\x5f\x44\xb8\x3b\xc3\x4d\x01\x02\x61\x27\x30\xcf\xee\x95\x12\x08\xd2\x3b\xf3\xe6\xf9\xd0\xc3\x21\x9f\x64\xd7\x60\x9e\x15\xa4\xb9\xac\x9c\xa5\x42\x7a\x2a\x51\x12\xdf\x73\xd4\x31\x74\x00\xb3\xc1\x92\xc3\x6d\xa5\x55\xdb\x8c\xd8\x2f\x59\x6b\xa0\x3a\xb9\xf9\x3c\xc9\xa7\x42\xba\xf2\xd8\x34\x57\x7b\xec\x66\xf5\xef\x6a\xd5\xe7\x79\x82\x0b\x5e\x37\x02\x87\xef\x62\xff\xa8\x51\x92\xf9\xf9\xcd\xfe\xbf\xdf\x6c\x5a\x70\xff\x69\x88\x53\x5f\x8c\xe0\x5f\x98\x3d\x10\x69\xbe\x6b\x09\xcd\x40\x99\x86\xbb\x1b\x19\xae\x8e\x07\xa5\xe9\x01\x0d\xd3\xbc\xe9\xeb\x1e\x17\xa3\x0c\x81\x24\x58\xdd\x67\x1e\xa5\xf2\x7d\xfe\x72\xaf\x2b\xc7\x07\x7d\x44\x7a\x11\xc0\x70\xad\xca\xc9\x27\xd1\x48\xc0\x05\x6a\x87\x4c\x22\x0d\x60\x71\x6c\xf3\x24\x8d\x01\xfc\x0a\x75\x7c\xd2\xa2\x4c\x23\x10\x96\x77\xd4\x43\x5b\x5e\xa3\xb5\xb6\x4d\x39\x63\x7d\xe7\x86\xef\x04\x3e\x6a\x55\xe7\x89\x75\xab\x46\xae\x35\xaf\x81\xb0\x52\x9a\xbb\x52\x9f\x6e\xee\xed\x9d\x6b\xff\x9b\x57\x4b\xde\x00\x97\x03\x10\xb4\xc0\x84\xeb\x33\x1e\x8c\xaa\xaf\xc5\x06\x9a\x86\xcb\x2a\xcf\xbe\xfb\x6b\xf9\x62\x1a\x52\xec\xf8\x8c\xef\x28\xf2\xf8\xc2\x47\xec\x7e\x28\x5d\x9a\xf0\x84\x1b\xa5\xaf\xf8\xc3\xd6\x67\x6c\xdc\xa4\xe7\x92\xc9\xf5\x4d\x67\xbb\x64\xeb\x3a\xcf\x39\x3d\xdb\x3c\x81\x34\xbc\x68\xce\x30\x4b\xbb\xde\xda\x9f\xe4\x5e\x4d\xb9\x4f\xa6\x2f\xbf\xfd\xef\x38\x62\x96\xd3\x68\x2c\x50\x20\x23\x2c\xff\x02\xcd\x41\x92\xed\x86\x20\xa2\xed\xc7\x2c\x5f\xee\xc2\x33\x4d\x18\xa4\x71\xf7\x0e\x5c\xf4\xea\x61\x93\x30\x73\x0e\xc7\xec\xbd\x53\x77\xf0\x59\x75\x20\xa8\x1b\xc1\xf4\xd6\x53\xc6\xe2\x17\x98\x12\x87\xe1\x4f\x94\xac\xc4\x3d\xb4\x82\xa2\x50\x9c\xe1\xa0\x9e\x0f\xdc\x30\xd5\x4a\xc2\x72\xa2\x57\x65\x00\xcc\x1d\x1d\xf0\x2d\x7e\x50\x1c\xb1\xe6\xf2\x45\x71\x49\x66\xab\x8a\x06\x25\xe5\xd9\xa3\x50\x40\x1e\x84\x8f\x65\x90\x29\x49\xd6\x5d\x1c\x70\xed\xcc\xb3\xed\x78\x82\x7d\x05\x58\x6b\x48\xd5\xa8\x7f\x8f\x64\xd6\x41\x07\x90\x12\x45\x2a\x2d\x42\x31\x10\x81\x6d\xa9\xe8\xf1\x02\xf2\x01\xad\x9e\x9b\x99\x1e\x88\xd8\x36\x9d\x9b\xd7\xeb\x5c\x5b\x72\x2c\xe7\x61\xba\x28\xc9\xa9\xca\xf9\x90\xbf\x49\xd2\xdd\xb5\x21\x2d\xd9\x87\x14\xb0\x0b\x45\xc3\x4a\x1c\x88\x16\x23\xc5\x58\xf4\xea\x37\xab\xf7\x75\x71\x3f\x4f\xb4\x74\x30\xf7\x62\x3a\x1a\x81\x11\x7f\x47\x3f\xf2\xe7\x85\x02\xe2\xd9\xbc\x62\x34\x45\x30\x49\x73\xfc\xe9\xa4\x9d\xb9\xbb\x95\x91\x68\x12\xd3\x95\xc8\x6b\xdc\x3a\x28\x34\xfb\x47\xc0\x94\x1d\xec\xf2\x78\xa1\xed\x51\xa3\x64\x57\xb4\xec\x69\x9b\xfb\xbc\xc6\x07\xc0\x17\xec\x92\x15\x10\xbe\x0e\x92\x42\x8c\xae\x3c\xf9\x00\x66\x34\xdd\x1e\xb1\x9b\x79\xfa\x0c\x2f\x9f\x45\xde\x62\x8c\xe4\xa4\xb9\xef\xbe\x60\xd7\x9f\x0f\xb3\xfe\x7c\x21\xcf\x8b\x65\x19\x54\x23\xd1\x84\xde\xf2\x9c\x0e\xc2\xcc\x6c\xb4\x92\x53\xac\x29\x49\xd4\x78\x8d\x2f\xc6\x6c\x80\x0e\xb1\x45\xda\x27\x4b\xcc\xd1\x56\x36\x17\x7c\x2f\xde\x78\xba\x29\x16\xf6\xc2\x05\xed\xbf\x20\xfd\x6e\x50\xef\xc1\x60\xa4\xeb\x27\xf3\x5d\xdd\x1f\x5c\x00\xff\x94\x9c\xa2\x75\x3b\xbb\x9e\xfc\x13\xa8\x6e\x80\x57\xf2\x5b\x2b\x30\xe9\xe2\x12\x65\xb7\x51\x1a\x87\xc3\x26\x3e\xfb\xe9\xf2\xaa\x71\xc3\x07\x1f\x6b\x01\xc6\x04\xcf\xa7\xa5\xe2\xbe\x15\x08\x9a\x1d\xbe\xa1\x69\xc5\xb0\x88\x1a\x07\xcd\x4d\x92\xcf\xf3\x53\xf6\x08\x0c\xc9\x64\xc6\x9e\x2e\x55\x0d\x5c\xfe\x6a\x6d\x6b\x25\xfa\x97\x0c\x57\x2e\x65\xd3\x56\x15\x1a\xff\xb3\xe8\xe4\xce\x45\x7d\x2b\x46\xd4\x3b\x76\xfe\x36\x18\xfd\xf2\xf1\xe4\xde\x6a\x3b\x14\x3f\x08\x65\x69\x75\x2a\xfb\xa3\xc5\x51\xfd\xa7\xc9\xde\xba\x07\x51\x93\xbc\xca\x67\xc6\x77\xd6\x81\x8f\x7c\x6b\x7c\x8d\xfe\x6e\xd1\x50\x9a\x98\x07\x66\xbc\x46\xd5\xed\x73\xff\x2f\x00\x00\xff\xff\x54\xd2\x74\xde\x53\x10\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
