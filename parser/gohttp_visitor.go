// Code generated from /Users/ljh/Documents/personal/antrl4test/src/main/java/online/indigo6a/gohttp/GoHttp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoHttp
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoHttpParser.
type GoHttpVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoHttpParser#https.
	VisitHttps(ctx *HttpsContext) interface{}

	// Visit a parse tree produced by GoHttpParser#http.
	VisitHttp(ctx *HttpContext) interface{}

	// Visit a parse tree produced by GoHttpParser#request.
	VisitRequest(ctx *RequestContext) interface{}

	// Visit a parse tree produced by GoHttpParser#request_line.
	VisitRequest_line(ctx *Request_lineContext) interface{}

	// Visit a parse tree produced by GoHttpParser#url.
	VisitUrl(ctx *UrlContext) interface{}

	// Visit a parse tree produced by GoHttpParser#http_version.
	VisitHttp_version(ctx *Http_versionContext) interface{}

	// Visit a parse tree produced by GoHttpParser#http_name.
	VisitHttp_name(ctx *Http_nameContext) interface{}

	// Visit a parse tree produced by GoHttpParser#comment_line.
	VisitComment_line(ctx *Comment_lineContext) interface{}

	// Visit a parse tree produced by GoHttpParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by GoHttpParser#method.
	VisitMethod(ctx *MethodContext) interface{}

	// Visit a parse tree produced by GoHttpParser#header_field.
	VisitHeader_field(ctx *Header_fieldContext) interface{}

	// Visit a parse tree produced by GoHttpParser#field_name.
	VisitField_name(ctx *Field_nameContext) interface{}

	// Visit a parse tree produced by GoHttpParser#field_value.
	VisitField_value(ctx *Field_valueContext) interface{}

	// Visit a parse tree produced by GoHttpParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by GoHttpParser#token.
	VisitToken(ctx *TokenContext) interface{}

	// Visit a parse tree produced by GoHttpParser#body_token.
	VisitBody_token(ctx *Body_tokenContext) interface{}

	// Visit a parse tree produced by GoHttpParser#char.
	VisitChar(ctx *CharContext) interface{}

	// Visit a parse tree produced by GoHttpParser#body_char.
	VisitBody_char(ctx *Body_charContext) interface{}

	// Visit a parse tree produced by GoHttpParser#nchar.
	VisitNchar(ctx *NcharContext) interface{}
}
