package model

import (
	"gohttp/listener"
	"gohttp/mvc"
	"image"
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
	Vars        map[string]string

	CurPreviewImage image.Image

	// 变量文件列表
	VarFiles []HttpFile

	ResBodyBytes []byte
}

// func (main MainModel) send() {

// }
