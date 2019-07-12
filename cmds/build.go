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
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/pkg/transpiler"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers a¬pplications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Building Application")
		cwd, err := os.Getwd()
		errors.Panic(err)

		transpiler.ProcessAll(filepath.Join(cwd, "components"))
		transpiler.ProcessAll(filepath.Join(cwd, "routes"))
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
