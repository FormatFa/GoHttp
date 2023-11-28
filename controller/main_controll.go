package controller

import (
	"fmt"
	"gohttp/listener"
	"gohttp/model"
	"gohttp/mvc"
	"gohttp/views"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MainController struct {
	mvc.BaseController
	Model *model.MainModel
	View  *views.MainPageView
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

		req, _ := http.NewRequest(method, url, strings.NewReader(bodyStr))

		for _, line := range strings.Split(headerStr, "\n") {
			kv := strings.Split(line, ":")
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
			view.ResBodyText.SetText(bodyStr)
			// res body
			keyValues := make([]string, 0, len(response.Header))
			for k := range response.Header {
				keyValues = append(keyValues, k+": "+response.Header.Get(k))
			}
			resHeadStr := strings.Join(keyValues, "\n")
			log.Println(resHeadStr)
			// TODO muliti values?
			view.ResHeaderEntry.SetText(resHeadStr)

		} else {
			log.Println("err=", err)
			// dialog.ShowError(err, myWindow)
		}
		view.Infinite.Stop()
		view.Infinite.Hide()

	}

	// 加载测试数据
	// https := listener.ReadFromFile("/Users/ljh/Documents/personal/gohttp/demo.http")
	// controller.Model.Https = https
}

func (controller *MainController) InitData(path string) {
	https := listener.ReadFromFile("/Users/ljh/Documents/personal/gohttp/demo.http")
	controller.Model.Https = https
	controller.View.HttpList.Refresh()
}
func (controller *MainController) loadFile(path string) {

}
