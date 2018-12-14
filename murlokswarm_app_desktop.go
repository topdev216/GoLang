package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

type Hello struct {
	Greeting string
}

func (h *Hello) Render() string {
	return `
<div style="text-align:center;">
	Hello, <span>{{if .Greeting}}{{html .Greeting}}{{else}}World{{end}}!</span>
	<input type="text" placeholder="What is your name?" onchange="OnInputChange" />
</div>
	`
}

func init() {
	app.RegisterComponent(&Hello{})
}

func (h *Hello)OnInputChange(arg app.ChangeArg) {
	h.Greeting = arg.Value
	app.Render(h)
}

func main() {
	app.OnLaunch = func() {
		win := app.NewWindow(app.Window{
			Title: "Hello world",
			Width:800,
			Height:600,
			TitlebarHidden:true,
		})

		win.Mount(&Hello{})
	}

	app.Run()
}
