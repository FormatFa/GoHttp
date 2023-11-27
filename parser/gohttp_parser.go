// Code generated from /Users/ljh/Documents/personal/antrl4test/src/main/java/online/indigo6a/gohttp/GoHttp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoHttp
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoHttpParser struct {
	*antlr.BaseParser
}

var GoHttpParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gohttpParserInit() {
	staticData := &GoHttpParserStaticData
	staticData.LiteralNames = []string{
		"", "'HTTP/'", "'GET'", "'HEAD'", "'POST'", "'PUT'", "'DELETE'", "'CONNECT'",
		"'OPTIONS'", "'TRACE'", "' '", "", "", "", "'.'", "'-'", "'#'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "SP", "ALPHA", "DIGIT", "HEXDIG",
		"Dot", "DASH", "Hashtag", "Colon", "CRLF", "OTHER", "WS",
	}
	staticData.RuleNames = []string{
		"https", "http", "request", "request_line", "url", "http_version", "http_name",
		"comment_line", "comment", "method", "header_field", "field_name", "field_value",
		"body", "token", "body_token", "char", "body_char", "nchar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 166, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 5, 0, 40, 8, 0, 10, 0,
		12, 0, 43, 9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 2, 1, 2, 1, 2, 4, 2, 62, 8,
		2, 11, 2, 12, 2, 63, 1, 3, 4, 3, 67, 8, 3, 11, 3, 12, 3, 68, 1, 3, 1, 3,
		1, 3, 3, 3, 74, 8, 3, 1, 3, 5, 3, 77, 8, 3, 10, 3, 12, 3, 80, 9, 3, 1,
		3, 5, 3, 83, 8, 3, 10, 3, 12, 3, 86, 9, 3, 1, 4, 5, 4, 89, 8, 4, 10, 4,
		12, 4, 92, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 5, 8, 105, 8, 8, 10, 8, 12, 8, 108, 9, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 5, 10, 115, 8, 10, 10, 10, 12, 10, 118, 9, 10, 1, 10, 1,
		10, 1, 11, 5, 11, 123, 8, 11, 10, 11, 12, 11, 126, 9, 11, 1, 12, 5, 12,
		129, 8, 12, 10, 12, 12, 12, 132, 9, 12, 1, 13, 1, 13, 1, 14, 4, 14, 137,
		8, 14, 11, 14, 12, 14, 138, 1, 15, 4, 15, 142, 8, 15, 11, 15, 12, 15, 143,
		1, 15, 1, 15, 5, 15, 148, 8, 15, 10, 15, 12, 15, 151, 9, 15, 1, 15, 4,
		15, 154, 8, 15, 11, 15, 12, 15, 155, 3, 15, 158, 8, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 2, 106, 130, 0, 19, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 0, 4, 1, 0, 2, 9, 3,
		0, 10, 12, 14, 17, 19, 19, 2, 0, 10, 12, 14, 19, 3, 0, 11, 12, 14, 16,
		19, 19, 163, 0, 41, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6,
		66, 1, 0, 0, 0, 8, 90, 1, 0, 0, 0, 10, 93, 1, 0, 0, 0, 12, 98, 1, 0, 0,
		0, 14, 100, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 20, 111,
		1, 0, 0, 0, 22, 124, 1, 0, 0, 0, 24, 130, 1, 0, 0, 0, 26, 133, 1, 0, 0,
		0, 28, 136, 1, 0, 0, 0, 30, 157, 1, 0, 0, 0, 32, 159, 1, 0, 0, 0, 34, 161,
		1, 0, 0, 0, 36, 163, 1, 0, 0, 0, 38, 40, 3, 2, 1, 0, 39, 38, 1, 0, 0, 0,
		40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 1, 1, 0,
		0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 3, 4, 2, 0, 45, 3, 1, 0, 0, 0, 46, 47,
		3, 14, 7, 0, 47, 48, 5, 18, 0, 0, 48, 49, 3, 6, 3, 0, 49, 55, 5, 18, 0,
		0, 50, 51, 3, 20, 10, 0, 51, 52, 5, 18, 0, 0, 52, 54, 1, 0, 0, 0, 53, 50,
		1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 5, 18, 0, 0, 59, 61, 3,
		26, 13, 0, 60, 62, 5, 18, 0, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 5, 1, 0, 0, 0, 65, 67, 3, 18,
		9, 0, 66, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 5, 10, 0, 0, 71, 73, 3, 8, 4, 0,
		72, 74, 5, 10, 0, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 78, 1,
		0, 0, 0, 75, 77, 3, 10, 5, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78,
		76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 84, 1, 0, 0, 0, 80, 78, 1, 0, 0,
		0, 81, 83, 5, 18, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 7, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		87, 89, 3, 32, 16, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 9, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93,
		94, 3, 12, 6, 0, 94, 95, 5, 12, 0, 0, 95, 96, 5, 14, 0, 0, 96, 97, 5, 12,
		0, 0, 97, 11, 1, 0, 0, 0, 98, 99, 5, 1, 0, 0, 99, 13, 1, 0, 0, 0, 100,
		101, 5, 16, 0, 0, 101, 102, 3, 16, 8, 0, 102, 15, 1, 0, 0, 0, 103, 105,
		3, 32, 16, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 107, 1,
		0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 17, 1, 0, 0, 0, 108, 106, 1, 0, 0,
		0, 109, 110, 7, 0, 0, 0, 110, 19, 1, 0, 0, 0, 111, 112, 3, 22, 11, 0, 112,
		116, 5, 17, 0, 0, 113, 115, 5, 10, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0,
		0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 3, 24, 12, 0, 120, 21, 1, 0, 0, 0,
		121, 123, 3, 36, 18, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 23, 1, 0, 0, 0, 126, 124, 1,
		0, 0, 0, 127, 129, 3, 32, 16, 0, 128, 127, 1, 0, 0, 0, 129, 132, 1, 0,
		0, 0, 130, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 25, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 133, 134, 3, 30, 15, 0, 134, 27, 1, 0, 0, 0, 135,
		137, 3, 32, 16, 0, 136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136,
		1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 29, 1, 0, 0, 0, 140, 142, 3, 32,
		16, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0,
		143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 5, 18, 0, 0, 146,
		148, 1, 0, 0, 0, 147, 141, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 158, 1, 0, 0, 0, 151, 149, 1, 0,
		0, 0, 152, 154, 3, 32, 16, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0,
		0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157,
		149, 1, 0, 0, 0, 157, 153, 1, 0, 0, 0, 158, 31, 1, 0, 0, 0, 159, 160, 7,
		1, 0, 0, 160, 33, 1, 0, 0, 0, 161, 162, 7, 2, 0, 0, 162, 35, 1, 0, 0, 0,
		163, 164, 7, 3, 0, 0, 164, 37, 1, 0, 0, 0, 17, 41, 55, 63, 68, 73, 78,
		84, 90, 106, 116, 124, 130, 138, 143, 149, 155, 157,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoHttpParserInit initializes any static state used to implement GoHttpParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoHttpParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoHttpParserInit() {
	staticData := &GoHttpParserStaticData
	staticData.once.Do(gohttpParserInit)
}

// NewGoHttpParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoHttpParser(input antlr.TokenStream) *GoHttpParser {
	GoHttpParserInit()
	this := new(GoHttpParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoHttpParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoHttp.g4"

	return this
}

// GoHttpParser tokens.
const (
	GoHttpParserEOF     = antlr.TokenEOF
	GoHttpParserT__0    = 1
	GoHttpParserT__1    = 2
	GoHttpParserT__2    = 3
	GoHttpParserT__3    = 4
	GoHttpParserT__4    = 5
	GoHttpParserT__5    = 6
	GoHttpParserT__6    = 7
	GoHttpParserT__7    = 8
	GoHttpParserT__8    = 9
	GoHttpParserSP      = 10
	GoHttpParserALPHA   = 11
	GoHttpParserDIGIT   = 12
	GoHttpParserHEXDIG  = 13
	GoHttpParserDot     = 14
	GoHttpParserDASH    = 15
	GoHttpParserHashtag = 16
	GoHttpParserColon   = 17
	GoHttpParserCRLF    = 18
	GoHttpParserOTHER   = 19
	GoHttpParserWS      = 20
)

// GoHttpParser rules.
const (
	GoHttpParserRULE_https        = 0
	GoHttpParserRULE_http         = 1
	GoHttpParserRULE_request      = 2
	GoHttpParserRULE_request_line = 3
	GoHttpParserRULE_url          = 4
	GoHttpParserRULE_http_version = 5
	GoHttpParserRULE_http_name    = 6
	GoHttpParserRULE_comment_line = 7
	GoHttpParserRULE_comment      = 8
	GoHttpParserRULE_method       = 9
	GoHttpParserRULE_header_field = 10
	GoHttpParserRULE_field_name   = 11
	GoHttpParserRULE_field_value  = 12
	GoHttpParserRULE_body         = 13
	GoHttpParserRULE_token        = 14
	GoHttpParserRULE_body_token   = 15
	GoHttpParserRULE_char         = 16
	GoHttpParserRULE_body_char    = 17
	GoHttpParserRULE_nchar        = 18
)

// IHttpsContext is an interface to support dynamic dispatch.
type IHttpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllHttp() []IHttpContext
	Http(i int) IHttpContext

	// IsHttpsContext differentiates from other interfaces.
	IsHttpsContext()
}

type HttpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpsContext() *HttpsContext {
	var p = new(HttpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_https
	return p
}

func InitEmptyHttpsContext(p *HttpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_https
}

func (*HttpsContext) IsHttpsContext() {}

func NewHttpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpsContext {
	var p = new(HttpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_https

	return p
}

func (s *HttpsContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpsContext) AllHttp() []IHttpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHttpContext); ok {
			len++
		}
	}

	tst := make([]IHttpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHttpContext); ok {
			tst[i] = t.(IHttpContext)
			i++
		}
	}

	return tst
}

