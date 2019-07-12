package codegen

import (
	"bytes"
	"github.com/pubgo/errors"
	"text/template"
)

// ClientGoMain returns the Go code for the main function for the wasm app
func ClientGoMain(appPath string) string {
	b := new(bytes.Buffer)
	defer b.Reset()

	data := map[string]string{"AppPath": appPath}
	errors.Panic(clMainTemplate.Execute(b, data))
	return string(b.Bytes())
}

var clMainTemplate = template.Must(template.New("cl").Parse(clMainTemplateStr))

const clMainTemplateStr = `// build +js,wasm
package main

import (
	"{{.AppPath}}/routes"
	"github.com/pubgo/vrouter"
	"github.com/gopherjs/vecty"
)

func main() {
	c := make(chan struct{}, 0)
	// Create a new Router object
	r := router.New()
	//r.ShouldInterceptLinks = true
	// Use HandleFunc to add routes.
	r.HandleFunc("/", func(context *router.Context) {

		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		vecty.SetTitle("Factor: Home")
		vecty.RenderBody(&routes.Index{})
	})
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()
	<-c
}
`
