package views

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SetTempView struct {
	View OBView
	UI   *fyne.Container
}

func (v SetTempView) GoNext() (View, error) {
	if v.View.Valid {
		return v.View.Next, nil
	}

	return NOT_SET, errors.New("not valid")
}

func (v SetTempView) GoPrev() (View, error) {
	if v.View.Prev != NOT_SET {
		return v.View.Prev, nil
	}

	return NOT_SET, errors.New("not valid")
}

func (v *SetTempView) Create() *fyne.Container {
	v.View = OBView{
		Next:  NOT_SET,
		Prev:  WELCOME,
		Valid: false,
	}
	input := widget.NewEntry()
	input.SetPlaceHolder("Thermostat IP")
	ipForm := container.NewVBox()
	ipForm.Add(input)
	ipForm.Add(widget.NewButton("Save", func() {
		//onClick(input.Text, nil)

	}),
	)

	v.UI = ipForm
	return v.UI
}
