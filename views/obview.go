package views

type View int

const (
	NOT_SET View = -1
	WELCOME      = iota
	SET_TEMP
	SHCEDULE_TEMP
)

type IOBView interface {
	GoNext() (View, error)
	GoPrev() (View, error)
}

type OBView struct {
	Next  View
	Prev  View
	Valid bool
}
