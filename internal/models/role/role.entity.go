package role

import (
	"GameAdmin/internal/models/util"
	"context"
	"gorm.io/gorm"
)

type Role struct {
	ID       uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserID   uint64 `gorm:"column:user_id;NOT NULL"`
	UserName string `gorm:"column:user_name"`
	Value    string `gorm:"column:value"`
}

func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Role))
}
