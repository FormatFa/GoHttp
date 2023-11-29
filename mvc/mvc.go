package mvc

import (
	"fyne.io/fyne/v2"
)

type Model interface {
}
type BaseController interface {
	// BindView(view *BaseView)
}
type BaseView interface {
	GetCanvas() fyne.CanvasObject
	Init(window fyne.Window)
	GetController() *BaseController
	SetController(controller BaseController)
}