func (s *HttpsContext) Http(i int) IHttpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpContext)
}

func (s *HttpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterHttps(s)
	}
}

func (s *HttpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitHttps(s)
	}
}

func (s *HttpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitHttps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Https() (localctx IHttpsContext) {
	localctx = NewHttpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoHttpParserRULE_https)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoHttpParserHashtag {
		{
			p.SetState(38)
			p.Http()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpContext is an interface to support dynamic dispatch.
type IHttpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Request() IRequestContext

	// IsHttpContext differentiates from other interfaces.
	IsHttpContext()
}

type HttpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpContext() *HttpContext {
	var p = new(HttpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_http
	return p
}

func InitEmptyHttpContext(p *HttpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_http
}

func (*HttpContext) IsHttpContext() {}

func NewHttpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpContext {
	var p = new(HttpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_http

	return p
}

func (s *HttpContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpContext) Request() IRequestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *HttpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterHttp(s)
	}
}

func (s *HttpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitHttp(s)
	}
}

func (s *HttpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitHttp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Http() (localctx IHttpContext) {
	localctx = NewHttpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoHttpParserRULE_http)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Request()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comment_line() IComment_lineContext
	AllCRLF() []antlr.TerminalNode
	CRLF(i int) antlr.TerminalNode
	Request_line() IRequest_lineContext
	Body() IBodyContext
	AllHeader_field() []IHeader_fieldContext
	Header_field(i int) IHeader_fieldContext

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_request
	return p
}

func InitEmptyRequestContext(p *RequestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_request
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) Comment_line() IComment_lineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComment_lineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComment_lineContext)
}

