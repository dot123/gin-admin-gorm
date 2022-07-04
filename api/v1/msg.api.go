package v1

import (
	"GameAdmin/internal/ginx"
	"GameAdmin/internal/models/msg"
	"GameAdmin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var MsgSet = wire.NewSet(wire.Struct(new(MsgApi), "*"))

type MsgApi struct {
	MsgSrv *service.MsgSrv
}

// @Tags     MsgApi
// @Summary  获取公告列表
// @Accept   application/json
// @Produce  application/json
// @Param    page   query     int                                               true  "页"   limit  default(1)
// @Param    limit  query     int                                               true  "数量"  default(10)
// @Success  200    {object}  ginx.ResponseData{data=schema.NoticeQueryResult}  "成功结果"
// @Failure  500    {object}  ginx.ResponseFail{}                               "失败结果"
// @Router   /msg/notice [get]
func (a *MsgApi) GetNotices(c *gin.Context) {
	ctx := c.Request.Context()
	pageNum, pageSize := ginx.GetPage(c)
	result, err := a.MsgSrv.GetNotices(ctx, pageNum, pageSize)
	if err != nil {
		ginx.RespError(c, err, "获取公告列表失败")
	} else {
		ginx.RespData(c, result)
	}
}

// @Tags      MsgApi
// @Summary   新建公告
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     data  query     msg.SchemaNotice     true  "请求参数"
// @Success   200   {object}  ginx.ResponseData{}  "成功结果"
// @Failure   500   {object}  ginx.ResponseFail{}  "失败结果"
// @Router    /msg/notice [post]
func (a *MsgApi) AddNotice(c *gin.Context) {
	ctx := c.Request.Context()
	var params msg.SchemaNotice
	if err := ginx.Bind(c, &params); err != nil {
		ginx.RespError(c, err, "新建公告失败")
		return
	}

	if err := a.MsgSrv.AddNotice(ctx, &params); err != nil {
		ginx.RespError(c, err, "新建公告失败")
	} else {
		ginx.RespOk(c)
	}
}

// @Tags      MsgApi
// @Summary   更新公告
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     data  query     msg.SchemaNotice     true  "请求参数"
// @Success   200   {object}  ginx.ResponseData{}  "成功结果"
// @Failure   500   {object}  ginx.ResponseFail{}  "失败结果"
// @Router    /msg/notice [put]
func (a *MsgApi) UpdateNotice(c *gin.Context) {
	ctx := c.Request.Context()
	var params msg.SchemaNotice
	if err := ginx.Bind(c, &params); err != nil {
		ginx.RespError(c, err, "更新公告失败")
		return
	}

	if err := a.MsgSrv.UpdateNotice(ctx, &params); err != nil {
		ginx.RespError(c, err, "更新公告失败")
	} else {
		ginx.RespOk(c)
	}
}

// @Tags      MsgApi
// @Summary   删除公告
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     id   path      uint64               true  "公告id"
// @Success   200  {object}  ginx.ResponseData{}  "成功结果"
// @Failure   500  {object}  ginx.ResponseFail{}  "失败结果"
// @Router    /msg/notice/{id} [delete]
func (a *MsgApi) DeleteNotice(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "id")
	if err := a.MsgSrv.DeleteNotice(ctx, id); err != nil {
		ginx.RespError(c, err, "删除公告失败")
	} else {
		ginx.RespOk(c)
	}
}
