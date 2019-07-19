package jsbuild

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/envy"
	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/compiler/prelude"
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/internal/config"
	"path/filepath"
	"strings"
	"text/template"
)

const jsGoPrelude = `$load.prelude=function(){};`

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
	return &_Build{Options: &gbuild.Options{CreateMapFile: true, Watch: true}}
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
}

func (t *_Build) Hash(d []byte) string {
	h := sha1.New()
	h.Write(d)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *_Build) Build() {
	defer errors.Assert()

	cfg := config.Default()

	// gen prelude
	_preludeCode := []byte(prelude.Minified + jsGoPrelude)
	_pre := &_Pkg{
		Content: _preludeCode,
		Hash:    t.Hash(_preludeCode),
		Path:    "prelude",
	}

	var sess *gbuild.Session
	errors.Wrap(errors.Try(func() {
		sess = gbuild.NewSession(t.Options)
	}), "session error")

	errors.T(sess.Watcher == nil, "file watcher error")
	errors.Wrap(sess.Watcher.Add(t.RootPath), "watch error")

	mainPkg, err := gbuild.Import(envy.CurrentPackage(), 0, sess.InstallSuffix(), t.Options.BuildTags)
	errors.Wrap(err, "import error")

	fmt.Print("11111\n\nn\n\n")
	errors.T(!mainPkg.IsCommand(), "not main package")
	errors.Wrap(sess.Watcher.Add(mainPkg.Dir), "watch main pkg error")

	fmt.Print("222222\n\nn\n\n")
	for {
		archive, err := sess.BuildPackage(mainPkg)
		errors.Wrap(err, "BuildPackage error")

		deps, err := compiler.ImportDependencies(archive, sess.BuildImportPath)
		errors.Wrap(err, "ImportDependencies error")

		t.pkgIndex = []*_Pkg{}
		t.pkgData = make(map[string]*_Pkg)

		t.pkgIndex = append(t.pkgIndex, _pre)
		t.pkgData[_pre.Path] = _pre

		// gen pkgs
		_vendor := filepath.Join(t.RootPath, "vendor/")
		for _, dep := range deps {
			if strings.HasPrefix(dep.ImportPath, _vendor) {
				dep.ImportPath = strings.ReplaceAll(dep.ImportPath, _vendor, "")
			}

			content := t.GetPackageCode(dep, t.Options.Minify)
			_pkg := &_Pkg{
				Content: content,
				Path:    dep.ImportPath,
				Hash:    t.Hash(content),
			}
			t.pkgIndex = append(t.pkgIndex, _pkg)
			t.pkgData[_pkg.Path] = _pkg

			fmt.Println(dep.Name, dep.ImportPath, t.Hash(content), string(content)[:100])
		}

		// gen main
		dt, err := json.Marshal(t.pkgIndex)
		errors.Wrap(err,"pkgIndex json error")

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

		errors.ErrHandle(errors.Try(sess.WaitForChange), func(err *errors.Err) {
			fmt.Println(err.P())
		})
	}
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