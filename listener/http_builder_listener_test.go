package listener

import (
	"fmt"
	"gohttp/parser"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1+2 expected be 3,but %d got", ans)
	}
	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10+-20 expected be -30,but %d got", ans)
	}
}

func TestParse(t *testing.T) {
	input, _ := antlr.NewFileStream("/Users/ljh/Documents/personal/gohttp/demo.http")
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

}
