package engine

import "pathpro-go/pkg/errno"

func newResponse(code errno.ErrCode, msg string, data any) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewResponse(code errno.ErrCode, data any) *Response {
	return newResponse(code, code.Error(), data)
}

func NewResponseWithMsg(code errno.ErrCode, msg string, data any) *Response {
	return newResponse(code, msg, data)
}

func NewSuccessResponse(data any) *Response {
	return newResponse(errno.OK, errno.OK.Error(), data)
}

func NewErrorResponse(code errno.ErrCode) *Response {
	return newResponse(code, code.Error(), nil)
}
