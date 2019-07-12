package transpiler

import (
	"github.com/gobuffalo/envy"
	"github.com/pubgo/errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pubgo/vapper/pkg"
)

// ProcessAll processes components starting at base
func ProcessAll(base string) {
	errors.Wrap(filepath.Walk(base, func(path string, info os.FileInfo, err error) (e error) {
		defer errors.Resp(func(_err *errors.Err) {
			e = _err
		})

		errors.Panic(err)

		if !info.IsDir() && pkg.IsHTML(info) {
			f, err := os.Open(path)
			errors.Panic(err)

			comp := pkg.ComponentName(path)
			gfn := filepath.Join(base, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			var makeStruct bool
			if os.IsNotExist(err) {
				makeStruct = true
			}

			gf, err := os.Create(pkg.GeneratedGoFileName(base, comp))
			errors.Wrap(err, "error")
			defer errors.Panic(gf.Close)

			_, err = io.WriteString(gf, NewTranspiler(f, makeStruct, envy.CurrentPackage(), comp, "components").Code())
			errors.Panic(err)
		}
		return
	}), "error walking the path %s \n", base)
}
