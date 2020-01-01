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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x96\xdf\x6e\x1a\x3b\x10\xc6\xef\x79\x8a\x41\xb9\x49\x8e\x94\xf3\x00\x7b\x79\x38\xfd\x93\x8b\x2a\x34\xb4\xe9\x45\x55\xa1\x89\x3d\xb0\x56\xbd\xf6\x6a\x3c\x1b\xb2\xaa\xf2\xee\x95\xbd\x06\x96\x10\x28\x21\xb4\xdd\x0b\x24\xd6\xf6\xf7\xfd\x66\x3c\x5e\xcf\x19\x8c\x7c\x55\x11\x2b\x9a\x8e\x4a\x52\xdf\x7d\x23\xd3\xb1\x45\x45\xd7\xac\x89\x47\xde\x09\x3d\x08\x30\xd5\x4c\x81\x9c\x04\x90\x92\x80\x29\x34\x56\xc0\xcf\xd2\x3f\xd5\x30\x93\x13\x38\xe7\xc6\x39\xe3\xe6\x17\x50\x47\x01\xf0\x51\x01\xaa\x46\x50\x8c\x77\x03\x69\x6b\x3a\xc8\xec\xc7\x00\x00\xe0\x0c\x3e\x95\x04\x23\x64\x01\x29\x51\xc0\x04\x98\x7b\xe3\xe6\x20\x1e\xee\xa8\xb3\xd0\x69\xa6\x42\x96\x62\xad\xfc\x3f\x29\xcf\x28\xa4\xe3\xda\x9e\x54\xb7\x22\x53\x19\x07\x0a\xc3\x92\xd1\x04\x40\xcb\x84\xba\xed\xeb\xa6\xb1\x2b\x37\xf3\xa1\xd8\xc5\xad\xaf\x57\x73\xb2\xd3\x44\x50\x08\x34\xd5\xe4\x74\xa4\xf5\x2e\xe5\x28\xa4\xd7\x7e\x06\x35\xb6\x55\x4c\x16\x3a\xbd\x91\xa6\xcb\x3c\xa5\xc2\x16\x94\x77\x82\xc6\x01\x6a\x6d\x62\xea\xd0\x82\x59\x59\xa4\x69\x3b\x81\x12\x4f\x62\x98\xa6\xdf\xc1\xe3\x60\x30\xd8\xb9\xc7\x3d\x7c\xb8\xec\x3c\x00\xef\x7c\x23\x89\xb9\x9f\xb0\x38\x2e\x6d\x6d\x14\x5a\xdb\x42\x28\xfd\xc2\xc5\xd0\x10\x42\xa3\x88\x42\x80\x1a\xe7\xb4\x77\x8b\xfb\x5e\xdd\x0e\xe7\x54\xe4\x0c\xe7\xe7\xeb\xbe\xc0\xc6\xeb\x15\xc3\x6f\x9d\xc6\x13\xe9\xe2\x89\x06\xf2\x96\x7d\x5e\x49\x15\x1a\xbb\xb2\xcd\xcf\x44\xd8\xb8\xf9\x30\x66\x2d\xc5\xb2\x37\xcd\x3d\x9a\x1c\xd1\x1c\x85\x16\xd8\x16\x5b\x7a\xbd\x70\xc7\xec\xef\x8d\x26\x2e\x36\x06\x2b\x92\xd2\xeb\x62\x9b\x24\x55\x95\x62\xd2\x46\x46\xc8\x3a\x99\xc1\x3f\x29\xec\xb4\x31\xff\x8e\x36\xc6\xd2\x7c\xac\x7c\xe3\xa4\x27\xb6\x8a\x62\xcc\x46\x51\x27\x2a\x46\x2c\x15\xcf\xc7\x6e\x9c\x10\xcf\x62\x71\x1e\x58\x66\x39\xfe\x99\x71\x68\x0b\xf8\xcf\x7b\x4b\xe8\x52\xed\xfd\xf2\xd4\xf7\x54\xa6\x5f\xd0\x08\x98\xaa\xb6\x54\xa5\x6f\xcd\x6b\xdd\x5f\x64\x3e\x69\x94\x8a\x95\xfc\xb7\xfc\xdf\xa2\xa0\x7d\xc3\xec\xf9\x64\x08\xa9\xce\xa3\xe2\xb2\xd8\x5e\x9e\x94\xd2\x2f\xae\x66\x8c\x15\x9d\x14\xaa\x61\xfb\x2a\xa4\xf7\x52\xd9\x93\x02\x95\x52\x1d\x4f\x74\x43\xda\x30\xa9\xd3\x95\xee\x33\x29\x7a\x19\xd1\x08\x9d\x22\x6b\x49\x9f\x14\x49\x75\xaa\xe9\x3a\xbf\x21\x0c\xde\x1d\x7c\x0f\x2d\x89\x7a\x6b\x8f\xfa\xd0\x3c\xa3\x93\x79\x39\x03\x1d\xb9\x87\xdb\xba\xd3\xfc\x7d\x3f\xf6\x48\xfe\x49\xd4\x5b\xb4\x46\xa7\x17\xbf\x9d\x36\xbe\xba\x5f\xd9\xdd\xa4\x4e\xb0\x5f\x06\xf1\xc6\xbd\x7d\x32\x9e\x0a\x98\x1e\x84\x9c\x86\x14\xea\xc7\x86\xb8\x5d\x35\x7a\xef\x28\xf7\x95\x16\x83\x6c\x74\x45\x2a\xb7\x84\x0b\x23\x25\x18\x09\xb9\x49\x5a\x77\x43\x07\xf4\x93\xfb\x6b\x34\x4f\x1a\x6e\x21\x7e\xc8\x7d\xeb\x92\xf2\xda\xd9\x16\x6a\x1f\x82\xb9\xb3\x04\x66\xd6\x75\x7a\x15\x06\x55\x1a\x47\xe0\xbc\x00\x2a\x31\xf7\x11\x3d\xb6\x97\xd8\x9d\xa0\x8e\x78\x07\xeb\x44\x90\x65\xcd\x72\xce\x24\x0d\xbb\xcf\xeb\x93\x3f\xbc\x58\x9d\xc0\xdc\x0a\xc0\xb8\x87\x90\xb3\x11\x92\x7b\xb2\xdb\x61\xd4\x6d\xee\xda\xa9\xa7\xfa\x38\xf8\x19\x00\x00\xff\xff\x56\xf6\x3f\x9a\x03\x0c\x00\x00")

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
