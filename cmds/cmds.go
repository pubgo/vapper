package cmds

import (
	"fmt"
	"github.com/pubgo/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func initGenCmdCmd(cmd *cobra.Command) *cobra.Command {
	return cmd
}

func init() {
	var _cmd = "test"
	rootCmd.AddCommand(initGenCmdCmd(&cobra.Command{
		Use:     "cmd",
		Short:   "db2rest gen cmd",
		Example: "db2rest cmd hello",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

			if len(args) > 0 {
				_cmd = args[0]
			}

			_file := fmt.Sprintf("cmds/%s.go", _cmd)
			log.Info().Msgf("gen cmd file: %s", _file)

			_cmd1 := strings.ToUpper(string(_cmd[0])) + _cmd[1:]
			errors.Wrap(ioutil.WriteFile(_file, []byte(fmt.Sprintf(genCmdTpl, _cmd1, _cmd, _cmd)), 0644), "gen cmd error")
		},
	}))
}

var genCmdTpl = `package cmds

import (
	"github.com/pubgo/errors"
	"github.com/spf13/cobra"
)

func init%sCmd(cmd *cobra.Command) *cobra.Command {
	return cmd
}

func init() {
	rootCmd.AddCommand(init%sCmd(&cobra.Command{
		Use:   "%s",
		Short: "db2rest %s",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

		},
	}))
}`
