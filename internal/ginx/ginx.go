package ginx

import (
	"GameAdmin/internal/errors"
	"GameAdmin/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"strings"
)

const (
	prefix     = "GameAdmin"
	ReqBodyKey = prefix + "/req-body"
	ResBodyKey = prefix + "/res-body"
)

// ResponseData 数据返回结构体
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseFail 返回成功结构体
type ResponseFail struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
	Msg  string `json:"msg"`
}

// RespData 数据返回
func RespData(c *gin.Context, data interface{}) {
	resp := ResponseData{
		Code: 200,
		Data: data,
		Msg:  "success",
	}
	RespJSON(c, http.StatusOK, resp)
}

// RespOk 返回操作成功
func RespOk(c *gin.Context) {
	RespError(c, errors.SUCCESS, "success")
}

// RespJSON 返回JSON数据
func RespJSON(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
}

// GetPage 获取每页数量
func GetPage(c *gin.Context) (pageNum, pageSize int) {
	pageNum, _ = strconv.Atoi(c.Query("page"))
	pageSize, _ = strconv.Atoi(c.Query("limit"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	return
}

// GetToken Get jwt token from header (Authorization: Bearer xxx)
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetBodyData Get body data from context
func GetBodyData(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}

// ParseParamID Param returns the value of the URL param
func ParseParamID(c *gin.Context, key string) uint64 {
	id, err := strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// ParseJSON Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

// ParseQuery Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request query failed: %s", err.Error()))
	}
	return nil
}

func Bind(c *gin.Context, obj interface{}) error {
	if err := c.Bind(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Bind failed: %s", err.Error()))
	}
	return nil
}

// ParseForm Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

// RespError Response error object and parse error status code
func RespError(c *gin.Context, err error, msg string) {
	ctx := c.Request.Context()
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.Err = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	errMsg := ""
	if err = res.Err; err != nil {
		errMsg = err.Error()
		if status := res.Code; status >= 400 && status < 500 {
			logger.WithContext(ctx).Warnf(err.Error())
		} else if status >= 500 {
			logger.WithContext(logger.NewStackContext(ctx, err)).Errorf(err.Error())
		}
	}

	if msg == "" {
		msg = res.Msg
	}

	RespJSON(c, res.Code, ResponseFail{
		Code: res.Code,
		Err:  errMsg,
		Msg:  msg,
	})
}
