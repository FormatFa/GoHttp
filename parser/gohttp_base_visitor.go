// Code generated from /Users/ljh/Documents/personal/antrl4test/src/main/java/online/indigo6a/gohttp/GoHttp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoHttp
import "github.com/antlr4-go/antlr/v4"

type BaseGoHttpVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoHttpVisitor) VisitHttps(ctx *HttpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitHttp(ctx *HttpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitRequest(ctx *RequestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitRequest_line(ctx *Request_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitUrl(ctx *UrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitHttp_version(ctx *Http_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitHttp_name(ctx *Http_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitComment_line(ctx *Comment_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitMethod(ctx *MethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitHeader_field(ctx *Header_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitField_name(ctx *Field_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitField_value(ctx *Field_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitToken(ctx *TokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitBody_token(ctx *Body_tokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitChar(ctx *CharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitBody_char(ctx *Body_charContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoHttpVisitor) VisitNchar(ctx *NcharContext) interface{} {
	return v.VisitChildren(ctx)
}
