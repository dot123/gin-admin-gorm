package service

import (
	"GameAdmin/internal/config"
	"GameAdmin/internal/models/msg"
	"GameAdmin/internal/schema"
	"GameAdmin/pkg/redisHelper"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/sync/singleflight"
	"time"
)

func NewMsgSrv(msgRepo *msg.MsgRepo) *MsgSrv {
	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       0,
	})
	return &MsgSrv{MsgRepo: msgRepo, Ring: ring, g: singleflight.Group{}}
}

type MsgSrv struct {
	MsgRepo *msg.MsgRepo
	Ring    *redis.Ring
	g       singleflight.Group
}

// GetNotices 获取公告信息
func (a *MsgSrv) GetNotices(ctx context.Context, pageNum int, pageSize int) (*schema.NoticeQueryResult, error) {
	// 尝试从缓存中取
	var data schema.NoticeQueryResult
	key := fmt.Sprintf("%s:%d-%d", NoticesKey, pageNum, pageSize)

	err := redisHelper.Get(a.Ring, key, &data)
	if err != nil {
		if redis.Nil != err {
			return nil, err
		}
	} else {
		// 从缓存中取到了
		return &data, nil
	}

	// 防止缓存击穿
	val, err, _ := a.g.Do(key, func() (interface{}, error) {
		// 从数据库中取
		result, err := a.MsgRepo.GetNotices(ctx, pageNum, pageSize)
		if err != nil {
			return nil, err
		}

		// 再放入缓存
		if err = redisHelper.Set(a.Ring, key, result, 60*time.Second); err != nil {
			return nil, err
		}
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
	redisHelper.LikeDeletes(a.Ring, NoticesKey)
	return err
}

// UpdateNotice 更新公告
func (a *MsgSrv) UpdateNotice(ctx context.Context, params *msg.SchemaNotice) error {
	err := a.MsgRepo.UpdateNotice(ctx, params)
	redisHelper.LikeDeletes(a.Ring, NoticesKey)
	return err
}

// DeleteNotice 删除公告
func (a *MsgSrv) DeleteNotice(ctx context.Context, id uint64) error {
	err := a.MsgRepo.DeleteNotice(ctx, id)
	redisHelper.LikeDeletes(a.Ring, NoticesKey)
	return err
}
