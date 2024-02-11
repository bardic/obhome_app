package main

import (
	"obhome/views"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("OBHome")
	content := views.CreateFactory(views.WELCOME)
	
	w.SetContent(content)
	w.ShowAndRun()
}
