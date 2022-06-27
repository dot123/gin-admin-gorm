package file

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var FileSet = wire.NewSet(wire.Struct(new(FileRepo), "*"))

type FileRepo struct {
	DB *gorm.DB
}

// Upload 创建文件上传记录
func (a *FileRepo) Upload(ctx context.Context, file *File) error {
	err := GetFileDB(ctx, a.DB).Create(file).Error
	return err
}

// FindFile 删除文件切片记录
func (a *FileRepo) FindFile(ctx context.Context, id uint64) (*File, error) {
	var file File

	err := GetFileDB(ctx, a.DB).Where("`id` = ?", id).Take(&file).Error
	return &file, err
}

// DeleteFile 删除文件记录
func (a *FileRepo) DeleteFile(ctx context.Context, id uint64) error {
	err := GetFileDB(ctx, a.DB).Where("`id` = ?", id).Unscoped().Delete(new(File)).Error
	return err
}
