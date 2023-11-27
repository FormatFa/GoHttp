// Code generated from /Users/ljh/Documents/personal/antrl4test/src/main/java/online/indigo6a/gohttp/GoHttp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoHttp
import "github.com/antlr4-go/antlr/v4"

// GoHttpListener is a complete listener for a parse tree produced by GoHttpParser.
type GoHttpListener interface {
	antlr.ParseTreeListener

	// EnterHttps is called when entering the https production.
	EnterHttps(c *HttpsContext)

	// EnterHttp is called when entering the http production.
	EnterHttp(c *HttpContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterRequest_line is called when entering the request_line production.
	EnterRequest_line(c *Request_lineContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterHttp_version is called when entering the http_version production.
	EnterHttp_version(c *Http_versionContext)

	// EnterHttp_name is called when entering the http_name production.
	EnterHttp_name(c *Http_nameContext)

	// EnterComment_line is called when entering the comment_line production.
	EnterComment_line(c *Comment_lineContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterHeader_field is called when entering the header_field production.
	EnterHeader_field(c *Header_fieldContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// EnterField_value is called when entering the field_value production.
	EnterField_value(c *Field_valueContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterToken is called when entering the token production.
	EnterToken(c *TokenContext)

	// EnterBody_token is called when entering the body_token production.
	EnterBody_token(c *Body_tokenContext)

	// EnterChar is called when entering the char production.
	EnterChar(c *CharContext)

	// EnterBody_char is called when entering the body_char production.
	EnterBody_char(c *Body_charContext)

	// EnterNchar is called when entering the nchar production.
	EnterNchar(c *NcharContext)

	// ExitHttps is called when exiting the https production.
	ExitHttps(c *HttpsContext)

	// ExitHttp is called when exiting the http production.
	ExitHttp(c *HttpContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitRequest_line is called when exiting the request_line production.
	ExitRequest_line(c *Request_lineContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitHttp_version is called when exiting the http_version production.
	ExitHttp_version(c *Http_versionContext)

	// ExitHttp_name is called when exiting the http_name production.
	ExitHttp_name(c *Http_nameContext)

	// ExitComment_line is called when exiting the comment_line production.
	ExitComment_line(c *Comment_lineContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitHeader_field is called when exiting the header_field production.
	ExitHeader_field(c *Header_fieldContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)

	// ExitField_value is called when exiting the field_value production.
	ExitField_value(c *Field_valueContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitToken is called when exiting the token production.
	ExitToken(c *TokenContext)

	// ExitBody_token is called when exiting the body_token production.
	ExitBody_token(c *Body_tokenContext)

	// ExitChar is called when exiting the char production.
	ExitChar(c *CharContext)

	// ExitBody_char is called when exiting the body_char production.
	ExitBody_char(c *Body_charContext)

	// ExitNchar is called when exiting the nchar production.
	ExitNchar(c *NcharContext)
}
