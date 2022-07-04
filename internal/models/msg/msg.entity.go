package msg

import (
	"GameAdmin/internal/models/util"
	"GameAdmin/internal/schema"
	"GameAdmin/pkg/structure"
	"GameAdmin/pkg/types"
	"context"
	"gorm.io/gorm"
)

type Notice struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;NOT NULL;"`
	CreatedAt types.Time `gorm:"column:created_at;type:dateTime;comment:'创建时间';"`
	StartTime types.Time `gorm:"column:start_time;type:dateTime;comment:'开始时间';"`
	EndTime   types.Time `gorm:"column:end_time;type:dateTime;comment:'结束时间';"`
	Title     string     `gorm:"column:title;comment:'标题';"`
	Content   string     `gorm:"column:content;comment:'内容';"`
	Operator  string     `gorm:"column:operator;comment:'操作者';"`
}

func GetNoticeDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Notice))
}

type Notices []*Notice

type SchemaNotice schema.Notice

func (a SchemaNotice) ToNotice() *Notice {
	item := new(Notice)
	structure.Copy(a, item)
	return item
}

func (a Notice) ToSchemaNotice() *schema.Notice {
	item := new(schema.Notice)
	structure.Copy(a, item)
	return item
}

func (a Notices) ToSchemaNotices() []*schema.Notice {
	list := make([]*schema.Notice, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaNotice()
	}
	return list
}
