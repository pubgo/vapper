package vapper

import (
	"fmt"
	"github.com/dave/flux"
	"github.com/gopherjs/gopherjs/js"
	dom "github.com/siongui/godom"
	"reflect"
)

type Vapper struct {
	dispatcher flux.DispatcherInterface
	watcher    flux.WatcherInterface
	notifier   flux.NotifierInterface

	routes *Router
	cfg    reflect.Value
}

func (a *Vapper) Init() {

	n := flux.NewNotifier()
	a.notifier = n
	a.watcher = n

	a.dispatcher = flux.NewDispatcher(
		// Notifier:
		a.notifier,

		// Stores:
	)
}

func (a *Vapper) Dispatch(action flux.ActionInterface) chan struct{} {
	return a.dispatcher.Dispatch(action)
}

func (a *Vapper) Watch(key interface{}, f func(done chan struct{})) {
	a.watcher.Watch(key, f)
}

func (a *Vapper) Delete(key interface{}) {
	a.watcher.Delete(key)
}

func (a *Vapper) Fail(err error) {
	js.Global.Call("alert", err.Error())
}

func (a *Vapper) Debug(message ...interface{}) {
	js.Global.Get("console").Call("log", message...)
}

func (a *Vapper) Log(message ...interface{}) {
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

func (a *Vapper) Logf(format string, args ...interface{}) {
	a.Log(fmt.Sprintf(format, args...))
}

func requestAnimationFrame() {
	c := make(chan struct{})
	js.Global.Call("requestAnimationFrame", func() { close(c) })
	<-c
}
