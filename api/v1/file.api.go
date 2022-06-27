package v1

import (
	"GameAdmin/internal/ginx"
	"GameAdmin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var FileSet = wire.NewSet(wire.Struct(new(FileApi), "*"))

type FileApi struct {
	FileSrv *service.FileSrv
}

// @Tags     FileApi
// @Summary  上传文件
// @Accept   multipart/form-data
// @Produce  application/json
// @Param    file  formData  file                 true  "file"
// @Success  200   {object}  ginx.ResponseData{}  "成功结果"
// @Failure  500   {object}  ginx.ResponseFail{}  "失败结果"
// @Router   /uploadFile [post]
func (a *FileApi) UploadFile(c *gin.Context) {
	ctx := c.Request.Context()

	file, err := c.FormFile("file")
	if err != nil {
		ginx.RespError(c, err, "上传文件出错")
		return
	}
	url, err := a.FileSrv.UploadFile(ctx, file, "0")
	if err != nil {
		ginx.RespError(c, err, "上传文件失败")
	} else {
		ginx.RespData(c, url)
	}
}

// @Tags      FileApi
// @Summary   删除文件
// @Accept    application/json
// @Produce   application/json
// @Security  ApiKeyAuth
// @Param     id   path      uint64               true  "id"
// @Success   200  {object}  ginx.ResponseData{}  "成功结果"
// @Failure   500  {object}  ginx.ResponseFail{}  "失败结果"
// @Router    /deleteFile/{id} [delete]
func (a *FileApi) DeleteFile(c *gin.Context) {
	ctx := c.Request.Context()

	id := ginx.ParseParamID(c, "id")
	if err := a.FileSrv.DeleteFile(ctx, id); err != nil {
		ginx.RespError(c, err, "删除文件失败")
	} else {
		ginx.RespOk(c)
	}
}
