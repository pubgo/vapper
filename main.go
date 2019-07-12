package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/pubgo/vapper/app"
	"github.com/pubgo/vapper/router"
	"github.com/pubgo/vapper/views"
)

func main() {
	_app := &app.App{}
	_app.Init()

	v1 := views.NewPage(_app)

	r := router.New()
	// Use HandleFunc to add routes.
	r.HandleFunc("/", func(context *router.Context) {
		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		fmt.Printf("Hello, %s\n", context.Params["name"])
		vecty.RenderBody(v1)
	})

	r.HandleFunc("/a2/{name}", func(context *router.Context) {
		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		fmt.Printf("Hello, %s\n", context.Params["name"])
		vecty.RenderBody(v1)
	})
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()

}