func (s *RequestContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(GoHttpParserCRLF)
}

func (s *RequestContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(GoHttpParserCRLF, i)
}

func (s *RequestContext) Request_line() IRequest_lineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequest_lineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequest_lineContext)
}

func (s *RequestContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *RequestContext) AllHeader_field() []IHeader_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeader_fieldContext); ok {
			len++
		}
	}

	tst := make([]IHeader_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeader_fieldContext); ok {
			tst[i] = t.(IHeader_fieldContext)
			i++
		}
	}

	return tst
}

func (s *RequestContext) Header_field(i int) IHeader_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeader_fieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeader_fieldContext)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (s *RequestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitRequest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoHttpParserRULE_request)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Comment_line()
	}
	{
		p.SetState(47)
		p.Match(GoHttpParserCRLF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Request_line()
	}
	{
		p.SetState(49)
		p.Match(GoHttpParserCRLF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&776192) != 0 {
		{
			p.SetState(50)
			p.Header_field()
		}
		{
			p.SetState(51)
			p.Match(GoHttpParserCRLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(GoHttpParserCRLF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Body()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoHttpParserCRLF {
		{
			p.SetState(60)
			p.Match(GoHttpParserCRLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRequest_lineContext is an interface to support dynamic dispatch.
type IRequest_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode
	Url() IUrlContext
	AllMethod() []IMethodContext
	Method(i int) IMethodContext
	AllHttp_version() []IHttp_versionContext
	Http_version(i int) IHttp_versionContext
	AllCRLF() []antlr.TerminalNode
	CRLF(i int) antlr.TerminalNode

	// IsRequest_lineContext differentiates from other interfaces.
	IsRequest_lineContext()
}

type Request_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequest_lineContext() *Request_lineContext {
	var p = new(Request_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_request_line
	return p
}

func InitEmptyRequest_lineContext(p *Request_lineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_request_line
}

func (*Request_lineContext) IsRequest_lineContext() {}

func NewRequest_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Request_lineContext {
	var p = new(Request_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_request_line

	return p
}

func (s *Request_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Request_lineContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(GoHttpParserSP)
}

func (s *Request_lineContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(GoHttpParserSP, i)
}

func (s *Request_lineContext) Url() IUrlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
}

func (s *Request_lineContext) AllMethod() []IMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodContext); ok {
			len++
		}
	}

	tst := make([]IMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodContext); ok {
			tst[i] = t.(IMethodContext)
			i++
		}
	}

	return tst
}

func (s *Request_lineContext) Method(i int) IMethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *Request_lineContext) AllHttp_version() []IHttp_versionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHttp_versionContext); ok {
			len++
		}
	}

	tst := make([]IHttp_versionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHttp_versionContext); ok {
			tst[i] = t.(IHttp_versionContext)
			i++
		}
	}

	return tst
}

