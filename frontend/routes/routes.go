package routes

import (
	"github.com/pubgo/vapper/frontend/views"
)

func (t *Router) render() {
	t.routes = t.routes[:0]

	t.route("/", views.NewPage(t.app), true)
}
