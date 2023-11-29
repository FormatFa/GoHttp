package views

import (
	"fmt"
	"path"
	"strings"

	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainPageView struct {
	mvc.BaseView

	// mainModel      *model.MainModel
	HttpList       *widget.List
	FileList       *widget.List
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

func (view *MainPageView) Init(window fyne.Window) {
	fmt.Println("main page view init...")
	toolbar := widget.NewToolbar(

		widget.NewToolbarAction(theme.FileIcon(), func() {
			diag := dialog.NewFileOpen(func(f fyne.URIReadCloser, e error) {
				fmt.Printf("open:%s", f.URI())
				https := listener.ReadFromIo(f)
				view.Model.Https = https
			}, window)
			diag.Show()
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			diag := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				fmt.Println("1open :" + lu.Path())
				view.Model.Files = make([]model.HttpFile, 0)
				dirs, err := os.ReadDir(lu.Path())
				fmt.Printf("radd.%d\n", len(dirs))

				fmt.Printf("files len:%d", len(dirs))
				for _, file := range dirs {
					if !file.Type().IsDir() && strings.HasSuffix(file.Name(), ".http") {
						fmt.Println(file.Name())
						item := &model.HttpFile{Name: file.Name(), Path: path.Join(lu.Path(), file.Name())}
						view.Model.Files = append(view.Model.Files, *item)
					}
				}
				if len(view.Model.Files) == 0 {
					fmt.Println("没有找到文件")
				}
				view.FileList.Refresh()
			}, window)
			diag.Show()
		}),
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

	view.FileList = widget.NewList(func() int { return len(view.Model.Files) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		}, func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(view.Model.Files[lii].Name)

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
	view.FileList.OnSelected = func(id widget.ListItemID) {
		fmt.Printf("on select:%d", id)
		view.loadFileData(view.Model.Files[id].Path)
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

	view.Canvas = container.NewBorder(container.NewVBox(toolbar, top), nil, container.NewVSplit(view.HttpList, view.FileList), nil, center)
	fmt.Println("set canvas done" + strconv.FormatBool(view.Canvas == nil))

	// view.Controller = new(controller.MainController)
	// controller.BindView(view)

}

func (view *MainPageView) loadFileData(path string) {
	https := listener.ReadFromFile(path)
	fmt.Printf("load result,len=%d\n", len(https))
	view.Model.Https = https
	view.HttpList.Refresh()
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
