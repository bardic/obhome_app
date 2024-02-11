package views

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type WelcomeView struct {
	View OBView
	UI   *fyne.Container
}

func (v WelcomeView) GoNext() (View, error) {
	if v.View.Valid {
		return v.View.Next, nil
	}

	return NOT_SET, errors.New("not valid")
}

func (v WelcomeView) GoPrev() (View, error) {
	if v.View.Prev != NOT_SET {
		return v.View.Prev, nil
	}

	return NOT_SET, errors.New("not valid")
}

func (v *WelcomeView) Create() *fyne.Container {
	v.View = OBView{Next: SET_TEMP, Prev: NOT_SET}

	ipForm := container.NewVBox()
	w := widget.NewLabel("Welcome to OBHome")
	ipForm.Add(w)
	ipForm.Add(widget.NewButton("Save", func() {
		
	}),
	)

	return ipForm
}
