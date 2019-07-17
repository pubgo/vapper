package vapper

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/siongui/godom"
	"time"
)

type BaseView struct {
	vecty.Core
	App *Vapper
}

// ReadyStateComplete call when ReadyState is Complete
func (t *BaseView) ReadyStateComplete() {
}

func (t *BaseView) Mount() {
	t.App.Watch(t, func(done chan struct{}) {
		defer close(done)
		vecty.Rerender(t)
	})

	go func() {
		for {
			if godom.Document.Get("readyState").String() != "complete" {
				time.Sleep(time.Millisecond * 10)
				continue
			}

			t.ReadyStateComplete()
			return
		}
	}()
}

func (t *BaseView) Unmount() {
	t.App.Delete(t)
}

func (t *BaseView) Render() vecty.ComponentOrHTML {
	return elem.Body()
}
