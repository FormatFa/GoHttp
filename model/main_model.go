package model

import (
	"gohttp/listener"
	"gohttp/mvc"
	"net/http"
)

type MainModel struct {
	mvc.Model
	// response
	response *http.Response
	header   http.Header
	Https    listener.Https
}

// func (main MainModel) send() {

// }
