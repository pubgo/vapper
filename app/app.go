package app

import (
	"fmt"
	"github.com/dave/flux"
	"github.com/gopherjs/gopherjs/js"
	dom "github.com/siongui/godom"
)

type App struct {
	Dispatcher flux.DispatcherInterface
	Watcher    flux.WatcherInterface
	Notifier   flux.NotifierInterface
}

func (a *App) Init() {

	n := flux.NewNotifier()
	a.Notifier = n
	a.Watcher = n

	a.Dispatcher = flux.NewDispatcher(
		// Notifier:
		a.Notifier,
		// Stores:
	)
}

func (a *App) Dispatch(action flux.ActionInterface) chan struct{} {
	return a.Dispatcher.Dispatch(action)
}

func (a *App) Watch(key interface{}, f func(done chan struct{})) {
	a.Watcher.Watch(key, f)
}

func (a *App) Delete(key interface{}) {
	a.Watcher.Delete(key)
}

func (a *App) Fail(err error) {
	// TODO: improve this
	js.Global.Call("alert", err.Error())
}

func (a *App) Debug(message ...interface{}) {
	js.Global.Get("console").Call("log", message...)
}

func (a *App) Log(message ...interface{}) {
	m := dom.Document.GetElementById("message")
	if len(message) == 0 {
		m.SetInnerHTML("")
		return
	}
	s := fmt.Sprint(message[0])
	if m.InnerHTML() != s {
		requestAnimationFrame()
		m.SetInnerHTML(s)
		requestAnimationFrame()
	}
	js.Global.Get("console").Call("log", message...)
}

func (a *App) Logf(format string, args ...interface{}) {
	a.Log(fmt.Sprintf(format, args...))
}

func requestAnimationFrame() {
	c := make(chan struct{})
	js.Global.Call("requestAnimationFrame", func() { close(c) })
	<-c
}
