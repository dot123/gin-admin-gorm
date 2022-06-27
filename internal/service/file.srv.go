package service

import (
	"GameAdmin/internal/models/file"
	"GameAdmin/pkg/fileStore"
	"context"
	"errors"
	"github.com/google/wire"
	"mime/multipart"
	"strings"
)

var FileSet = wire.NewSet(wire.Struct(new(FileSrv), "*"))

type FileSrv struct {
	FileRepo *file.FileRepo
	Local    *fileStore.Local
}

// Upload 创建文件上传记录
func (a *FileSrv) Upload(ctx context.Context, file *file.File) error {
	return a.FileRepo.Upload(ctx, file)
}

// FindFile 查找文件切片记录
func (a *FileSrv) FindFile(ctx context.Context, id uint64) (*file.File, error) {
	return a.FileRepo.FindFile(ctx, id)
}

// DeleteFile 删除文件记录
func (a *FileSrv) DeleteFile(ctx context.Context, id uint64) error {
	file, err := a.FindFile(ctx, id)
	if err != nil {
		return err
	}
	if err = a.Local.DeleteFile(file.Key); err != nil {
		return errors.New("文件删除失败")
	}
	return a.FileRepo.DeleteFile(ctx, file.ID)
}

// UploadFile 根据配置文件判断是文件上传到本地或者七牛云
func (a *FileSrv) UploadFile(ctx context.Context, header *multipart.FileHeader, noSave string) (string, error) {
	filePath, key, uploadErr := a.Local.UploadFile(header)
	if uploadErr != nil {
		return "", uploadErr
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := file.File{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		err := a.Upload(ctx, &f)

		return f.Url, err
	}
	return "", nil
}