func (s *Request_lineContext) Http_version(i int) IHttp_versionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttp_versionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttp_versionContext)
}

func (s *Request_lineContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(GoHttpParserCRLF)
}

func (s *Request_lineContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(GoHttpParserCRLF, i)
}

func (s *Request_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Request_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Request_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterRequest_line(s)
	}
}

func (s *Request_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitRequest_line(s)
	}
}

func (s *Request_lineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitRequest_line(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Request_line() (localctx IRequest_lineContext) {
	localctx = NewRequest_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoHttpParserRULE_request_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1020) != 0) {
		{
			p.SetState(65)
			p.Method()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(GoHttpParserSP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Url()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoHttpParserSP {
		{
			p.SetState(72)
			p.Match(GoHttpParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoHttpParserT__0 {
		{
			p.SetState(75)
			p.Http_version()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(81)
				p.Match(GoHttpParserCRLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChar() []ICharContext
	Char(i int) ICharContext

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_url
	return p
}

func InitEmptyUrlContext(p *UrlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_url
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) AllChar() []ICharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharContext); ok {
			len++
		}
	}

	tst := make([]ICharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharContext); ok {
			tst[i] = t.(ICharContext)
			i++
		}
	}

	return tst
}

func (s *UrlContext) Char(i int) ICharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharContext)
}

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitUrl(s)
	}
}

func (s *UrlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitUrl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Url() (localctx IUrlContext) {
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoHttpParserRULE_url)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(87)
				p.Char()
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttp_versionContext is an interface to support dynamic dispatch.
type IHttp_versionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Http_name() IHttp_nameContext
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode
	Dot() antlr.TerminalNode

	// IsHttp_versionContext differentiates from other interfaces.
	IsHttp_versionContext()
}

type Http_versionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttp_versionContext() *Http_versionContext {
	var p = new(Http_versionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_http_version
	return p
}

func InitEmptyHttp_versionContext(p *Http_versionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_http_version
}

func (*Http_versionContext) IsHttp_versionContext() {}

func NewHttp_versionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Http_versionContext {
	var p = new(Http_versionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_http_version

	return p
}

func (s *Http_versionContext) GetParser() antlr.Parser { return s.parser }

func (s *Http_versionContext) Http_name() IHttp_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttp_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttp_nameContext)
}

func (s *Http_versionContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(GoHttpParserDIGIT)
}

func (s *Http_versionContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(GoHttpParserDIGIT, i)
}

func (s *Http_versionContext) Dot() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDot, 0)
}

