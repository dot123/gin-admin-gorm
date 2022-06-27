package file

import (
	"GameAdmin/internal/models/util"
	"context"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID        uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	Name      string    `gorm:"column:name"`
	Url       string    `gorm:"column:url"`
	Tag       string    `gorm:"column:tag"`
	Key       string    `gorm:"column:key"`
}

func GetFileDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(File))
}
