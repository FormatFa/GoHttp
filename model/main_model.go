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
	response   *http.Response
	header     http.Header
	Https      listener.Https
	CurHttpRef *listener.Http
	// 文件夹的列表
	Files []HttpFile
	// TODO 当前文件是否改变
	// 当前工作区
	CurWorkPath string
	// 当前编辑的文件
	CurEditPath string
	Vars        map[string]string

	CurPreviewImage image.Image

	// 变量文件列表
	VarFiles []HttpFile

	ResBodyBytes []byte
	// request config
	// timeout second
	Timeout int
}

// func (main MainModel) send() {

// }
