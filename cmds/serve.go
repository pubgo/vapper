package cmds

import (
	"bytes"
	"fmt"
	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"github.com/pubgo/errors"
	"github.com/spf13/cobra"
	"go/build"
	"golang.org/x/crypto/ssh/terminal"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	options = &gbuild.Options{CreateMapFile: true}
	pkgObj  string
	tags    string
	addr    string
)

func initServeCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolVarP(&options.Verbose, "verbose", "v", false, "print the names of packages as they are compiled")
	cmd.Flags().BoolVarP(&options.Quiet, "quiet", "q", false, "suppress non-fatal warnings")
	cmd.Flags().BoolVarP(&options.Minify, "minify", "m", false, "minify generated code")
	cmd.Flags().BoolVar(&options.Color, "color", terminal.IsTerminal(int(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb", "colored output")
	cmd.Flags().StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")
	cmd.Flags().BoolVar(&options.MapToLocalDisk, "localmap", false, "use local paths for sourcemap")
	cmd.Flags().StringVarP(&addr, "http", "", ":8080", "HTTP bind address to serve")
	return cmd
}

func init() {
	rootCmd.AddCommand(initServeCmd(&cobra.Command{
		Use:   "serve [root]",
		Short: "compile on-the-fly and serve",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

			options.BuildTags = strings.Fields(tags)
			dirs := append(filepath.SplitList(build.Default.GOPATH), build.Default.GOROOT)

			errors.T(len(args) != 1, "root num error")

			root := args[0]

			sourceFiles := http.FileServer(serveCommandFileSystem{
				serveRoot:  root,
				options:    options,
				dirs:       dirs,
				sourceMaps: make(map[string][]byte),
			})

			ln, err := net.Listen("tcp", addr)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if tcpAddr := ln.Addr().(*net.TCPAddr); tcpAddr.IP.Equal(net.IPv4zero) || tcpAddr.IP.Equal(net.IPv6zero) { // Any available addresses.
				fmt.Printf("serving at http://localhost:%d and on port %d of any available addresses\n", tcpAddr.Port, tcpAddr.Port)
			} else { // Specific address.
				fmt.Printf("serving at http://%s\n", tcpAddr)
			}
			fmt.Fprintln(os.Stderr, http.Serve(ln, sourceFiles))
		},
	}))
}

type serveCommandFileSystem struct {
	serveRoot  string
	options    *gbuild.Options
	dirs       []string
	sourceMaps map[string][]byte
}

func (fs serveCommandFileSystem) Open(requestName string) (http.File, error) {
	name := path.Join(fs.serveRoot, requestName[1:]) // requestName[0] == '/'

	dir, file := path.Split(name)
	base := path.Base(dir) // base is parent folder name, which becomes the output file name.

	isPkg := file == base+".js"
	isMap := file == base+".js.map"
	isIndex := file == "index.html"

	if isPkg || isMap || isIndex {
		// If we're going to be serving our special files, make sure there's a Go command in this folder.
		s := gbuild.NewSession(fs.options)
		pkg, err := gbuild.Import(path.Dir(name), 0, s.InstallSuffix(), fs.options.BuildTags)
		if err != nil || pkg.Name != "main" {
			isPkg = false
			isMap = false
			isIndex = false
		}

		switch {
		case isPkg:
			buf := new(bytes.Buffer)
			browserErrors := new(bytes.Buffer)
			err := func() error {
				archive, err := s.BuildPackage(pkg)
				if err != nil {
					return err
				}

				sourceMapFilter := &compiler.SourceMapFilter{Writer: buf}
				m := &sourcemap.Map{File: base + ".js"}
				sourceMapFilter.MappingCallback = gbuild.NewMappingCallback(m, fs.options.GOROOT, fs.options.GOPATH, fs.options.MapToLocalDisk)

				deps, err := compiler.ImportDependencies(archive, s.BuildImportPath)
				if err != nil {
					return err
				}
				if err := compiler.WriteProgramCode(deps, sourceMapFilter); err != nil {
					return err
				}

				mapBuf := new(bytes.Buffer)
				m.WriteTo(mapBuf)
				buf.WriteString("//# sourceMappingURL=" + base + ".js.map\n")
				fs.sourceMaps[name+".map"] = mapBuf.Bytes()

				return nil
			}()
			handleError(err, fs.options, browserErrors)
			if err != nil {
				buf = browserErrors
			}
			return newFakeFile(base+".js", buf.Bytes()), nil

		case isMap:
			if content, ok := fs.sourceMaps[name]; ok {
				return newFakeFile(base+".js.map", content), nil
			}
		}
	}

	for _, d := range fs.dirs {
		dir := http.Dir(filepath.Join(d, "src"))

		f, err := dir.Open(name)
		if err == nil {
			return f, nil
		}

		// source maps are served outside of serveRoot
		f, err = dir.Open(requestName)
		if err == nil {
			return f, nil
		}
	}

	if isIndex {
		// If there was no index.html file in any dirs, supply our own.
		return newFakeFile("index.html", []byte(`<html><head><meta charset="utf-8"><script src="`+base+`.js"></script></head><body></body></html>`)), nil
	}

	return nil, os.ErrNotExist
}