func (s *Http_versionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Http_versionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Http_versionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterHttp_version(s)
	}
}

func (s *Http_versionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitHttp_version(s)
	}
}

func (s *Http_versionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitHttp_version(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Http_version() (localctx IHttp_versionContext) {
	localctx = NewHttp_versionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoHttpParserRULE_http_version)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Http_name()
	}
	{
		p.SetState(94)
		p.Match(GoHttpParserDIGIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(GoHttpParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(GoHttpParserDIGIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttp_nameContext is an interface to support dynamic dispatch.
type IHttp_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsHttp_nameContext differentiates from other interfaces.
	IsHttp_nameContext()
}

type Http_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttp_nameContext() *Http_nameContext {
	var p = new(Http_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_http_name
	return p
}

func InitEmptyHttp_nameContext(p *Http_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_http_name
}

func (*Http_nameContext) IsHttp_nameContext() {}

func NewHttp_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Http_nameContext {
	var p = new(Http_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_http_name

	return p
}

func (s *Http_nameContext) GetParser() antlr.Parser { return s.parser }
func (s *Http_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Http_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Http_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterHttp_name(s)
	}
}

func (s *Http_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitHttp_name(s)
	}
}

func (s *Http_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitHttp_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Http_name() (localctx IHttp_nameContext) {
	localctx = NewHttp_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoHttpParserRULE_http_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(GoHttpParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComment_lineContext is an interface to support dynamic dispatch.
type IComment_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Hashtag() antlr.TerminalNode
	Comment() ICommentContext

	// IsComment_lineContext differentiates from other interfaces.
	IsComment_lineContext()
}

type Comment_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComment_lineContext() *Comment_lineContext {
	var p = new(Comment_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_comment_line
	return p
}

func InitEmptyComment_lineContext(p *Comment_lineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_comment_line
}

func (*Comment_lineContext) IsComment_lineContext() {}

func NewComment_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comment_lineContext {
	var p = new(Comment_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_comment_line

	return p
}

func (s *Comment_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Comment_lineContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(GoHttpParserHashtag, 0)
}

func (s *Comment_lineContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Comment_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comment_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comment_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterComment_line(s)
	}
}

func (s *Comment_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitComment_line(s)
	}
}

func (s *Comment_lineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitComment_line(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Comment_line() (localctx IComment_lineContext) {
	localctx = NewComment_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoHttpParserRULE_comment_line)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(GoHttpParserHashtag)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Comment()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChar() []ICharContext
	Char(i int) ICharContext

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) AllChar() []ICharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharContext); ok {
			len++
		}
	}

	tst := make([]ICharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharContext); ok {
			tst[i] = t.(ICharContext)
			i++
		}
	}

	return tst
}

func (s *CommentContext) Char(i int) ICharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoHttpParserRULE_comment)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(103)
				p.Char()
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_method
	return p
}

func InitEmptyMethodContext(p *MethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_method
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }
func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (s *MethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoHttpParserRULE_method)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1020) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeader_fieldContext is an interface to support dynamic dispatch.
type IHeader_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_name() IField_nameContext
	Colon() antlr.TerminalNode
	Field_value() IField_valueContext
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode

	// IsHeader_fieldContext differentiates from other interfaces.
	IsHeader_fieldContext()
}

