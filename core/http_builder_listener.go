package core

import (
	"fmt"
	"gohttp/parser"
)

type Header struct {
	Key   string
	Value string
}
type Http struct {
	Method  string
	Url     string
	Body    string
	Headers []Header
}
type Https []Http
type HttpBuilderListener struct {
	*parser.BaseGoHttpListener
	https     Https
	curHttp   *Http
	curHeader Header
}

func (listener *HttpBuilderListener) EnterHttps(ctx *parser.HttpsContext) {
	fmt.Println("enter https:")
	listener.https = make(Https, 10)
}
func (s *HttpBuilderListener) EnterHttp(ctx *parser.HttpContext) {
	s.curHttp = new(Http)
}
func (s *HttpBuilderListener) ExitHttp(ctx *parser.HttpContext) {
	s.https = append(s.https, *s.curHttp)
}
func (s *HttpBuilderListener) EnterMethod(ctx *parser.MethodContext) {
	s.curHttp.Method = ctx.GetText()
}

func (s *HttpBuilderListener) EnterUrl(ctx *parser.UrlContext) {
	fmt.Println("url:" + ctx.GetText())
	s.curHttp.Url = ctx.GetText()
}

// 解析成https对象
func (h *HttpBuilderListener) GetHttps() Https {
	return h.https

}
func NewHttpBuilderListener() *HttpBuilderListener {
	return new(HttpBuilderListener)
}
func Add(a int, b int) int {
	return a + b
}
