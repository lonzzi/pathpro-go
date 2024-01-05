package engine

import "pathpro-go/pkg/errno"

func newResponse[T any](code errno.ErrCode, msg string, data T) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewResponse[T any](code errno.ErrCode, data T) *Response {
	return newResponse[T](code, code.Error(), data)
}

func NewResponseWithMsg[T any](code errno.ErrCode, msg string, data T) *Response {
	return newResponse[T](code, msg, data)
}

func NewSuccessResponse[T any](data T) *Response {
	return newResponse[T](errno.OK, errno.OK.Error(), data)
}

func NewErrorResponse(code errno.ErrCode) *Response {
	return newResponse[any](code, code.Error(), nil)
}
