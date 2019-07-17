package vapper

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type BaseView struct {
	vecty.Core
	App *Vapper
}

func (t *BaseView) Mount() {
	t.App.Watch(t, func(done chan struct{}) {
		defer close(done)
		vecty.Rerender(t)
	})
}

func (t *BaseView) Unmount() {
	t.App.Delete(t)
}

func (t *BaseView) Render() vecty.ComponentOrHTML {
	return elem.Body()
}
