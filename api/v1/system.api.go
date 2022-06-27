package v1

import (
	"GameAdmin/internal/ginx"
	"GameAdmin/internal/service"
	"GameAdmin/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SystemSet = wire.NewSet(wire.Struct(new(SystemApi), "*"))

type SystemApi struct {
	SystemSrv *service.SystemSrv
}

// @Tags      SystemApi
// @Summary   重启系统
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Success   200  {object}  ginx.ResponseData{}  "成功结果"
// @Failure   500  {object}  ginx.ResponseFail{}  "失败结果"
// @Router    /system/reloadSystem [get]
func (a *SystemApi) ReloadSystem(c *gin.Context) {
	if err := helper.Reload(); err != nil {
		ginx.RespError(c, err, "重启系统失败")
	} else {
		ginx.RespOk(c)
	}
}

// @Tags      SystemApi
// @Summary   服务器状态
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Success   200  {object}  ginx.ResponseData{data=helper.Server}  "成功结果"
// @Failure   500  {object}  ginx.ResponseFail{}                    "失败结果"
// @Router    /system/serverInfo [get]
func (a *SystemApi) GetServerInfo(c *gin.Context) {
	ctx := c.Request.Context()
	s, err := a.SystemSrv.GetServerInfo(ctx)
	if err != nil {
		ginx.RespError(c, err, "获取服务器状态失败")
	} else {
		ginx.RespData(c, s)
	}
}
