package cmds

import (
	"github.com/pubgo/errors"
	"github.com/spf13/cobra"
)

func initConfigCmd(cmd *cobra.Command) *cobra.Command {
	return cmd
}

func init() {
	rootCmd.AddCommand(initConfigCmd(&cobra.Command{
		Use:   "cfg",
		Short: "cmd config",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

			//_curDir, err := os.Getwd()
			//errors.Panic(err)

		},
	}))
}
