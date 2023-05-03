package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	_widget := app.NewWindow("Hello World!")

	_widget.SetContent(widget.NewLabel("Welcome home!"))
	_widget.ShowAndRun()
}
