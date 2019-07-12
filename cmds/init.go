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
	"fmt"
	"github.com/pubgo/vapper/templates"
	"os"
	"path/filepath"
	"text/template"

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
