package cmds

import (
	gbuild "github.com/gopherjs/gopherjs/build"
	"github.com/kisielk/gotool"
	"github.com/pubgo/errors"
	"github.com/spf13/cobra"
	"go/build"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
	"strings"
)

func initBuildCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&pkgObj, "output", "o", "", "output file")
	cmd.Flags().BoolVarP(&options.Verbose, "verbose", "v", false, "print the names of packages as they are compiled")
	cmd.Flags().BoolVarP(&options.Quiet, "quiet", "q", false, "suppress non-fatal warnings")
	cmd.Flags().BoolVarP(&options.Minify, "minify", "m", false, "minify generated code")
	cmd.Flags().BoolVar(&options.Color, "color", terminal.IsTerminal(int(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb", "colored output")
	cmd.Flags().StringVar(&tags, "tags", "", "a list of build tags to consider satisfied during the build")
	cmd.Flags().BoolVar(&options.MapToLocalDisk, "localmap", false, "use local paths for sourcemap")
	cmd.Flags().BoolVarP(&options.Watch, "watch", "w", true, "watch for changes to the source files")
	return cmd
}

func init() {
	rootCmd.AddCommand(initBuildCmd(&cobra.Command{
		Use:   "build [packages]",
		Short: "compile packages and dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()
			options.BuildTags = strings.Fields(tags)

			_func := func(s *gbuild.Session) {
				// Handle "gopherjs build [files]" ad-hoc package mode.

				if len(args) > 0 && (strings.HasSuffix(args[0], ".go") || strings.HasSuffix(args[0], ".inc.js")) {
					for _, arg := range args {
						errors.T(!strings.HasSuffix(arg, ".go") && !strings.HasSuffix(arg, ".inc.js"), "named files must be .go or .inc.js files")
					}

					if pkgObj == "" {
						basename := filepath.Base(args[0])
						pkgObj = basename[:len(basename)-3] + ".js"
					}
					names := make([]string, len(args))
					for i, name := range args {
						name = filepath.ToSlash(name)
						names[i] = name
						if s.Watcher != nil {
							errors.Panic(s.Watcher.Add(name))
						}
					}
					errors.Panic(s.BuildFiles(args, pkgObj, currentDirectory))
				}

				// Expand import path patterns.
				patternContext := gbuild.NewBuildContext("", options.BuildTags)
				pkgs := (&gotool.Context{BuildContext: *patternContext}).ImportPaths(args)

				for _, pkgPath := range pkgs {
					if s.Watcher != nil {
						pkg, err := gbuild.NewBuildContext(s.InstallSuffix(), options.BuildTags).Import(pkgPath, "", build.FindOnly)
						errors.Panic(err)
						errors.Panic(s.Watcher.Add(pkg.Dir))
					}

					pkg, err := gbuild.Import(pkgPath, 0, s.InstallSuffix(), options.BuildTags)
					errors.Panic(err)

					archive, err := s.BuildPackage(pkg)
					errors.Panic(err)

					if len(pkgs) == 1 { // Only consider writing output if single package specified.
						if pkgObj == "" {
							pkgObj = filepath.Base(pkg.Dir) + ".js"
						}

						if pkg.IsCommand() && !pkg.UpToDate {
							errors.Panic(s.WriteCommandPackage(archive, pkgObj))
						}
					}
				}
			}

			for {
				s := gbuild.NewSession(options)
				errors.ErrHandle(errors.Try(_func)(s), func(err *errors.Err) {
					err.P()
					//exitCode := handleError(err.Err(), options, nil)
					//if s.Watcher == nil {
					//	os.Exit(exitCode)
					//}
				})
				s.WaitForChange()
			}
		},
	}))
}
