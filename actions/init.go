package actions

import (
	"github.com/dave/flux"
)

type Load struct{}

type ConsoleFirstWrite struct{}
type ConsoleToggleClick struct{}
type MinifyToggleClick struct{}

type ShowAllDepsChange struct{ State bool }

type ChangeSplit struct{ Sizes []float64 }
type ChangeFile struct {
	Path string
	Name string
}

type LoadSource struct {
	Source         map[string]map[string]string
	Tags           []string
	CurrentPackage string
	CurrentFile    string
	Save           bool // Save directly after loading? false during initialising, true for load package.
	Update         bool // Update directly after loading?
}

type UserChangedSplit struct{ Sizes []float64 }
type UserChangedText struct {
	Text    string
	Changed bool
}
type UserChangedFile struct{ Name string }
type UserChangedPackage struct{ Path string }

type DownloadClick struct{}
type BuildTags struct{ Tags []string }

type AddFile struct{ Name string }
type AddPackage struct{ Path string }
type DeleteFile struct{ Name string }
type RemovePackage struct{ Path string }

type FormatCode struct{ Then flux.ActionInterface }

// CompileStart compiles the app and injects the js into the iframe
type CompileStart struct{}

type DragEnter struct{}
type DragLeave struct{}
type DragDrop struct {
	Changed map[string]map[string]bool
}

type Dial struct {
	Url     string
	Open    func() flux.ActionInterface
	Message func(interface{}) flux.ActionInterface
	Close   func() flux.ActionInterface
}

type ShareStart struct{}
type ShareOpen struct{}
type ShareMessage struct{ Message interface{} }
type ShareClose struct{}

type DeployStart struct{}
type DeployOpen struct{}
type DeployMessage struct{ Message interface{} }
type DeployClose struct{}

type RequestStart struct {
	Path string // Path to get (for GetRequest and InitialiseRequest)
	Run  bool   // Run after update? (for UpdateRequest)
}
type RequestOpen struct {
	*RequestStart
}
type RequestMessage struct {
	*RequestStart
	Message interface{}
}
type RequestClose struct {
	*RequestStart
}
