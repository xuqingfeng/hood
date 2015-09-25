// Code generated by go-bindata.
// sources:
// ../config/hood.json
// ../templates/.c.tpl
// ../templates/.css.tpl
// ../templates/.gitignore.tpl
// ../templates/.go.tpl
// ../templates/.html.tpl
// ../templates/.java.tpl
// ../templates/.js.tpl
// ../templates/.php.tpl
// ../templates/.py.tpl
// DO NOT EDIT!

package hood

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _ConfigHoodJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x50\x50\x0a\xf0\x08\x88\x77\x8d\x08\x71\xf5\x0b\xf6\xf4\xf7\x53\xb2\x52\x50\xd2\x2b\xc8\x28\x50\xd2\x01\x49\xb9\x7b\x86\xc4\x7b\xba\xfb\xf9\x07\xb9\xa2\xa9\x48\xcf\x2c\xc9\x4c\xcf\xcb\x2f\x4a\x85\xaa\xf3\x47\x97\xcf\x87\x48\x78\x05\xa3\x49\x64\x15\x43\x24\x9c\x83\xd1\x65\x92\x8b\xa1\x52\x5e\x8e\x61\x8e\xe8\xba\x12\xcb\x12\x21\x92\x01\x91\x21\x1e\xfe\x7e\xe8\xee\xad\x84\x48\x7a\x84\xf8\xfa\xa0\x49\x65\x94\xe4\xe6\x40\x6d\x44\xb7\x0f\xdd\x8b\x2e\xf1\x3e\x9e\xc1\x21\x40\xb9\x68\xa0\x38\x50\x46\xcf\x25\x38\x3e\xb8\x04\xe6\x47\xa0\x40\x66\x4a\x6a\xa2\x12\x90\x19\xcb\x55\x0b\x08\x00\x00\xff\xff\x31\x6c\x5e\xc2\x39\x01\x00\x00")

func ConfigHoodJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigHoodJson,
		"../config/hood.json",
	)
}

func ConfigHoodJson() (*asset, error) {
	bytes, err := ConfigHoodJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/hood.json", size: 313, mode: os.FileMode(420), modTime: time.Unix(1443157418, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesCTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\xe2\x52\x00\x02\xe7\xa2\xd4\xc4\x92\xd4\x14\x85\xa4\x4a\x85\xea\x6a\xbd\xd0\xe2\xd4\xa2\xda\x5a\x85\xfc\x3c\x10\xc7\x05\x28\x51\x5b\xab\xc7\xa5\xa5\x0f\x08\x00\x00\xff\xff\x6d\xb8\x6e\xbe\x2c\x00\x00\x00")

func TemplatesCTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesCTpl,
		"../templates/.c.tpl",
	)
}

func TemplatesCTpl() (*asset, error) {
	bytes, err := TemplatesCTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.c.tpl", size: 44, mode: os.FileMode(420), modTime: time.Unix(1443157497, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesCssTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\xd2\xe2\x52\xd0\x52\x70\x2e\x4a\x4d\x2c\x49\x4d\x51\x48\xaa\x54\xa8\xae\xd6\x0b\x2d\x4e\x2d\xaa\xad\x55\xc8\xcf\x03\x71\x5c\x80\x12\xb5\xb5\x7a\x40\x55\xfa\x5c\x80\x00\x00\x00\xff\xff\x71\x77\x91\xf8\x2e\x00\x00\x00")

func TemplatesCssTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesCssTpl,
		"../templates/.css.tpl",
	)
}

func TemplatesCssTpl() (*asset, error) {
	bytes, err := TemplatesCssTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.css.tpl", size: 46, mode: os.FileMode(420), modTime: time.Unix(1442901457, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesGitignoreTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\xf3\x4c\xcf\xcb\x2f\x4a\x4d\xf1\xc9\x2c\x2e\xa9\xad\xad\xae\xd6\xab\xad\xe5\xaa\xae\x4e\xcd\x4b\xa9\xad\x05\x04\x00\x00\xff\xff\x72\x27\x24\x7a\x23\x00\x00\x00")

func TemplatesGitignoreTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesGitignoreTpl,
		"../templates/.gitignore.tpl",
	)
}

func TemplatesGitignoreTpl() (*asset, error) {
	bytes, err := TemplatesGitignoreTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.gitignore.tpl", size: 35, mode: os.FileMode(420), modTime: time.Unix(1442899475, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesGoTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\x0b\x80\x30\x6b\x6b\x01\x01\x00\x00\xff\xff\x83\x69\xe6\xd5\x14\x00\x00\x00")

func TemplatesGoTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesGoTpl,
		"../templates/.go.tpl",
	)
}

func TemplatesGoTpl() (*asset, error) {
	bytes, err := TemplatesGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.go.tpl", size: 20, mode: os.FileMode(420), modTime: time.Unix(1442901368, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesHtmlTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xb2\x81\x51\xa9\x89\x29\x76\x5c\x0a\x40\x60\x93\x9b\x5a\x92\xa8\x90\x9c\x91\x58\x54\x9c\x5a\x62\xab\x14\x1a\xe2\xa6\x6b\xa1\x04\x95\x2a\xc9\x2c\xc9\x49\xb5\xab\xae\xd6\x0b\x01\x31\x6a\x6b\x6d\xf4\x21\x22\x5c\x36\xfa\x10\x03\x6c\x92\xf2\x53\x2a\xed\xb8\x80\x7c\x08\x03\x28\x0e\x32\x1f\x10\x00\x00\xff\xff\xb1\x30\x76\xf2\x76\x00\x00\x00")

func TemplatesHtmlTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesHtmlTpl,
		"../templates/.html.tpl",
	)
}

func TemplatesHtmlTpl() (*asset, error) {
	bytes, err := TemplatesHtmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.html.tpl", size: 118, mode: os.FileMode(420), modTime: time.Unix(1442901496, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesJavaTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\xd2\xe2\x52\xd0\x52\x70\x2e\x4a\x4d\x2c\x49\x4d\x51\x48\xaa\x54\xa8\xae\xd6\x0b\x2d\x4e\x2d\xaa\xad\x55\xc8\xcf\x03\x71\x5c\x80\x12\xb5\xb5\x7a\x40\x55\xfa\x80\x00\x00\x00\xff\xff\xf1\x68\xb3\xfc\x2d\x00\x00\x00")

func TemplatesJavaTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesJavaTpl,
		"../templates/.java.tpl",
	)
}

func TemplatesJavaTpl() (*asset, error) {
	bytes, err := TemplatesJavaTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.java.tpl", size: 45, mode: os.FileMode(420), modTime: time.Unix(1442901520, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesJsTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\xd2\xe2\x52\xd0\x52\x70\x2e\x4a\x4d\x2c\x49\x4d\x51\x48\xaa\x54\xa8\xae\xd6\x0b\x2d\x4e\x2d\xaa\xad\x55\xc8\xcf\x03\x71\x5c\x80\x12\xb5\xb5\x7a\x40\x55\xfa\x80\x00\x00\x00\xff\xff\xf1\x68\xb3\xfc\x2d\x00\x00\x00")

func TemplatesJsTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesJsTpl,
		"../templates/.js.tpl",
	)
}

func TemplatesJsTpl() (*asset, error) {
	bytes, err := TemplatesJsTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.js.tpl", size: 45, mode: os.FileMode(420), modTime: time.Unix(1442901555, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesPhpTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xb1\x2f\xc8\x28\xe0\xd2\xd7\xd2\xe2\x52\xd0\x52\x70\x2c\x2d\xc9\xc8\x2f\x52\xa8\xae\xd6\x0b\x2d\x4e\x2d\xaa\xad\x05\x89\xb9\x24\x96\xa4\x82\x44\x40\x34\x58\x44\x1f\x10\x00\x00\xff\xff\xd4\x12\x33\x03\x33\x00\x00\x00")

func TemplatesPhpTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPhpTpl,
		"../templates/.php.tpl",
	)
}

func TemplatesPhpTpl() (*asset, error) {
	bytes, err := TemplatesPhpTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.php.tpl", size: 51, mode: os.FileMode(420), modTime: time.Unix(1442901140, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _TemplatesPyTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x52\x52\xe2\x72\x2e\x4a\x4d\x2c\x49\x4d\x51\x48\xaa\x54\xa8\xae\xd6\x0b\x2d\x4e\x2d\xaa\xad\x55\xc8\xcf\x03\x71\x5c\x80\x12\xb5\xb5\x7a\x5c\x40\x65\x80\x00\x00\x00\xff\xff\x98\x61\xe9\xaa\x2a\x00\x00\x00")

func TemplatesPyTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPyTpl,
		"../templates/.py.tpl",
	)
}

func TemplatesPyTpl() (*asset, error) {
	bytes, err := TemplatesPyTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/.py.tpl", size: 42, mode: os.FileMode(420), modTime: time.Unix(1442901684, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"../config/hood.json": ConfigHoodJson,
	"../templates/.c.tpl": TemplatesCTpl,
	"../templates/.css.tpl": TemplatesCssTpl,
	"../templates/.gitignore.tpl": TemplatesGitignoreTpl,
	"../templates/.go.tpl": TemplatesGoTpl,
	"../templates/.html.tpl": TemplatesHtmlTpl,
	"../templates/.java.tpl": TemplatesJavaTpl,
	"../templates/.js.tpl": TemplatesJsTpl,
	"../templates/.php.tpl": TemplatesPhpTpl,
	"../templates/.py.tpl": TemplatesPyTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"config": &bintree{nil, map[string]*bintree{
			"hood.json": &bintree{ConfigHoodJson, map[string]*bintree{
			}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			".c.tpl": &bintree{TemplatesCTpl, map[string]*bintree{
			}},
			".css.tpl": &bintree{TemplatesCssTpl, map[string]*bintree{
			}},
			".gitignore.tpl": &bintree{TemplatesGitignoreTpl, map[string]*bintree{
			}},
			".go.tpl": &bintree{TemplatesGoTpl, map[string]*bintree{
			}},
			".html.tpl": &bintree{TemplatesHtmlTpl, map[string]*bintree{
			}},
			".java.tpl": &bintree{TemplatesJavaTpl, map[string]*bintree{
			}},
			".js.tpl": &bintree{TemplatesJsTpl, map[string]*bintree{
			}},
			".php.tpl": &bintree{TemplatesPhpTpl, map[string]*bintree{
			}},
			".py.tpl": &bintree{TemplatesPyTpl, map[string]*bintree{
			}},
		}},
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

