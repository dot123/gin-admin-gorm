package service

import (
	"GameAdmin/internal/models/msg"
	"GameAdmin/internal/schema"
	"GameAdmin/pkg/store"
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
)

func NewMsgSrv(msgRepo *msg.MsgRepo) *MsgSrv {
	return &MsgSrv{MsgRepo: msgRepo, s: store.New(map[string]*schema.NoticeQueryResult{}), g: singleflight.Group{}}
}

type MsgSrv struct {
	MsgRepo *msg.MsgRepo
	s       *store.Store[*schema.NoticeQueryResult]
	g       singleflight.Group
}

// GetNotices 获取公告信息
func (a *MsgSrv) GetNotices(ctx context.Context, pageNum int, pageSize int) (*schema.NoticeQueryResult, error) {
	// 尝试从缓存中取
	key := fmt.Sprintf("%s:%d-%d", NoticesKey, pageNum, pageSize)

	result := a.s.Get(key)
	if result != nil {
		return result, nil
	}

	// 防止缓存击穿
	val, err, _ := a.g.Do(key, func() (interface{}, error) {
		// 从数据库中取
		result, err := a.MsgRepo.GetNotices(ctx, pageNum, pageSize)
		if err != nil {
			return nil, err
		}

		// 再放入缓存
		a.s.Set(key, result)
		return result, nil
	})

	if err != nil {
		return nil, err
	}

	return val.(*schema.NoticeQueryResult), nil
}

// AddNotice 新建公告
func (a *MsgSrv) AddNotice(ctx context.Context, params *msg.SchemaNotice) error {
	err := a.MsgRepo.AddNotice(ctx, params)
	a.s.LikeDeletes(NoticesKey)
	return err
}

// UpdateNotice 更新公告
func (a *MsgSrv) UpdateNotice(ctx context.Context, params *msg.SchemaNotice) error {
	err := a.MsgRepo.UpdateNotice(ctx, params)
	a.s.LikeDeletes(NoticesKey)
	return err
}

// DeleteNotice 删除公告
func (a *MsgSrv) DeleteNotice(ctx context.Context, id uint64) error {
	err := a.MsgRepo.DeleteNotice(ctx, id)
	a.s.LikeDeletes(NoticesKey)
	return err
}
