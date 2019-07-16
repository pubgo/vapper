package cmds

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/internal/config"
	"github.com/pubgo/vapper/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const homeFlag = "home"
const debugFlag = "debug"

var rootCmd = &cobra.Command{
	Use:     "db2rest",
	Short:   "db2rest app",
	Version: version.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		defer errors.Assert()

		for _, name := range []string{"version"} {
			if name == cmd.Name() {
				return
			}
		}

		errors.Wrap(viper.BindPFlags(cmd.Flags()), "Flags Error")

		viper.SetConfigType("toml")
		viper.SetConfigName("config")

		homeDir := viper.GetString(homeFlag)
		viper.AddConfigPath(homeDir)                          // search root directory
		viper.AddConfigPath(filepath.Join(homeDir, "config")) // search root directory /kdata

		home, err := homedir.Dir()
		errors.Wrap(err, "home dir error")
		viper.AddConfigPath(home)

		viper.AddConfigPath("/app/config")
		viper.AddConfigPath("config")
		viper.AddConfigPath("pdd/config")

		// load data
		errors.Wrap(viper.ReadInConfig(), "check kata error")

		config.Default().Parse()
	},
}

func Execute(envPrefix, defaultHome string) {
	defer errors.Assert()

	cobra.OnInitialize(func() { initEnv(envPrefix) })
	rootCmd.PersistentFlags().StringP(homeFlag, "", defaultHome, "directory for data")
	rootCmd.PersistentFlags().BoolP(debugFlag, "d", true, "debug mode")
	rootCmd.PersistentFlags().StringP("ll", "l", "debug", "log level")
	rootCmd.PersistentFlags().StringP("env", "e", "dev", "project environment")
	errors.Panic(rootCmd.Execute())
}

// initEnv sets to use ENV variables if set.
func initEnv(prefix string) {
	copyEnvVars(prefix)

	// env variables with TM prefix (eg. TM_ROOT)
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
}

// This copies all variables like TMROOT to TM_ROOT,
// so we can support both formats for the user
func copyEnvVars(prefix string) {
	prefix = strings.ToUpper(prefix)
	ps := prefix + "_"
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		if len(kv) == 2 {
			k, v := kv[0], kv[1]
			if strings.HasPrefix(k, prefix) && !strings.HasPrefix(k, ps) {
				k2 := strings.Replace(k, prefix, ps, 1)
				errors.Wrap(os.Setenv(k2, v), "env set error")
			}
		}
	}
}
