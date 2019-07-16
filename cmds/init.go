// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmds

import (
	"bytes"
	"fmt"
	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/pubgo/vapper/templates"
	"go/build"
	"go/scanner"
	"go/types"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"text/template"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/pubgo/errors"
	"github.com/spf13/cobra"
)

var appName = "example"
var directories = []string{
	"app",
	"assets",
	"client",
	"components",
	"models",
	"routes",
	"server",
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Example: "./factor init example",
	Short:   "initialize a new factor application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			appName = args[0]
		}

		cwd, err := os.Getwd()
		errors.Panic(err)

		// make directories under appName
		for _, dir := range directories {
			errors.Wrap(os.MkdirAll(filepath.Join(cwd, appName, dir), 0755), "MkdirAll Error: %s", dir)
		}

		appPkg := filepath.Join(envy.CurrentPackage(), appName)
		errors.T(appPkg == "", "couldn't get the current package for the app")
		fmt.Println("app package", appPkg)

		// put the new files there
		writeTemplate(filepath.Join(cwd, appName, "app", "index.html"), templates.IndexHTML)
		//writeTemplate(filepath.Join(cwd, appName, "app", "wasm_exec.js"), templates.WasmJS)
		writeTemplate(filepath.Join(cwd, appName, "assets", "global.css"), templates.GlobalCSS)
		writeTemplate(filepath.Join(cwd, appName, "client", "main.go"), templates.ClientGoMain(appPkg))
		writeTemplate(filepath.Join(cwd, appName, "components", "Nav.html"), templates.NavComponentHTML)
		writeTemplate(filepath.Join(cwd, appName, "components", "nav.go"), templates.NavComponentGo)
		writeTemplate(filepath.Join(cwd, appName, "routes", "Index.html"), templates.RoutesHTML)
		writeTemplate(filepath.Join(cwd, appName, "routes", "index.go"), templates.RoutesGo)
		writeTemplate(filepath.Join(cwd, appName, "models", "todo.go"), templates.TodoClient)
		writeTemplate(filepath.Join(cwd, appName, "server", "main.go"), templates.ServerGoMain(appPkg))
		writeTemplate(filepath.Join(cwd, appName, "Makefile"), templates.Makefile)
	},
}

func writeTemplate(filePath string, templateName string) {
	tf, err := os.Create(filePath)
	errors.Wrap(err, "os Create Error: %s", filePath)
	defer errors.Panic(tf.Close)

	errors.Panic(template.Must(template.New("component").Parse(templateName)).Execute(tf, nil))
}

func init() {
	rootCmd.AddCommand(initCmd)
}

type fakeFile struct {
	name string
	size int
	io.ReadSeeker
}

func newFakeFile(name string, content []byte) *fakeFile {
	return &fakeFile{name: name, size: len(content), ReadSeeker: bytes.NewReader(content)}
}

func (f *fakeFile) Close() error {
	return nil
}

func (f *fakeFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, os.ErrInvalid
}

func (f *fakeFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *fakeFile) Name() string {
	return f.name
}

func (f *fakeFile) Size() int64 {
	return int64(f.size)
}

func (f *fakeFile) Mode() os.FileMode {
	return 0
}

func (f *fakeFile) ModTime() time.Time {
	return time.Time{}
}

func (f *fakeFile) IsDir() bool {
	return false
}

func (f *fakeFile) Sys() interface{} {
	return nil
}

// handleError handles err and returns an appropriate exit code.
// If browserErrors is non-nil, errors are written for presentation in browser.
func handleError(err error, options *gbuild.Options, browserErrors *bytes.Buffer) int {
	switch err := err.(type) {
	case nil:
		return 0
	case compiler.ErrorList:
		for _, entry := range err {
			printError(entry, options, browserErrors)
		}
		return 1
	case *exec.ExitError:
		return err.Sys().(syscall.WaitStatus).ExitStatus()
	default:
		printError(err, options, browserErrors)
		return 1
	}
}

// printError prints err to Stderr with options. If browserErrors is non-nil, errors are also written for presentation in browser.
func printError(err error, options *gbuild.Options, browserErrors *bytes.Buffer) {
	e := sprintError(err)
	options.PrintError("%s\n", e)
	if browserErrors != nil {
		fmt.Fprintln(browserErrors, `console.error("`+template.JSEscapeString(e)+`");`)
	}
}

// sprintError returns an annotated error string without trailing newline.
func sprintError(err error) string {
	makeRel := func(name string) string {
		if relname, err := filepath.Rel(currentDirectory, name); err == nil {
			return relname
		}
		return name
	}

	switch e := err.(type) {
	case *scanner.Error:
		return fmt.Sprintf("%s:%d:%d: %s", makeRel(e.Pos.Filename), e.Pos.Line, e.Pos.Column, e.Msg)
	case types.Error:
		pos := e.Fset.Position(e.Pos)
		return fmt.Sprintf("%s:%d:%d: %s", makeRel(pos.Filename), pos.Line, pos.Column, e.Msg)
	default:
		return fmt.Sprintf("%s", e)
	}
}

var currentDirectory string

func init() {
	var err error
	currentDirectory, err = os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	currentDirectory, err = filepath.EvalSymlinks(currentDirectory)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	gopaths := filepath.SplitList(build.Default.GOPATH)
	if len(gopaths) == 0 {
		fmt.Fprintf(os.Stderr, "$GOPATH not set. For more details see: go help gopath\n")
		os.Exit(1)
	}
}
