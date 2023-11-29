package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"strings"

	"net/http"

	"gohttp/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := new(core.Application)

	application.Start()

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
	combo := widget.NewSelect([]string{"GET", "POST", "DELETE"}, func(value string) {
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
