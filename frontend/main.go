package main

import (
	"github.com/gopherjs/vecty"
	"github.com/pubgo/vapper/frontend/stores"
	"github.com/pubgo/vapper/frontend/views"
	"github.com/vincent-petithory/dataurl"
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
	vecty.AddStylesheet(dataurl.New([]byte(styles), "text/css").String())
	app := &stores.App{}
	app.Init()

	vecty.RenderBody(views.NewPage(app))
	//vecty.RenderBody(routes.NewRouter(app))
}

var styles = `
	html, body {
		height: 100%;
	}
	.editor {
		height: 100%;
		width: 100%;
	}
	.split {
		height: 100%;
		width: 100%;
	}
	.gutter {
		height: 100%;
		background-color: #eee;
		background-repeat: no-repeat;
		background-position: 50%;
	}
	.gutter.gutter-horizontal {
		cursor: col-resize;
		background-image:  url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAeCAYAAADkftS9AAAAIklEQVQoU2M4c+bMfxAGAgYYmwGrIIiDjrELjpo5aiZeMwF+yNnOs5KSvgAAAABJRU5ErkJggg==')
	}
	.split {
		-webkit-box-sizing: border-box;
		-moz-box-sizing: border-box;
		box-sizing: border-box;
	}
	.split, .gutter.gutter-horizontal {
		float: left;
	}
`
