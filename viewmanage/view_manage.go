package viewmanage

import (
	"errors"
	"fmt"
	"gohttp/controller"
	"gohttp/core"
	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
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
		Https: make(listener.Https, 1),
		Files: make([]model.HttpFile, 0),
		Vars:  make(map[string]string),
	}
	view.Init(core.GetInstance().Window)
	// 初始化数据
	mainModel.Https[0].Name = "未命名"
	mainModel.Https[0].Method = "GET"
	mainModel.CurHttpRef = &mainModel.Https[0]
	view.Model = mainModel
	controller := new(controller.MainController)
	controller.Model = mainModel
	controller.BindView(view)

	canvas := view.GetCanvas()
	core.GetInstance().Window.SetContent(canvas)
	lastDir := core.GetInstance().FyneApp.Preferences().String("last_dir")
	if len(lastDir) > 0 {
		fmt.Printf("last dir:%s", lastDir)
		view.OpenDir(lastDir)
	} else {
		fmt.Printf("last dir not found")
	}
	// controller.InitData("")
	return view
}
