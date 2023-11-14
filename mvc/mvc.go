package mvc

import (
	"fyne.io/fyne/v2"
)

type BaseController interface {
	// BindView(view *BaseView)
}
type BaseView interface {
	GetCanvas() fyne.CanvasObject
	Init()
	GetController() *BaseController
	SetController(controller BaseController)
}
