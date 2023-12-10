package controller

import (
	"fmt"
	"gohttp/core"
	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
	"gohttp/views"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/valyala/fasttemplate"
)

type MainController struct {
	mvc.BaseController
	Model *model.MainModel
	View  *views.MainPageView
}

func replaceVar(s string, vars map[string]interface{}) string {
	return fasttemplate.New(s, "{{", "}}").ExecuteString(vars)
}
func (controller *MainController) BindView(view *views.MainPageView) {
	// controller.view = view
	controller.View = view

	view.SendBtn.OnTapped = func() {
		fmt.Println("my click.")
		method := view.MethodCombo.Selected
		url := view.UrlEntry.Text
		headerStr := view.ReqHeaderEntry.Text
		bodyStr := view.BodyEntry.Text
		fmt.Println("" + method + " " + url + " " + headerStr + " " + bodyStr)

		view.Infinite.Show()
		view.Infinite.Start()

		varsParam := make(map[string]interface{}, len(view.Model.Vars))
		for k, v := range view.Model.Vars {
			varsParam[k] = v
		}
		url = replaceVar(url, varsParam)
		bodyStr = replaceVar(bodyStr, varsParam)
		req, _ := http.NewRequest(method, url, strings.NewReader(bodyStr))

		for _, line := range strings.Split(headerStr, "\n") {
			kv := strings.Split(line, ":")
			if len(kv) == 2 {
				log.Println("add header key=" + kv[0] + " value=" + kv[1])
				req.Header.Add(replaceVar(kv[0], varsParam), replaceVar(kv[1], varsParam))
			} else {
				log.Println("discard invalid header ,because len is not 2")
			}
		}

		response, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(response.Body)
			// 根据content-type 相应类型决定
			// TODO: 如果大小大于某个值就不保存到内存
			controller.Model.ResBodyBytes = body
			id := method + " " + url
			if len(body) < 3000 {
				core.GetInstance().Cache.Set(id, body, time.Hour)
			} else {
				core.GetInstance().Cache.Delete(id)
			}

			// res body
			keyValues := make([]string, 0, len(response.Header))
			for k := range response.Header {
				keyValues = append(keyValues, k+": "+response.Header.Get(k))
			}
			resHeadStr := strings.Join(keyValues, "\n")
			log.Println(resHeadStr)
			// TODO muliti values?
			view.ResHeaderEntry.SetText(resHeadStr)
			view.ResStatusCode.Text = response.Status
			if response.StatusCode >= 200 && response.StatusCode < 300 {
				view.ResStatusCode.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}

			} else if response.StatusCode >= 400 {
				view.ResStatusCode.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

			} else {
				view.ResStatusCode.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

			}
			// 情况之前生成的数据
			controller.Model.CurPreviewImage = nil

		} else {
			log.Println("err=", err)
			// dialog.ShowError(err, myWindow)
		}
		view.Infinite.Stop()
		view.Infinite.Hide()
		view.UpdateResponsePreview()

	}

	// 加载测试数据
	// https := listener.ReadFromFile("/Users/ljh/Documents/personal/gohttp/demo.http")
	// controller.Model.Https = https
}

func (controller *MainController) LoadHttpFile(f fyne.URIReadCloser) {
	https := listener.ReadFromIo(f)
	controller.Model.Https = https
	controller.View.HttpList.Refresh()
}
func (controller *MainController) loadFile(path string) {

}
