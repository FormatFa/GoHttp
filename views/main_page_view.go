package views

import (
	"fmt"

	"gohttp/model"
	"gohttp/mvc"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainPageView struct {
	mvc.BaseView

	mainModel      model.MainModel
	HttpList       *widget.List
	Canvas         fyne.CanvasObject
	Controller     *mvc.BaseController
	Infinite       *widget.ProgressBarInfinite
	ReqHeaderEntry *widget.Entry
	BodyEntry      *widget.Entry
	MethodCombo    *widget.Select

	// url
	UrlEntry *widget.Entry

	SendBtn    *widget.Button
	TypeChoice *widget.RadioGroup
	// response
	ResHeaderEntry *widget.Entry
	ResBodyText    *widget.Label
	Model          *model.MainModel
}

func (view *MainPageView) Init() {
	fmt.Println("main page view init...")
	toolbar := widget.NewToolbar(

		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {}),
	)
	// ---- left--
	view.HttpList = widget.NewList(
		func() int {
			return len(view.Model.Https)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(view.Model.Https[lii].Name)
		},
	)

	view.HttpList.OnSelected = func(id widget.ListItemID) {
		fmt.Printf("on select:%d", id)
		http := view.Model.Https[id]
		view.UrlEntry.SetText(http.Url)
		view.MethodCombo.Selected = http.Method
		view.BodyEntry.SetText(http.Body)
		view.ReqHeaderEntry.SetText(http.HeaderRaw)
	}

	view.Infinite = widget.NewProgressBarInfinite()
	view.Infinite.Hide()
	view.Infinite.Stop()

	view.MethodCombo = widget.NewSelect([]string{"GET", "POST", "DELETE"}, func(value string) {
		log.Println("Select set to", value)
	})
	view.MethodCombo.Selected = "GET"
	view.SendBtn = widget.NewButton("send", nil)

	view.UrlEntry = widget.NewEntry()
	top := container.NewBorder(view.Infinite, nil, view.MethodCombo, view.SendBtn, view.UrlEntry)

	view.ReqHeaderEntry = widget.NewMultiLineEntry()
	view.BodyEntry = widget.NewMultiLineEntry()
	view.TypeChoice = widget.NewRadioGroup([]string{"text", "form", "x-www-form-urlencoded"}, func(value string) {
		log.Println("Radio set to", value)
	})
	view.TypeChoice.Horizontal = true
	reqDiv := container.NewHSplit(view.ReqHeaderEntry, container.NewVBox(view.TypeChoice, view.BodyEntry))

	// response body
	view.ResHeaderEntry = widget.NewEntry()
	view.ResBodyText = widget.NewLabel("")
	resDiv := container.NewHSplit(view.ResHeaderEntry, view.ResBodyText)

	center := container.NewVSplit(reqDiv, resDiv)

	view.Canvas = container.NewBorder(container.NewVBox(toolbar, top), nil, view.HttpList, nil, center)
	fmt.Println("set canvas done" + strconv.FormatBool(view.Canvas == nil))

	// view.Controller = new(controller.MainController)
	// controller.BindView(view)

}
func (view *MainPageView) GetCanvas() fyne.CanvasObject {
	fmt.Println("call get canvas" + strconv.FormatBool(view.Canvas == nil))
	return view.Canvas
}
func (view *MainPageView) setController(controller *mvc.BaseController) {
	view.Controller = controller
}
func (view *MainPageView) GetController() *mvc.BaseController {
	// view.Controller = controller
	return view.Controller
}

// 更新http响应，header,body
func (view *MainPageView) updateResponse() {

}
