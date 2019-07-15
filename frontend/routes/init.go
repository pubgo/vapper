package routes

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/pubgo/vapper/frontend/stores"
	router "marwan.io/vecty-router"
)

func NewRouter(app *stores.App) *Router {
	return &Router{app: app}
}

type Router struct {
	vecty.Core

	routes []vecty.MarkupOrChild
	app    *stores.App
}

func (t *Router) Run() {
	vecty.RenderBody(t)

	//t.app.Watch(nil, func(done chan struct{}) {
	//	defer close(done)
	//	vecty.Rerender(t)
	//})
	//t.app.Dispatch(&actions.Load{})
}

func (t *Router) Render() vecty.ComponentOrHTML {
	t.render()
	
	return elem.Body(
		t.routes...
	)
}

func (t *Router) route(pattern string, c vecty.Component, exactMatch bool) {
	t.routes = append(t.routes, router.NewRoute(pattern, c, router.NewRouteOpts{ExactMatch: exactMatch}))
}
