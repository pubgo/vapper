package main


import (
	"github.com/pubgo/vapper/router"
	"github.com/pubgo/vapper/stores"
	"github.com/pubgo/vapper/views"
)

func main() {
	app := &stores.App{}
	app.Init()

	r := router.New()
	// Use HandleFunc to add routes.
	r.HandleFunc("/", views.NewPage(app).Handle)
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()
}

