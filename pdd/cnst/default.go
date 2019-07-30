package cnst

import (
	"go/build"
	"path/filepath"
)

var Default = struct {
	JsPkg string
}{
	JsPkg: filepath.Join(build.Default.GOPATH, "js_pkg"),
}

