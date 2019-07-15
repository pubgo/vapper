package actions

import (
	"github.com/dave/flux"
	"github.com/dave/services"
)

type Load struct{}

type UserChangedTextAction struct {
	Text string
}

type ChangeTextAction struct {
	Text string
}

type Send struct{ Message services.Message }
type Dial struct {
	Url     string
	Open    func() flux.ActionInterface
	Message func(interface{}) flux.ActionInterface
	Close   func() flux.ActionInterface
}
