package viewmanage

import (
	"errors"
	"fmt"
	"gohttp/controller"
	"gohttp/core"
	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
	"gohttp/storage"
	"gohttp/views"
)

func openView(viewId string) {

	fmt.Printf("open view point:%p\n", &viewId)
	// var view *mvc.BaseView = nil
	if viewId == "main" {
		OpenMainPageView()
	} else {
		errors.New("unknow view id" + viewId)
	}
	// canvas := view.GetCanvas()
	// application.window.SetContent(canvas)

	// application.CurVie
	// application.window.SetContent(view.GetCanvas())

}

func OpenMainPageView() mvc.BaseView {
	view := new(views.MainPageView)
	mainModel := &model.MainModel{
		Https: make(listener.Https, 5),
		Files: make([]model.HttpFile, 5),
		Vars:  make(map[string]string),
	}
	view.Init(core.GetInstance().Window)
	view.Model = mainModel
	controller := new(controller.MainController)
	controller.Model = mainModel
	controller.BindView(view)

	canvas := view.GetCanvas()
	core.GetInstance().Window.SetContent(canvas)

	if x, found := storage.GetInstance().Cache.Get("last_dir"); found {
		fmt.Printf("last dir:%s", x)
		view.OpenDir(x.(string))
	} else {
		fmt.Printf("last dir not found")
	}
	// controller.InitData("")
	return view
}
