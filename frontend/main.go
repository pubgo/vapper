package main

import (
	"github.com/pubgo/vapper/frontend/routes"
	"github.com/pubgo/vapper/frontend/stores"
)

func main() {
	//if Document.Get("ReadyState") == js.Undefined {
	//	go run()
	//	return
	//}
	//
	//Document.AddEventListener("DOMContentLoaded", func(event Event) {
	//	go run()
	//})

	run()
}

func run() {
	//vecty.AddStylesheet(dataurl.New([]byte(views.Styles), "text/css").String())

	app := &stores.App{}
	app.Init()

	routes.NewRouter(app).Run()
}


