package transpiler

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var fileNameRegexp = regexp.MustCompile("[^a-zA-Z0-9]+")
var reg = regexp.MustCompile("[^a-zA-Z0-9]+")

func SanitizedName(path string) string {
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(fileNameRegexp.ReplaceAllString(base, ""))
}

func IsHTML(info os.FileInfo) bool {
	_ext := filepath.Ext(info.Name())
	return _ext == ".html" || _ext == ".ghtml" || _ext == ".gohtml"
}

func GeneratedGoFileName(base, name string) string {
	return filepath.Join(base, strings.ToLower(name)+"_generated.go")
}
func ComponentName(path string) string {
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(reg.ReplaceAllString(base, ""))
}
