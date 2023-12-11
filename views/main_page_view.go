package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"path"
	"strings"

	"gohttp/core"
	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
	"log"
	"os"
	"strconv"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
	VarsCombo      *widget.Select

	// url
	UrlEntry *widget.Entry

	SendBtn    *widget.Button
	TypeChoice *widget.RadioGroup
	// response
	ResTypeChoice  *widget.Select
	ResBodyLayout  *container.Scroll
	ResBodyCur     *fyne.Widget
	ResStatusCode  *canvas.Text
	ResHeaderEntry *widget.Entry
	ResBodyText    *widget.Entry
	ResBodyImage   *canvas.Image

	// 全局状态底部
	InfoText *widget.Label
	Model    *model.MainModel
}

// load variables file to memory
func (view *MainPageView) loadVarsFile(path string) {
	fmt.Println("加载环境文件:" + path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件错误")
		return
	}
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	for key, value := range result {
		fmt.Println("load vars:" + key + ":" + value.(string))
		view.Model.Vars[key] = value.(string)
	}
	core.GetInstance().FyneApp.SendNotification(fyne.NewNotification("tip", "共加载:"+strconv.Itoa(len(result))+"个变量"))

}

// 打开扫描指定目录
func (view *MainPageView) OpenDir(dirpath string) {
	fmt.Println("1open :" + dirpath)
	view.Model.Files = make([]model.HttpFile, 0)
	view.Model.VarFiles = make([]model.HttpFile, 0)
	dirs, _ := os.ReadDir(dirpath)
	fmt.Printf("radd.%d\n", len(dirs))

	fmt.Printf("files len:%d", len(dirs))
	for _, file := range dirs {
		if !file.Type().IsDir() && strings.HasSuffix(file.Name(), ".http") {
			fmt.Println(file.Name())
			item := &model.HttpFile{Name: file.Name(), Path: path.Join(dirpath, file.Name())}
			view.Model.Files = append(view.Model.Files, *item)
		}
		if !file.Type().IsDir() && strings.HasSuffix(file.Name(), "vars.json") {
			fmt.Println(file.Name())
			item := &model.HttpFile{Name: file.Name(), Path: path.Join(dirpath, file.Name())}
			view.Model.VarFiles = append(view.Model.VarFiles, *item)
		}
	}
	if len(view.Model.Files) == 0 {
		fmt.Println("没有找到文件")
	} else {
		// 保存上次打开目录
		// storage.GetInstance().Cache.Set("last_dir", dirpath, -1)
		// storage.GetInstance().SaveToFile()
		core.GetInstance().FyneApp.Preferences().SetString("last_dir", dirpath)
		fmt.Println("已保存目录到缓存")

	}
	view.FileList.Refresh()
	options := make([]string, len(view.Model.VarFiles))
	for i, val := range view.Model.VarFiles {
		options[i] = val.Name
	}
	view.VarsCombo.SetOptions(options)
	core.GetInstance().FyneApp.SendNotification(fyne.NewNotification("tip", "已打开目录:"+dirpath+" 共"+strconv.Itoa(len(view.Model.Files))+"个http文件"))

	view.InfoText.Text = dirpath
	view.InfoText.Refresh()
}
func (view *MainPageView) Init(window fyne.Window) {

	fmt.Println("main page view init...")

	// main menu
	window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("文件", fyne.NewMenuItem("打开工作区", func() {
			diag := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				if err == nil {
					view.OpenDir(lu.Path())
				} else {
					core.GetInstance().Toast(err.Error())
				}

			}, window)
			diag.Show()
		})),
		fyne.NewMenu("请求", fyne.NewMenuItem("保存修改", func() {
			view.SaveCurHttp()
		})),
		fyne.NewMenu("变量", fyne.NewMenuItem("预览", func() {
			s, _ := json.Marshal(view.Model.Vars)
			dialog := dialog.NewInformation("vars", string(s), core.GetInstance().Window)
			dialog.Show()
		}), fyne.NewMenuItem("new json", func() {
			// TODO 改成判断目录是否存在
			if view.Model.CurFilePath == "" {
				core.GetInstance().Toast("还没有打开工作区，请先打开一个目录作为工作区")
				diag := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
					if err == nil {
						view.OpenDir(lu.Path())
					} else {
						core.GetInstance().Toast(err.Error())
					}

				}, window)
				diag.Show()
				return
			}
			dialog.NewEntryDialog("tip", "将会保存到:"+view.Model.CurFilePath, func(name string) {

			}, window).Show()

		})),
	))

	toolbar := widget.NewToolbar(

		// widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		// 	view.SaveCurHttp()
		// }),
		// widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		// 	diag := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		// 		view.OpenDir(lu.Path())
		// 	}, window)
		// 	diag.Show()
		// }),
		widget.NewToolbarSeparator(),
		// widget.NewToolbarAction(theme.InfoIcon(), func() {
		// 	fmt.Println("加载变量")
		// 	// core.GetInstance().Window.SetMainMenu()

		// }),
	)
	// ---- left--
	view.HttpList = widget.NewList(
		func() int {
			return len(view.Model.Https) + 1
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			if lii == 0 {
				co.(*widget.Label).SetText("+")
			} else {
				if view.Model.Https[lii-1].IsChange {
					co.(*widget.Label).SetText("*" + view.Model.Https[lii-1].Name)

				} else {
					co.(*widget.Label).SetText(view.Model.Https[lii-1].Name)

				}

			}
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
		if view.Model.CurHttpRef.IsChange {
			dialog.NewConfirm("GoHttp", "请求已修改，是否保存?", func(b bool) {
				if b {
					view.SaveCurHttp()
				} else {
					view.Model.CurHttpRef.IsChange = false
				}

			}, window).Show()
			return
		}
		if id == 0 {
			dialog.NewEntryDialog("tip", "名字", func(s string) {
				fmt.Println("name:" + s)
				view.Model.Https = append(view.Model.Https, listener.Http{
					Name:   s,
					Method: "GET",
				})
				view.HttpList.Refresh()
			}, core.GetInstance().Window).Show()
			return
		}

		view.Model.CurHttpRef = &view.Model.Https[id-1]
		fmt.Printf("on select:%d", id-1)
		http := view.Model.Https[id-1]
		view.UrlEntry.SetText(http.Url)

		view.MethodCombo.Selected = http.Method
		view.BodyEntry.SetText(http.Body)
		view.ReqHeaderEntry.SetText(http.HeaderRaw)
		cacheKey := http.Method + " " + http.Url
		data, found := core.GetInstance().Cache.Get(cacheKey)
		if found {
			fmt.Println("load cache response:")
			view.Model.ResBodyBytes = data.([]byte)
			view.ResStatusCode.Text = "当前显示为上次请求缓存"
			view.ResStatusCode.Color = color.RGBA{R: 255, G: 255, B: 0, A: 255}
			view.ResStatusCode.Refresh()
			view.UpdateResponsePreview()
		}
		view.Model.CurHttpRef.IsChange = false

	}
	view.FileList.OnSelected = func(id widget.ListItemID) {
		fmt.Printf("on select:%d", id)
		view.loadFileData(view.Model.Files[id].Path)
	}

	view.Infinite = widget.NewProgressBarInfinite()
	view.Infinite.Hide()
	view.Infinite.Stop()

	varsTip := widget.NewLabel("选择变量文件:")
	// 环境切换等
	view.VarsCombo = widget.NewSelect([]string{}, func(s string) {
		fmt.Println("切换变量:" + s)
	})
	loadvarsBtn := widget.NewButton("读取", func() {
		// 查找打开的那个文件
		for i := 0; i < len(view.Model.VarFiles); i += 1 {
			if view.Model.VarFiles[i].Name == view.VarsCombo.Selected {
				view.loadVarsFile(view.Model.VarFiles[i].Path)
				// TODO 抽取出来
				s, _ := json.Marshal(view.Model.Vars)
				dialog := dialog.NewInformation("vars", string(s), core.GetInstance().Window)
				dialog.Show()
				return
			}
		}
		fmt.Println("没找到变量文件:" + view.VarsCombo.Selected)
	})

	toolLayout := container.NewHBox(varsTip, view.VarsCombo, loadvarsBtn)
	view.MethodCombo = widget.NewSelect([]string{"GET", "POST", "DELETE"}, func(value string) {
		log.Println("Select set to", value)
	})

	view.MethodCombo.Selected = "GET"
	view.SendBtn = widget.NewButton("发送", nil)

	// saveBtn := widget.NewButton("保存", func() { view.SaveCurHttp() })

	view.UrlEntry = widget.NewEntry()
	urlLayout := container.NewBorder(view.Infinite, nil, view.MethodCombo, container.NewHBox(view.SendBtn), view.UrlEntry)

	view.ReqHeaderEntry = widget.NewMultiLineEntry()
	view.BodyEntry = widget.NewMultiLineEntry()
	view.TypeChoice = widget.NewRadioGroup([]string{"text", "form", "x-www-form-urlencoded"}, func(value string) {
		log.Println("Radio set to", value)
	})
	view.TypeChoice.Horizontal = true
	reqDiv := container.NewHSplit(view.ReqHeaderEntry, container.NewBorder(nil, nil, nil, nil, view.BodyEntry))

	// response status
	view.ResStatusCode = canvas.NewText("", color.Black)
	// response body
	view.ResHeaderEntry = widget.NewEntry()
	view.ResTypeChoice = widget.NewSelect([]string{"text", "image"}, func(s string) {
		fmt.Println("切换类型:" + s)
	})
	view.ResTypeChoice.Selected = "text"
	view.ResTypeChoice.OnChanged = func(s string) {
		fmt.Println("select:"+s, "body len:"+strconv.Itoa(len(view.Model.ResBodyBytes)))
		view.UpdateResponsePreview()
	}
	viewBtn := widget.NewButton("下载", func() {
		core.GetInstance().Toast("尚未实现")
	})
	view.ResBodyText = widget.NewMultiLineEntry()
	view.ResBodyImage = canvas.NewImageFromResource(theme.FileImageIcon())
	view.ResBodyImage.FillMode = canvas.ImageFillOriginal
	// view.ResBodyImage.Hide()
	// 自动换行
	// view.ResBodyText.Wrapping = fyne.TextWrapBreak
	// view.ResBodyCur = view.ResBodyImage

	view.ResBodyLayout = container.NewVScroll(view.ResBodyText)
	resDiv := container.NewHSplit(container.NewBorder(view.ResStatusCode, nil, nil, nil, view.ResHeaderEntry),
		container.NewBorder(container.NewHBox(view.ResTypeChoice, viewBtn), nil, nil, nil, view.ResBodyLayout),
	)

	center := container.NewVSplit(reqDiv, resDiv)

	view.InfoText = widget.NewLabel("not open workspace")
	view.Canvas = container.NewBorder(container.NewVBox(toolbar, toolLayout, urlLayout), view.InfoText, container.NewVSplit(view.HttpList, view.FileList), nil, center)
	fmt.Println("set canvas done" + strconv.FormatBool(view.Canvas == nil))

	view.UrlEntry.OnChanged = func(s string) {
		view.Model.CurHttpRef.IsChange = true
		view.HttpList.Refresh()
	}

	view.MethodCombo.OnChanged = func(s string) {
		view.Model.CurHttpRef.IsChange = true
		view.HttpList.Refresh()

	}

	view.ReqHeaderEntry.OnChanged = func(s string) {
		view.Model.CurHttpRef.IsChange = true
		view.HttpList.Refresh()

	}

	view.BodyEntry.OnChanged = func(s string) {
		view.Model.CurHttpRef.IsChange = true
		view.HttpList.Refresh()

	}
	// view.Controller = new(controller.MainController)
	// controller.BindView(view)
	// temp := ""
	// for i := 0; i <= 100; i++ {
	// 	temp += "hello"
	// }
	// view.ResBodyText.SetText(temp)

}

