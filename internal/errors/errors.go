package errors

import (
	"github.com/pkg/errors"
)

// Define alias
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	SUCCESS                  = NewResponse(200, "ok")
	ERROR                    = NewResponse(500, "fail")
	InvalidParams            = NewResponse(400, "请求参数错误")
	ErrAuthCheckTokenFail    = NewResponse(401, "Token鉴权失败")
	ErrAuthCheckTokenTimeout = NewResponse(401, "Token已超时")
	ErrAuthToken             = NewResponse(401, "Token生成失败")
	ErrAuth                  = NewResponse(401, "Token错误")

	ErrExistUser       = NewResponse(412, "用户名已存在")
	ErrNotFound        = NewResponse(404, "not found")
	ErrMethodNotAllow  = NewResponse(405, "method not allowed")
	ErrInternalServer  = NewResponse(500, "internal server error")
	ErrTooManyRequests = NewResponse(429, "too many requests")
)