type Header_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeader_fieldContext() *Header_fieldContext {
	var p = new(Header_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_header_field
	return p
}

func InitEmptyHeader_fieldContext(p *Header_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_header_field
}

func (*Header_fieldContext) IsHeader_fieldContext() {}

func NewHeader_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Header_fieldContext {
	var p = new(Header_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_header_field

	return p
}

func (s *Header_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Header_fieldContext) Field_name() IField_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Header_fieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(GoHttpParserColon, 0)
}

func (s *Header_fieldContext) Field_value() IField_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_valueContext)
}

func (s *Header_fieldContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(GoHttpParserSP)
}

func (s *Header_fieldContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(GoHttpParserSP, i)
}

func (s *Header_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Header_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Header_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterHeader_field(s)
	}
}

func (s *Header_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitHeader_field(s)
	}
}

func (s *Header_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitHeader_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Header_field() (localctx IHeader_fieldContext) {
	localctx = NewHeader_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoHttpParserRULE_header_field)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Field_name()
	}
	{
		p.SetState(112)
		p.Match(GoHttpParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(113)
				p.Match(GoHttpParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Field_value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_nameContext is an interface to support dynamic dispatch.
type IField_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNchar() []INcharContext
	Nchar(i int) INcharContext

	// IsField_nameContext differentiates from other interfaces.
	IsField_nameContext()
}

type Field_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_nameContext() *Field_nameContext {
	var p = new(Field_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_field_name
	return p
}

func InitEmptyField_nameContext(p *Field_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_field_name
}

func (*Field_nameContext) IsField_nameContext() {}

func NewField_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_nameContext {
	var p = new(Field_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_field_name

	return p
}

func (s *Field_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_nameContext) AllNchar() []INcharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INcharContext); ok {
			len++
		}
	}

	tst := make([]INcharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INcharContext); ok {
			tst[i] = t.(INcharContext)
			i++
		}
	}

	return tst
}

func (s *Field_nameContext) Nchar(i int) INcharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INcharContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INcharContext)
}

func (s *Field_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterField_name(s)
	}
}

func (s *Field_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitField_name(s)
	}
}

func (s *Field_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitField_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Field_name() (localctx IField_nameContext) {
	localctx = NewField_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoHttpParserRULE_field_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&645120) != 0 {
		{
			p.SetState(121)
			p.Nchar()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_valueContext is an interface to support dynamic dispatch.
type IField_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChar() []ICharContext
	Char(i int) ICharContext

	// IsField_valueContext differentiates from other interfaces.
	IsField_valueContext()
}

type Field_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_valueContext() *Field_valueContext {
	var p = new(Field_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_field_value
	return p
}

func InitEmptyField_valueContext(p *Field_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_field_value
}

func (*Field_valueContext) IsField_valueContext() {}

func NewField_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_valueContext {
	var p = new(Field_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_field_value

	return p
}

func (s *Field_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_valueContext) AllChar() []ICharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharContext); ok {
			len++
		}
	}

	tst := make([]ICharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharContext); ok {
			tst[i] = t.(ICharContext)
			i++
		}
	}

	return tst
}

func (s *Field_valueContext) Char(i int) ICharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharContext)
}

func (s *Field_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterField_value(s)
	}
}

func (s *Field_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitField_value(s)
	}
}

func (s *Field_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitField_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Field_value() (localctx IField_valueContext) {
	localctx = NewField_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoHttpParserRULE_field_value)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(127)
				p.Char()
			}

		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Body_token() IBody_tokenContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) Body_token() IBody_tokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBody_tokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBody_tokenContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoHttpParserRULE_body)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Body_token()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChar() []ICharContext
	Char(i int) ICharContext

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_token
	return p
}

