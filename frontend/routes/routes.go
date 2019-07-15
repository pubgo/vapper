package routes

import (
	"github.com/pubgo/vapper/frontend/views"
)

func (t *Router) render() {
	t.route("/", views.NewPage(t.app), true)
}
