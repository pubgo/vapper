package jsbuild

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/compiler/prelude"
	"github.com/pubgo/errors"
	"path/filepath"
	"strings"
	"text/template"
)

const jsGoPrelude = `$load.prelude=function(){};`

var mainTemplateMinified = template.Must(template.New("main").Parse(
	`"use strict";var $mainPkg,$load={};!function(){for(var n=0,t=0,e={{ .Json }},o=(document.getElementById("log"),function(){n++,window.jsgoProgress&&window.jsgoProgress(n,t),n==t&&function(){for(var n=0;n<e.length;n++)$load[e[n].path]();$mainPkg=$packages["{{ .Path }}"],$synthesizeMethods(),$packages.runtime.$init(),$go($mainPkg.$init,[]),$flushConsole()}()}),a=function(n){t++;var e=document.createElement("script");e.src=n,e.onload=o,e.onreadystatechange=o,document.head.appendChild(e)},s=0;s<e.length;s++)a("{{ .PkgProtocol }}://{{ .PkgHost }}/"+e[s].path+"."+e[s].hash+".js")}();`,
))

type MainVars struct {
	Path        string
	Json        string
	PkgHost     string
	PkgProtocol string
}

type Pkg struct {
	Path    string `json:"path,omitempty"`
	Hash    string `json:"hash,omitempty"`
	Content []byte `json:"-"`
}

func Default() *_Build {
	return &_Build{Options: &gbuild.Options{CreateMapFile: true}}
}

type _Build struct {
	Options *gbuild.Options

	RootPath string
	Tags     string
	Addr     string

	pkgIndex []*Pkg
	pkgData  map[string]*Pkg
	pkgMain  Pkg

	OnlyHash bool
}

func (t *_Build) Hash(d []byte) string {
	h := sha1.New()
	h.Write(d)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *_Build) Build() {
	s := gbuild.NewSession(t.Options)
	//s.Watcher.Add(".")

	//if s.Watcher != nil {
	pkg, err := gbuild.Import(t.RootPath, 0, s.InstallSuffix(), t.Options.BuildTags)
	errors.Panic(err)
	errors.T(!pkg.IsCommand(), "not main package")
	//errors.Panic(s.Watcher.Add(pkg.Dir))
	//}

	//for {
	archive, err := s.BuildPackage(pkg)
	errors.Panic(err)

	deps, err := compiler.ImportDependencies(archive, s.BuildImportPath)
	errors.Panic(err)

	t.pkgIndex = []*Pkg{}
	t.pkgData = make(map[string]*Pkg)

	// gen prelude
	_preludeCode := []byte(prelude.Minified + jsGoPrelude)
	_pre := &Pkg{
		Content: _preludeCode,
		Hash:    t.Hash(_preludeCode),
		Path:    "prelude",
	}
	t.pkgIndex = append(t.pkgIndex, _pre)
	t.pkgData[_pre.Path] = _pre

	// gen pkgs
	_vendor := filepath.Join(t.RootPath, "vendor/")
	for _, dep := range deps {
		if strings.HasPrefix(dep.ImportPath, _vendor) {
			dep.ImportPath = strings.ReplaceAll(dep.ImportPath, _vendor, "")
		}


		content := t.GetPackageCode(dep, t.Options.Minify)
		_pkg := &Pkg{
			Content: content,
			Path:    dep.ImportPath,
			Hash:    t.Hash(content),
		}
		t.pkgIndex = append(t.pkgIndex, _pkg)
		t.pkgData[_pkg.Path] = _pkg

		fmt.Println(filepath.Join(t.RootPath, "vendor"), dep.Name, dep.ImportPath, t.Hash(content), string(content)[:100])
	}

	// gen main
	dt, err := json.Marshal(t.pkgIndex)
	errors.Panic(err)

	fmt.Println(pkg.ImportPath, pkg.Name)
	buf := &bytes.Buffer{}
	errors.Panic(mainTemplateMinified.Execute(buf, &MainVars{
		Path:        pkg.ImportPath,
		Json:        string(dt),
		PkgHost:     "localhost:8080/pkg",
		PkgProtocol: "http",
	}))
	t.pkgMain = Pkg{
		Content: buf.Bytes(),
		Path:    "main",
		Hash:    t.Hash(buf.Bytes()),
	}

	fmt.Println(string(t.pkgMain.Content))
	//s.WaitForChange()
	//}
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
