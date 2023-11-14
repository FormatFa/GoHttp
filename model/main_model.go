package model

import "net/http"

type MainModel struct {
	// response
	response *http.Response
	header   http.Header
}

func (main MainModel) send() {

}
