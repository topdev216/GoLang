package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
)

func main() {
	gl.StartDriver(func(driver gxui.Driver) {
		theme := dark.CreateTheme(driver)
		font, _ := driver.CreateFont(gxfont.Default, 14)
		color := gxui.Color{R: 1, G: 1, B: 1, A: 0.75}

		layout := theme.CreateLinearLayout()
		layout.SetSizeMode(gxui.Fill)

		label := func(name string) {
			l := theme.CreateLabel()
			l.SetText(name)
			l.SetColor(color)
			l.SetFont(font)
			layout.AddChild(l)
		}
		textBox := func(name string, width int) {
			b := theme.CreateTextBox()
			b.SetText(name)
			b.SetDesiredWidth(width)
			layout.AddChild(b)
		}
		button := func(name string, action func()) {
			b := theme.CreateButton()
			b.SetType(gxui.ToggleButton)
			b.SetText(name)
			b.OnClick(func(gxui.MouseEvent) { action() })
			layout.AddChild(b)
		}

		label("Host:")
		textBox("127.0.0.1", 60)
		label("Port:")
		textBox("3306", 40)
		label("Database:")
		textBox("", 60)
		label("Username:")
		textBox("root", 40)
		label("Password:")
		textBox("", 60)
		button("Connect", func() {
			//
		})

		layout.SetDirection(gxui.LeftToRight)
		layout.SetHorizontalAlignment(gxui.AlignLeft)
		layout.SetVerticalAlignment(gxui.AlignTop)

		window := theme.CreateWindow(800, 500, "Hello")
		window.SetScale(1)
		window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray10))
		window.AddChild(layout)
		window.OnClose(driver.Terminate)
		window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
	})
}
