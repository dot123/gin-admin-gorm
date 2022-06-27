package v1

import (
	"GameAdmin/internal/errors"
	"GameAdmin/internal/ginx"
	"GameAdmin/internal/schema"
	"GameAdmin/internal/service"
	"GameAdmin/pkg/logger"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

var UserSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSrv *service.UserSrv
}

// @Tags      UserApi
// @Summary   获取用户信息
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Success   200  {object}  ginx.ResponseData{data=schema.UserData}  "成功结果"
// @Failure   500  {object}  ginx.ResponseFail{}                      "失败结果"
// @Router    /user/info [get]
func (a *UserApi) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	roles := jwt.ExtractClaims(c)
	userName := roles["userName"].(string)
	avatar, err := a.UserSrv.GetUserAvatar(ctx, userName)
	if err != nil {
		ginx.RespError(c, err, "")
		return
	}
	arrRole, err := a.UserSrv.GetRoles(ctx, userName)
	if err != nil {
		ginx.RespError(c, err, "")
		return
	}

	data := schema.UserData{Roles: *arrRole, Introduction: "", Avatar: avatar, Name: userName}
	ginx.RespData(c, &data)
}

// @Tags      UserApi
// @Summary   获取用户列表
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     page   query     int                                             true  "页"   limit  default(1)
// @Param     limit  query     int                                             true  "数量"  default(10)
// @Param     name   query     string                                          true  "相似用户名"
// @Success   200    {object}  ginx.ResponseData{data=schema.UserQueryResult}  "成功结果"
// @Failure   500    {object}  ginx.ResponseFail{}                             "失败结果"
// @Router    /user/list [get]
func (a *UserApi) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	name := c.Query("name")
	pageNum, pageSize := ginx.GetPage(c)
	result, err := a.UserSrv.GetUsers(ctx, pageNum, pageSize, name)
	if err != nil {
		ginx.RespError(c, err, "获取用户信息失败")
		return
	}
	ginx.RespData(c, result)
}

// @Tags      UserApi
// @Summary   新建用户
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     data  query     schema.UserDataParam  true  "请求参数"
// @Success   200   {object}  ginx.ResponseData{}   "成功结果"
// @Failure   500   {object}  ginx.ResponseFail{}   "失败结果"
// @Router    /user [post]
func (a *UserApi) AddUser(c *gin.Context) {
	ctx := c.Request.Context()

	errMsg := errors.InvalidParams
	var params schema.UserDataParam
	if err := ginx.Bind(c, &params); err != nil {
		logger.WithContext(ctx).Error(err)
	} else {
		roles := jwt.ExtractClaims(c)
		createdBy := roles["userName"].(string)

		if !a.UserSrv.ExistUserByName(ctx, params.Username) {
			if err := a.UserSrv.AddUser(ctx, &params, createdBy); err != nil {
				errMsg = errors.ERROR
			} else {
				errMsg = errors.SUCCESS
			}
		} else {
			errMsg = errors.ErrExistUser
		}
	}

	ginx.RespError(c, errMsg, "")
}

// @Tags      UserApi
// @Summary   修改用户
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     data  query     schema.UserDataParam  true  "请求参数"
// @Success   200   {object}  ginx.ResponseData{}   "成功结果"
// @Failure   500   {object}  ginx.ResponseFail{}   "失败结果"
// @Router    /user [put]
func (a *UserApi) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	errMsg := errors.InvalidParams
	var params schema.UserDataParam
	if err := ginx.Bind(c, &params); err != nil {
		logger.WithContext(ctx).Error(err)
	} else {
		roles := jwt.ExtractClaims(c)
		modifiedBy := roles["userName"].(string)
		if err := a.UserSrv.UpdateUser(ctx, &params, modifiedBy); err != nil {
			errMsg = errors.ERROR
		} else {
			errMsg = errors.SUCCESS
		}
	}
	ginx.RespError(c, errMsg, "")
}

// @Tags      UserApi
// @Summary   删除用户
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     id   path      uint64               true  "id"
// @Success   200  {object}  ginx.ResponseData{}  "成功结果"
// @Failure   500  {object}  ginx.ResponseFail{}  "失败结果"
// @Router    /user/{id} [delete]
func (a *UserApi) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	id := ginx.ParseParamID(c, "id")
	if err := a.UserSrv.DeleteUser(ctx, id); err != nil {
		ginx.RespError(c, err, "")
	} else {
		ginx.RespOk(c)
	}
}
