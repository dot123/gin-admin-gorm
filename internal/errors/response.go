package errors

import "fmt"

// ResponseError 定义响应错误
type ResponseError struct {
	Code int    // 错误码
	Msg  string // 错误消息
	Err  error  // 响应错误
}

func (r *ResponseError) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return r.Msg
}

func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

func WrapResponse(err error, code int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code: code,
		Msg:  fmt.Sprintf(msg, args...),
		Err:  err,
	}
	return res
}

func Wrap400Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 400, msg, args...)
}

func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 500, msg, args...)
}

func NewResponse(code int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code: code,
		Msg:  fmt.Sprintf(msg, args...),
	}
	return res
}

func New400Response(msg string, args ...interface{}) error {
	return NewResponse(400, msg, args...)
}

func New500Response(msg string, args ...interface{}) error {
	return NewResponse(500, msg, args...)
}
