// Code generated from /Users/ljh/Documents/personal/antrl4test/src/main/java/online/indigo6a/gohttp/GoHttp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoHttp
import "github.com/antlr4-go/antlr/v4"

// BaseGoHttpListener is a complete listener for a parse tree produced by GoHttpParser.
type BaseGoHttpListener struct{}

var _ GoHttpListener = &BaseGoHttpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoHttpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoHttpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoHttpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoHttpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHttps is called when production https is entered.
func (s *BaseGoHttpListener) EnterHttps(ctx *HttpsContext) {}

// ExitHttps is called when production https is exited.
func (s *BaseGoHttpListener) ExitHttps(ctx *HttpsContext) {}

// EnterHttp is called when production http is entered.
func (s *BaseGoHttpListener) EnterHttp(ctx *HttpContext) {}

// ExitHttp is called when production http is exited.
func (s *BaseGoHttpListener) ExitHttp(ctx *HttpContext) {}

// EnterRequest is called when production request is entered.
func (s *BaseGoHttpListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production request is exited.
func (s *BaseGoHttpListener) ExitRequest(ctx *RequestContext) {}

// EnterRequest_line is called when production request_line is entered.
func (s *BaseGoHttpListener) EnterRequest_line(ctx *Request_lineContext) {}

// ExitRequest_line is called when production request_line is exited.
func (s *BaseGoHttpListener) ExitRequest_line(ctx *Request_lineContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseGoHttpListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseGoHttpListener) ExitUrl(ctx *UrlContext) {}

// EnterHttp_version is called when production http_version is entered.
func (s *BaseGoHttpListener) EnterHttp_version(ctx *Http_versionContext) {}

// ExitHttp_version is called when production http_version is exited.
func (s *BaseGoHttpListener) ExitHttp_version(ctx *Http_versionContext) {}

// EnterHttp_name is called when production http_name is entered.
func (s *BaseGoHttpListener) EnterHttp_name(ctx *Http_nameContext) {}

// ExitHttp_name is called when production http_name is exited.
func (s *BaseGoHttpListener) ExitHttp_name(ctx *Http_nameContext) {}

// EnterComment_line is called when production comment_line is entered.
func (s *BaseGoHttpListener) EnterComment_line(ctx *Comment_lineContext) {}

// ExitComment_line is called when production comment_line is exited.
func (s *BaseGoHttpListener) ExitComment_line(ctx *Comment_lineContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseGoHttpListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseGoHttpListener) ExitComment(ctx *CommentContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseGoHttpListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseGoHttpListener) ExitMethod(ctx *MethodContext) {}

// EnterHeader_field is called when production header_field is entered.
func (s *BaseGoHttpListener) EnterHeader_field(ctx *Header_fieldContext) {}

// ExitHeader_field is called when production header_field is exited.
func (s *BaseGoHttpListener) ExitHeader_field(ctx *Header_fieldContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BaseGoHttpListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BaseGoHttpListener) ExitField_name(ctx *Field_nameContext) {}

// EnterField_value is called when production field_value is entered.
func (s *BaseGoHttpListener) EnterField_value(ctx *Field_valueContext) {}

// ExitField_value is called when production field_value is exited.
func (s *BaseGoHttpListener) ExitField_value(ctx *Field_valueContext) {}

// EnterBody is called when production body is entered.
func (s *BaseGoHttpListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseGoHttpListener) ExitBody(ctx *BodyContext) {}

// EnterToken is called when production token is entered.
func (s *BaseGoHttpListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseGoHttpListener) ExitToken(ctx *TokenContext) {}

// EnterBody_token is called when production body_token is entered.
func (s *BaseGoHttpListener) EnterBody_token(ctx *Body_tokenContext) {}

// ExitBody_token is called when production body_token is exited.
func (s *BaseGoHttpListener) ExitBody_token(ctx *Body_tokenContext) {}

// EnterChar is called when production char is entered.
func (s *BaseGoHttpListener) EnterChar(ctx *CharContext) {}

// ExitChar is called when production char is exited.
func (s *BaseGoHttpListener) ExitChar(ctx *CharContext) {}

// EnterBody_char is called when production body_char is entered.
func (s *BaseGoHttpListener) EnterBody_char(ctx *Body_charContext) {}

// ExitBody_char is called when production body_char is exited.
func (s *BaseGoHttpListener) ExitBody_char(ctx *Body_charContext) {}

// EnterNchar is called when production nchar is entered.
func (s *BaseGoHttpListener) EnterNchar(ctx *NcharContext) {}

// ExitNchar is called when production nchar is exited.
func (s *BaseGoHttpListener) ExitNchar(ctx *NcharContext) {}
