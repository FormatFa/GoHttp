package listener

import (
	"fmt"
	"gohttp/parser"
	"io"

	"github.com/antlr4-go/antlr/v4"
)

type Header struct {
	Key   string
	Value string
}
type Http struct {
	Name      string
	Method    string
	Url       string
	Body      string
	Headers   []Header
	HeaderRaw string
	IsChange  bool
}
type Https []Http
type HttpBuilderListener struct {
	*parser.BaseGoHttpListener
	https     Https
	curHttp   *Http
	curHeader *Header
}

func ReadFromIo(reader io.Reader) Https {
	input := antlr.NewIoStream(reader)
	lexer := parser.NewGoHttpLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGoHttpParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Https()
	listener := NewHttpBuilderListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	https := listener.GetHttps()
	fmt.Println("parse done")
	for _, val := range https {
		fmt.Println(val.Method + " " + val.Url)
		fmt.Println(val.Headers)
		fmt.Println(val.Body)
	}
	return https
}
func ReadFromFile(path string) Https {
	input, _ := antlr.NewFileStream(path)
	lexer := parser.NewGoHttpLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGoHttpParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Https()
	listener := NewHttpBuilderListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	https := listener.GetHttps()
	for _, val := range https {
		fmt.Println(val.Method + " " + val.Url)
		fmt.Println(val.Headers)
		fmt.Println(val.Body)
	}
	return https
}
func (listener *HttpBuilderListener) EnterHttps(ctx *parser.HttpsContext) {
	fmt.Println("enter https:")
	listener.https = make(Https, 0)
}
func (s *HttpBuilderListener) EnterComment(ctx *parser.CommentContext) {
	s.curHttp.Name = ctx.GetText()
}

func (s *HttpBuilderListener) EnterHttp(ctx *parser.HttpContext) {
	s.curHttp = new(Http)
	s.curHttp.Headers = make([]Header, 0)
	s.curHttp.HeaderRaw = ""
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
func (s *HttpBuilderListener) EnterHeader_field(ctx *parser.Header_fieldContext) {
	fmt.Println("field:" + ctx.GetText())
	s.curHttp.HeaderRaw += ctx.GetText() + "\n"
}

func (s *HttpBuilderListener) EnterField_name(ctx *parser.Field_nameContext) {
	s.curHeader = new(Header)
	s.curHeader.Key = ctx.GetText()
}
func (s *HttpBuilderListener) EnterField_value(ctx *parser.Field_valueContext) {
	s.curHeader.Value = ctx.GetText()
}
func (s *HttpBuilderListener) ExitHeader_field(ctx *parser.Header_fieldContext) {
	s.curHttp.Headers = append(s.curHttp.Headers, *s.curHeader)
}
func (s *HttpBuilderListener) EnterBody(ctx *parser.BodyContext) {
	s.curHttp.Body = ctx.GetText()
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
