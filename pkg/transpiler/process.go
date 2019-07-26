package transpiler

import (
	"github.com/gobuffalo/envy"
	"github.com/pubgo/errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile("[^a-zA-Z0-9]+")

func _ComponentName(path string) string {
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(reg.ReplaceAllString(base, ""))
}

func _IsHTML(info os.FileInfo) bool {
	_ext := filepath.Ext(info.Name())
	return _ext == ".html" || _ext == ".ghtml" || _ext == ".gohtml"
}

func _GeneratedGoFileName(base, name string) string {
	return filepath.Join(base, strings.ToLower(name)+"_generated.go")
}

// ProcessAll processes components starting at base
func ProcessAll(base string, packageName string) {
	errors.Wrap(filepath.Walk(base, func(path string, info os.FileInfo, err error) (e error) {
		defer errors.Resp(func(_err *errors.Err) {
			e = _err
		})
		errors.Wrap(err, "file error")

		if !info.IsDir() && _IsHTML(info) {
			f, err := os.Open(path)
			errors.Wrap(err, "file open error")

			comp := _ComponentName(path)
			gfn := filepath.Join(base, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			var makeStruct bool
			if os.IsNotExist(err) {
				makeStruct = true
			}

			makeStruct = true

			gf, err := os.Create(_GeneratedGoFileName(base, comp))
			errors.Wrap(err, "file create error")
			defer errors.Panic(gf.Close)

			_, err = io.WriteString(gf, NewTranspiler(f, makeStruct, envy.CurrentPackage(), comp, packageName).Code())
			errors.Wrap(err, "file write error")
		}
		return
	}), "error walking the path %s \n", base)
}
