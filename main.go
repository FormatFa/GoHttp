package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"net/http"

	"gohttp/core"
	"gohttp/viewmanage"

	"github.com/flopp/go-findfont"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func init() {
	//设置中文环境
	fontPaths := findfont.List()
	for _, path := range fontPaths {

		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simhei.ttf") {
			log.Println("设置字体路径=", path)
			os.Setenv("FYNE_FONT", path)
			break
		}
		//mac
		if strings.Contains(path, "Arial Unicode.ttf") {
			log.Println("设置字体路径=", path)
			os.Setenv("FYNE_FONT", path)
			break
		}
		//mac
		if strings.Contains(path, "Apple Braille.ttf") {
			log.Println("设置字体路径=", path)
			os.Setenv("FYNE_FONT", path)
			break
		}

	}
}
func main4() {
	a := app.New()
	w := a.NewWindow("SysTray")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	// w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("hello", fyne.NewMenuItem("start", func() {}))))
	w.ShowAndRun()
}
func main() {

	core.GetInstance().Init()
	viewmanage.OpenMainPageView("")
	core.GetInstance().Start()

}
func main2() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	myWindow.Resize(fyne.NewSize(1024, 768))

	infinite := widget.NewProgressBarInfinite()
	infinite.Hide()
	infinite.Stop()

	resBodyRect := widget.NewLabel("resBodyRect")

	reqLabel := widget.NewLabel("request headers")
	resHeaderUi := widget.NewLabel("response headers")

	reqHeader := widget.NewMultiLineEntry()
	reqHeader.SetText("Content-type: application/json;charset=utf")

	bodyText := widget.NewMultiLineEntry()

	// 顶部
	combo := widget.NewSelect([]string{"GET", "POST", "DELETE", "HEAD", "PUT"}, func(value string) {
		log.Println("Select set to", value)
	})
	combo.Selected = "GET"

	urlEntry := widget.NewEntry()
	urlEntry.Text = "http://baidu.com"
	urlEntry.Resize(fyne.NewSize(600, 50))
	sendBtn := widget.NewButton("send", func() {
		infinite.Show()
		infinite.Start()
		method := combo.Selected
		url := urlEntry.Text
		headerStr := reqHeader.Text
		reqBodyStr := bodyText.Text

		req, _ := http.NewRequest(method, url, strings.NewReader(reqBodyStr))

		for _, line := range strings.Split(headerStr, "\n") {
			kv := strings.SplitN(line, ":", 2)
			if len(kv) == 2 {
				log.Println("add header key=" + kv[0] + " value=" + kv[1])
				req.Header.Add(kv[0], kv[1])
			} else {
				log.Println("discard invalid header ,because len is not 2")
			}
		}

		response, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(response.Body)
			bodyStr := string(body)
			log.Println("boody:", string(body))
			resBodyRect.SetText(bodyStr)
			// res body
			keyValues := make([]string, 0, len(response.Header))
			for k := range response.Header {
				keyValues = append(keyValues, k+": "+response.Header.Get(k))
			}
			resHeadStr := strings.Join(keyValues, "\n")
			log.Println(resHeadStr)
			// TODO muliti values?
			resHeaderUi.SetText(resHeadStr)

		} else {
			log.Println("err=", err)
			dialog.ShowError(err, myWindow)
		}
		log.Print("send: method=" + combo.Selected + " url=" + url)
		infinite.Stop()
		infinite.Hide()
	})
	// top := container.New(layout.NewHBoxLayout(), container.NewStack(urlEntry), sendBtn)

	top := container.NewBorder(infinite, nil, combo, sendBtn, urlEntry)
	// top=container.NewMax()

	typeChoice := widget.NewRadioGroup([]string{"text", "form", "x-www-form-urlencoded"}, func(value string) {
		log.Println("Radio set to", value)
	})
	typeChoice.Horizontal = true

	// when text

	reqRect := container.NewHSplit(reqHeader, container.NewVBox(typeChoice, bodyText))
	// req parameters. headers,body
	// reqHeaders:=
	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	// )
	// tabs.SetTabLocation(container.TabLocationTop)

	// body
	// headers

	headerRect := container.NewVSplit(reqLabel, resHeaderUi)

	resContent := container.NewHSplit(headerRect, resBodyRect)

	// 汇总
	// top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := container.NewVSplit(reqRect, resContent)
	content := container.NewBorder(top, nil, left, nil, middle)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