func InitEmptyTokenContext(p *TokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_token
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) AllChar() []ICharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharContext); ok {
			len++
		}
	}

	tst := make([]ICharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharContext); ok {
			tst[i] = t.(ICharContext)
			i++
		}
	}

	return tst
}

func (s *TokenContext) Char(i int) ICharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharContext)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitToken(s)
	}
}

func (s *TokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Token() (localctx ITokenContext) {
	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoHttpParserRULE_token)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&777216) != 0) {
		{
			p.SetState(135)
			p.Char()
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBody_tokenContext is an interface to support dynamic dispatch.
type IBody_tokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCRLF() []antlr.TerminalNode
	CRLF(i int) antlr.TerminalNode
	AllChar() []ICharContext
	Char(i int) ICharContext

	// IsBody_tokenContext differentiates from other interfaces.
	IsBody_tokenContext()
}

type Body_tokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBody_tokenContext() *Body_tokenContext {
	var p = new(Body_tokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_body_token
	return p
}

func InitEmptyBody_tokenContext(p *Body_tokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_body_token
}

func (*Body_tokenContext) IsBody_tokenContext() {}

func NewBody_tokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Body_tokenContext {
	var p = new(Body_tokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_body_token

	return p
}

func (s *Body_tokenContext) GetParser() antlr.Parser { return s.parser }

func (s *Body_tokenContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(GoHttpParserCRLF)
}

func (s *Body_tokenContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(GoHttpParserCRLF, i)
}

func (s *Body_tokenContext) AllChar() []ICharContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharContext); ok {
			len++
		}
	}

	tst := make([]ICharContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharContext); ok {
			tst[i] = t.(ICharContext)
			i++
		}
	}

	return tst
}

func (s *Body_tokenContext) Char(i int) ICharContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharContext)
}

func (s *Body_tokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Body_tokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Body_tokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterBody_token(s)
	}
}

func (s *Body_tokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitBody_token(s)
	}
}

func (s *Body_tokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitBody_token(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Body_token() (localctx IBody_tokenContext) {
	localctx = NewBody_tokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoHttpParserRULE_body_token)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&777216) != 0 {
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&777216) != 0) {
				{
					p.SetState(140)
					p.Char()
				}

				p.SetState(143)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(145)
				p.Match(GoHttpParserCRLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(151)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&777216) != 0) {
			{
				p.SetState(152)
				p.Char()
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharContext is an interface to support dynamic dispatch.
type ICharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALPHA() antlr.TerminalNode
	DIGIT() antlr.TerminalNode
	DASH() antlr.TerminalNode
	Hashtag() antlr.TerminalNode
	Colon() antlr.TerminalNode
	OTHER() antlr.TerminalNode
	Dot() antlr.TerminalNode
	SP() antlr.TerminalNode

	// IsCharContext differentiates from other interfaces.
	IsCharContext()
}

type CharContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharContext() *CharContext {
	var p = new(CharContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_char
	return p
}

func InitEmptyCharContext(p *CharContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_char
}

func (*CharContext) IsCharContext() {}

func NewCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharContext {
	var p = new(CharContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_char

	return p
}

func (s *CharContext) GetParser() antlr.Parser { return s.parser }

func (s *CharContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(GoHttpParserALPHA, 0)
}

func (s *CharContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDIGIT, 0)
}

func (s *CharContext) DASH() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDASH, 0)
}

func (s *CharContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(GoHttpParserHashtag, 0)
}

func (s *CharContext) Colon() antlr.TerminalNode {
	return s.GetToken(GoHttpParserColon, 0)
}

func (s *CharContext) OTHER() antlr.TerminalNode {
	return s.GetToken(GoHttpParserOTHER, 0)
}

func (s *CharContext) Dot() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDot, 0)
}

func (s *CharContext) SP() antlr.TerminalNode {
	return s.GetToken(GoHttpParserSP, 0)
}

func (s *CharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterChar(s)
	}
}

