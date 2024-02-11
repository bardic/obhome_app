package views

import "fyne.io/fyne/v2"

func CreateFactory(v View) *fyne.Container {
	switch v {
	case NOT_SET:

	case WELCOME:
		w := WelcomeView{}
		return w.Create()
	case SET_TEMP:
		t := SetTempView{}
		return t.Create()
	case SHCEDULE_TEMP:
		break
	}

	return nil
}
