package cmds

import (
	"fmt"
	"github.com/pubgo/vapper/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version", version.Version)
			fmt.Println("commitV", version.CommitV)
			fmt.Println("buildV", version.BuildV)
		},
	})
}