func (s *CharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitChar(s)
	}
}

func (s *CharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitChar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Char() (localctx ICharContext) {
	localctx = NewCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoHttpParserRULE_char)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&777216) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBody_charContext is an interface to support dynamic dispatch.
type IBody_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALPHA() antlr.TerminalNode
	DIGIT() antlr.TerminalNode
	DASH() antlr.TerminalNode
	Hashtag() antlr.TerminalNode
	Colon() antlr.TerminalNode
	OTHER() antlr.TerminalNode
	Dot() antlr.TerminalNode
	SP() antlr.TerminalNode
	CRLF() antlr.TerminalNode

	// IsBody_charContext differentiates from other interfaces.
	IsBody_charContext()
}

type Body_charContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBody_charContext() *Body_charContext {
	var p = new(Body_charContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_body_char
	return p
}

func InitEmptyBody_charContext(p *Body_charContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_body_char
}

func (*Body_charContext) IsBody_charContext() {}

func NewBody_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Body_charContext {
	var p = new(Body_charContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_body_char

	return p
}

func (s *Body_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Body_charContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(GoHttpParserALPHA, 0)
}

func (s *Body_charContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDIGIT, 0)
}

func (s *Body_charContext) DASH() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDASH, 0)
}

func (s *Body_charContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(GoHttpParserHashtag, 0)
}

func (s *Body_charContext) Colon() antlr.TerminalNode {
	return s.GetToken(GoHttpParserColon, 0)
}

func (s *Body_charContext) OTHER() antlr.TerminalNode {
	return s.GetToken(GoHttpParserOTHER, 0)
}

func (s *Body_charContext) Dot() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDot, 0)
}

func (s *Body_charContext) SP() antlr.TerminalNode {
	return s.GetToken(GoHttpParserSP, 0)
}

func (s *Body_charContext) CRLF() antlr.TerminalNode {
	return s.GetToken(GoHttpParserCRLF, 0)
}

func (s *Body_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Body_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Body_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterBody_char(s)
	}
}

func (s *Body_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitBody_char(s)
	}
}

func (s *Body_charContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitBody_char(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Body_char() (localctx IBody_charContext) {
	localctx = NewBody_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoHttpParserRULE_body_char)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1039360) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INcharContext is an interface to support dynamic dispatch.
type INcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALPHA() antlr.TerminalNode
	DIGIT() antlr.TerminalNode
	DASH() antlr.TerminalNode
	Hashtag() antlr.TerminalNode
	OTHER() antlr.TerminalNode
	Dot() antlr.TerminalNode

	// IsNcharContext differentiates from other interfaces.
	IsNcharContext()
}

type NcharContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNcharContext() *NcharContext {
	var p = new(NcharContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_nchar
	return p
}

func InitEmptyNcharContext(p *NcharContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoHttpParserRULE_nchar
}

func (*NcharContext) IsNcharContext() {}

func NewNcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NcharContext {
	var p = new(NcharContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoHttpParserRULE_nchar

	return p
}

func (s *NcharContext) GetParser() antlr.Parser { return s.parser }

func (s *NcharContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(GoHttpParserALPHA, 0)
}

func (s *NcharContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDIGIT, 0)
}

func (s *NcharContext) DASH() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDASH, 0)
}

func (s *NcharContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(GoHttpParserHashtag, 0)
}

func (s *NcharContext) OTHER() antlr.TerminalNode {
	return s.GetToken(GoHttpParserOTHER, 0)
}

func (s *NcharContext) Dot() antlr.TerminalNode {
	return s.GetToken(GoHttpParserDot, 0)
}

func (s *NcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.EnterNchar(s)
	}
}

func (s *NcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoHttpListener); ok {
		listenerT.ExitNchar(s)
	}
}

func (s *NcharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoHttpVisitor:
		return t.VisitNchar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoHttpParser) Nchar() (localctx INcharContext) {
	localctx = NewNcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoHttpParserRULE_nchar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&645120) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
