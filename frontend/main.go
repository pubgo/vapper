package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/pubgo/vapper/frontend/routes"
	"github.com/pubgo/vapper/frontend/stores"
	"github.com/pubgo/vapper/frontend/views"
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

	r := routes.New()
	r.ForceHashURL = false
	//r.ShouldInterceptLinks=true

	//// Use HandleFunc to add routes.
	r.HandleFunc("/greet/{name}", func(context *routes.Context) {
		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		fmt.Printf("Hello, %s\n", context.Params["name"])
		vecty.RenderBody(p)
	})

	r.HandleFunc("/", func(context *routes.Context) {
		fmt.Println(context, "context")
		vecty.RenderBody(p)
	})

	pt := Window.Get("location").Get("pathname").String()
	if r.CanNavigate(pt) {
		r.Navigate(pt)
	} else {
		r.Navigate("/")
	}
}
