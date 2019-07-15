package views

import (
	"fmt"
	"github.com/dave/splitter"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/pubgo/vapper/frontend/actions"
	"github.com/pubgo/vapper/frontend/stores"
	router "marwan.io/vecty-router"
	 "honnef.co/go/js/dom"
)

type Page struct {
	vecty.Core
	app *stores.App

	split *splitter.Split
}

func NewPage(app *stores.App) *Page {
	v := &Page{
		app: app,
	}
	return v
}

func (t *Page) Route() *router.Route {
	return router.NewRoute("/a", t, router.NewRouteOpts{ExactMatch: true})
}

func (t *Page) Mount() {
	t.app.Watch(t, func(done chan struct{}) {
		defer close(done)
		vecty.Rerender(t)
	})

	t.split = splitter.New("split")
	t.split.Init(
		js.S{"#left", "#right"},
		js.M{"sizes": []float64{50, 50}},
	)
}

func (t *Page) Unmount() {
	t.app.Delete(t)
}

func (t *Page) Render() vecty.ComponentOrHTML {
	fmt.Println(router.GetNamedVar(t), "hello")
	fmt.Println(dom.GetWindow().Document().DocumentURI(), "hello")

	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("container-fluid", "p-0", "split", "split-horizontal"),
			),
			t.renderLeft(),
			t.renderRight(),
		),
	)
}

func (t *Page) renderLeft() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			prop.ID("left"),
			vecty.Class("split"),
		),
		NewEditor(t.app, "html-editor", "html", t.app.Editor.Html(), true, func(value string) {
			t.app.Dispatch(&actions.UserChangedTextAction{
				Text: value,
			})
		}),
	)
}

func (t *Page) renderRight() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			prop.ID("right"),
			vecty.Class("split"),
		),
		NewEditor(t.app, "code-editor", "golang", t.app.Editor.Code(), false, nil),
	)
}
