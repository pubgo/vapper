package routes

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/pubgo/vapper/frontend/stores"
	"github.com/pubgo/vapper/frontend/views"
)

func NewRouter(app *stores.App) *Router {
	return &Router{app: app}
}

type Router struct {
	vecty.Core
	app *stores.App
}

func (t *Router) Render() vecty.ComponentOrHTML {
	return elem.Body(
		views.NewPage(t.app).Route(),
		views.NewNotFound().Route(),
	)
}
