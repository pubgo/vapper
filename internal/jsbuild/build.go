package jsbuild

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/gobuffalo/envy"
	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/compiler/prelude"
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/internal/config"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const jsGoPrelude = `$load.prelude=function(){};`

var ValidExtensions = []string{".go", ".jsgo.html", ".inc.js", ".md"}
var mainTemplateMinified = template.Must(template.New("main").Parse(
	`"use strict";var $mainPkg,$load={};!function(){for(var n=0,t=0,e={{ .Json }},o=(document.getElementById("log"),function(){n++,window.jsgoProgress&&window.jsgoProgress(n,t),n==t&&function(){for(var n=0;n<e.length;n++)$load[e[n].path]();$mainPkg=$packages["{{ .Path }}"],$synthesizeMethods(),$packages.runtime.$init(),$go($mainPkg.$init,[]),$flushConsole()}()}),a=function(n){t++;var e=document.createElement("script");e.src=n,e.onload=o,e.onreadystatechange=o,document.head.appendChild(e)},s=0;s<e.length;s++)a("{{ .PkgUrl }}/"+e[s].path+"."+e[s].hash+".js")}();`,
))

type _MainVars struct {
	Path   string
	Json   string
	PkgUrl string
}

type _Pkg struct {
	Path    string `json:"path,omitempty"`
	Hash    string `json:"hash,omitempty"`
	Content []byte `json:"-"`
}

func Default() *_Build {
	return &_Build{Options: &gbuild.Options{CreateMapFile: true, Watch: true}, deps: make(map[string]*_Pkg)}
}

type _Build struct {
	Options *gbuild.Options

	RootPath string
	Tags     string
	Addr     string

	pkgIndex []*_Pkg
	pkgData  map[string]*_Pkg
	pkgMain  _Pkg

	OnlyHash bool
	deps     map[string]*_Pkg
}

func (t *_Build) MainUrl() _Pkg {
	return t.pkgMain
}

func (t *_Build) GetByHash(h string) *_Pkg {
	return t.pkgData[h]
}

func (t *_Build) Hash(d []byte) string {
	h := sha1.New()
	_, err := h.Write(d)
	errors.Panic(err)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *_Build) XXHash(bytes []byte) []byte {
	h := xxhash.New()
	defer h.Reset()

	_, err := h.Write(bytes)
	errors.Panic(err)

	return h.Sum(nil)
}

func (t *_Build) Build() {
	cfg := config.Default()

	// gen prelude
	_preludeCode := []byte(prelude.Minified + jsGoPrelude)
	_pre := &_Pkg{
		Content: _preludeCode,
		Hash:    t.Hash(_preludeCode),
		Path:    "prelude",
	}

	errors.Ticker(func(_ time.Time) time.Duration {
		sess := gbuild.NewSession(t.Options)
		errors.T(sess.Watcher == nil, "file watcher error")
		errors.Wrap(sess.Watcher.Add(t.RootPath), "watch error")

		curPkg := envy.CurrentPackage()
		mainPkg, err := gbuild.Import(curPkg, 0, sess.InstallSuffix(), t.Options.BuildTags)
		errors.Wrap(err, "import error, path(%s)", curPkg)

		errors.T(!mainPkg.IsCommand(), "not main package")
		errors.Wrap(sess.Watcher.Add(mainPkg.Dir), "watch main pkg error")

		archive, err := sess.BuildPackage(mainPkg)
		errors.Wrap(err, "BuildPackage error")

		deps, err := compiler.ImportDependencies(archive, sess.BuildImportPath)
		errors.Wrap(err, "ImportDependencies error")

		t.pkgIndex = t.pkgIndex[:0]
		t.pkgData = make(map[string]*_Pkg)

		t.pkgIndex = append(t.pkgIndex, _pre)
		t.pkgData[_pre.Path] = _pre

		// gen pkgs
		for _, dep := range deps {
			_dt, err := json.Marshal(dep)
			errors.Panic(err)
			_dh := t.Hash(_dt)

			var _depPkg *_Pkg
			if _pkg, ok := t.deps[_dh]; ok {
				_depPkg = _pkg
			} else {
				content := t.GetPackageCode(dep, t.Options.Minify)
				_depPkg = &_Pkg{
					Content: content,
					Path:    dep.ImportPath,
					Hash:    t.Hash(content),
				}
			}
			t.deps[_dh] = _depPkg

			t.pkgIndex = append(t.pkgIndex, _depPkg)
			t.pkgData[_depPkg.Path] = _depPkg
		}

		// gen main
		dt, err := json.Marshal(t.pkgIndex)
		errors.Wrap(err, "pkgIndex json error")

		buf := &bytes.Buffer{}
		errors.Wrap(mainTemplateMinified.Execute(buf, &_MainVars{
			Path:   mainPkg.ImportPath,
			Json:   string(dt),
			PkgUrl: cfg.Cfg.Pkg.URL,
		}), "mainTemplateMinified error")

		_, name := filepath.Split(t.RootPath)
		t.pkgMain = _Pkg{
			Content: buf.Bytes(),
			Path:    name,
			Hash:    t.Hash(buf.Bytes()),
		}

		sess.WaitForChange()
		return time.Second
	})
}

func (t *_Build) GetPackageCode(archive *compiler.Archive, minify bool) []byte {
	buf := new(bytes.Buffer)
	defer buf.Reset()

	var s string
	if minify {
		s = `$load["%s"]=function(){`
	} else {
		s = `$load["%s"] = function () {` + "\n"
	}

	buf.WriteString(fmt.Sprintf(s, archive.ImportPath))

	dceSelection := make(map[*compiler.Decl]struct{})
	for _, d := range archive.Declarations {
		dceSelection[d] = struct{}{}
	}
	errors.Panic(compiler.WritePkgCode(archive, dceSelection, minify, &compiler.SourceMapFilter{Writer: buf}))

	if minify {
		// compilexr.WritePkgCode always finishes with a "\n". In minified mode we should remove this.
		buf.Truncate(buf.Len() - 1)
	}

	buf.WriteString("};")
	return buf.Bytes()
}

// buf := &bytes.Buffer{}
// errors.Panic(compiler.WriteArchive(deployer.StripArchive(archive), buf))
// MimeBin  = "application/octet-stream"

func getStandardLibraryPackages() []string {
	cmd := exec.Command("go", "list", "./...")
	fmt.Println(cmd.Env)
	cmd.Env = []string{
		fmt.Sprintf("GOPATH=%s", build.Default.GOPATH),
		fmt.Sprintf("GOROOT=%s", build.Default.GOROOT),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
		fmt.Sprintf("GOCACHE=%s", "/Users/barry/Library/Caches/go-build"),
	}
	cmd.Dir = filepath.Join(build.Default.GOROOT, "src")
	b, err := cmd.CombinedOutput()
	fmt.Println(string(b))
	fmt.Println(err)
	errors.Panic(err)

	all := strings.Split(strings.TrimSpace(string(b)), "\n")
	excluded := map[string]bool{
		"builtin":                true,
		"internal/cpu":           true,
		"net/http/pprof":         true,
		"plugin":                 true,
		"runtime/cgo":            true,
		"os/signal/internal/pty": true,
		"cmd/pprof":              true,
		"cmd/trace":              true,
	}
	var filtered []string
	for _, p := range all {
		if excluded[p] {
			continue
		}
		filtered = append(filtered, p)
	}
	return filtered
}

