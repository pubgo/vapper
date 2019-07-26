package cmds

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/internal/config"
	"github.com/pubgo/vapper/pkg/transpiler"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// buildCmd represents the build command

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "transpiler",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers a¬pplications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			cwd, err := os.Getwd()
			errors.Panic(err)

			cfg := config.Default()
			cfg.Init()

			transpiler.ProcessAll(filepath.Join(cwd, "components"), "components")
			transpiler.ProcessAll(filepath.Join(cwd, "views"), "views")
		},
	})
}
