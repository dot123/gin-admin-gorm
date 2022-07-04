package msg

import (
	"GameAdmin/internal/models/util"
	"GameAdmin/internal/schema"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var MsgSet = wire.NewSet(wire.Struct(new(MsgRepo), "*"))

type MsgRepo struct {
	DB *gorm.DB
}

// GetNotices 获得公告列表
func (a *MsgRepo) GetNotices(ctx context.Context, pageNum int, pageSize int) (*schema.NoticeQueryResult, error) {
	var list Notices

	total, err := util.GetPages(GetNoticeDB(ctx, a.DB), &list, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	var result schema.NoticeQueryResult
	result.List = list.ToSchemaNotices()
	result.Total = total

	return &result, nil
}

// AddNotice 新建公告
func (a *MsgRepo) AddNotice(ctx context.Context, params *SchemaNotice) error {
	if err := GetNoticeDB(ctx, a.DB).Create(params.ToNotice()).Error; err != nil {
		return err
	}
	return nil
}

// UpdateNotice 更新公告
func (a *MsgRepo) UpdateNotice(ctx context.Context, params *SchemaNotice) error {
	if err := GetNoticeDB(ctx, a.DB).Where("`id`=?", params.ID).Save(params.ToNotice()).Error; err != nil {
		return err
	}
	return nil
}

// DeleteNotice 删除公告
func (a *MsgRepo) DeleteNotice(ctx context.Context, id uint64) error {
	if err := GetNoticeDB(ctx, a.DB).Where("`id`=?", id).Delete(new(Notice)).Error; err != nil {
		return err
	}
	return nil
}

// GetNoticeByID 获取公告
func (a *MsgRepo) GetNoticeByID(ctx context.Context, id uint64) (*Notice, error) {
	result := new(Notice)
	if err := GetNoticeDB(ctx, a.DB).Where("`id`=?", id).Take(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
