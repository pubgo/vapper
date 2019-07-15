package main

import (
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
	vecty.RenderBody(routes.NewRouter(app))
}
