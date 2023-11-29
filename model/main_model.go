package model

import (
	"gohttp/listener"
	"gohttp/mvc"
	"net/http"
)

type HttpFile struct {
	Path string
	Name string
}
type MainModel struct {
	mvc.Model
	// response
	response *http.Response
	header   http.Header
	Https    listener.Https
	// 文件夹的列表
	Files       []HttpFile
	CurFilePath string
}

// func (main MainModel) send() {

// }