func (view *MainPageView) loadFileData(path string) {

	_, error := os.Stat(path)
	if os.IsNotExist(error) {
		core.GetInstance().Toast(path + " 不存在")
		return
	}

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

func (view *MainPageView) SaveCurHttp() {
	view.Model.CurHttpRef.Method = view.MethodCombo.Selected
	view.Model.CurHttpRef.Url = view.UrlEntry.Text
	view.Model.CurHttpRef.Body = view.BodyEntry.Text
	view.Model.CurHttpRef.IsChange = false
	view.HttpList.Refresh()
	// TODO header save
	core.GetInstance().Toast("save success")
}
func (view *MainPageView) UpdateResponsePreview() {
	s := view.ResTypeChoice.Selected
	if s == "text" {
		if len(view.Model.ResBodyBytes) < 1024 {
			bodyStr := string(view.Model.ResBodyBytes)
			log.Println("boody:", string(view.Model.ResBodyBytes))
			// 设置最大显示长度
			view.ResBodyText.SetText(bodyStr)
		} else {
			view.ResBodyText.SetText("gohttp: 文本太长不显示")
		}
		view.ResBodyLayout.Content = view.ResBodyText
	} else if s == "image" {
		fmt.Print(view.Model.CurPreviewImage)
		if view.Model.CurPreviewImage == nil {
			img, _, err := image.Decode(bytes.NewReader(view.Model.ResBodyBytes))
			if err != nil {
				core.GetInstance().Toast("image parse err" + err.Error())
				return
			}
			view.ResBodyImage = canvas.NewImageFromImage(img)
		}
		view.ResBodyImage.Refresh()

		view.ResBodyLayout.Content = view.ResBodyImage
	}
	view.ResBodyLayout.Refresh()
	view.ResStatusCode.Refresh()
	fmt.Println(view.ResBodyCur)
}

// 打开选择工作区
func (view *MainPageView) selectWorkspace() {

}

// 更新http响应，header,body
func (view *MainPageView) updateResponse() {

}
