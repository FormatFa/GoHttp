package core

import (
	"errors"
	"fmt"
	"gohttp/controller"
	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
	"gohttp/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// 1. 渲染显示内
type View struct {
	Canvas *fyne.Container
}

func (view *View) Init() {

}

// func (view BaseView) setController(controller BaseController) {
// 	view.Controller = controller
// }

type Application struct {
	// CurView fyne.CanvasObject
	CurView mvc.BaseView
	fyneApp fyne.App
	window  fyne.Window
}

func (application Application) Start() {

	application.fyneApp = app.New()
	application.window = application.fyneApp.NewWindow("Container")
	application.window.Resize(fyne.NewSize(1024, 768))
	application.openView("main")
	// con := new(controller.MainController)

	// application.openView(initView)
	// text := canvas.NewText("hello", color.Black)
	//  application.window.SetContent(container.NewCenter(text))

	// application.openView(intiView)
	application.window.ShowAndRun()

}
func (application Application) OpenMainPageView() mvc.BaseView {
	view := new(views.MainPageView)
	mainModel := &model.MainModel{
		Https: make(listener.Https, 5),
		Files: make([]model.HttpFile, 5),
	}
	view.Init(application.window)
	view.Model = mainModel
	controller := new(controller.MainController)
	controller.Model = mainModel
	controller.BindView(view)

	canvas := view.GetCanvas()
	application.window.SetContent(canvas)

	// controller.InitData("")
	return view
}
func (application Application) openView(viewId string) {

	fmt.Printf("open view point:%p\n", &viewId)
	// var view *mvc.BaseView = nil
	if viewId == "main" {
		application.OpenMainPageView()
	} else {
		errors.New("unknow view id" + viewId)
	}
	// canvas := view.GetCanvas()
	// application.window.SetContent(canvas)

	// application.CurVie
	// application.window.SetContent(view.GetCanvas())

}
