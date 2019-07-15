package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/pubgo/vapper/frontend/stores"
	"github.com/pubgo/vapper/frontend/views"
	"github.com/pubgo/vrouter"
	. "github.com/siongui/godom"
	"github.com/vincent-petithory/dataurl"
)

func main() {
	if Document.Get("ReadyState") == js.Undefined {
		go run()
		return
	}

	Document.AddEventListener("DOMContentLoaded", func(event Event) {
		go run()
	})
}

func run() {
	vecty.AddStylesheet(dataurl.New([]byte(views.Styles), "text/css").String())
	app := &stores.App{}
	app.Init()
	//vecty.RenderBody(routes.NewRouter(app))
	p := views.NewPage(app)

	r := vrouter.New()
	r.ForceHashURL=false
	//r.ShouldInterceptLinks=true

	//// Use HandleFunc to add routes.
	//r.HandleFunc("/greet/{name}", func(context *vrouter.Context) {
	//	// The handler for this route simply grabs the name parameter
	//	// from the map of params and says hello.
	//	fmt.Printf("Hello, %s\n", context.Params["name"])
	//	vecty.RenderBody(p)
	//})
	//r.HandleFunc("/", func(context *vrouter.Context) {
	//	fmt.Println(context,"context")
	//
	//	vecty.RenderBody(p)
	//})
	//// You must call Start in order to start listening for changes
	//// in the url and trigger the appropriate handler function.
	//r.Start()
	//r.Navigate("/")
}
