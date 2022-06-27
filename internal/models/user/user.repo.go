package user

import (
	"GameAdmin/internal/models/util"
	"GameAdmin/internal/schema"
	"GameAdmin/pkg/logger"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

type UserRepo struct {
	DB *gorm.DB
}

// CheckUser 身份验证
func (a *UserRepo) CheckUser(ctx context.Context, username string, password string) (uint64, error) {
	var user User

	if err := GetUserDB(ctx, a.DB).Where("`username` = ? AND `password` = ?", username, password).Take(&user).Error; err != nil {
		logger.WithContext(ctx).Errorf("用户名或密码错误: %v", err)
		return 0, err
	}

	return user.ID, nil
}

// GetUserAvatar 获取用户头像
func (a *UserRepo) GetUserAvatar(ctx context.Context, username string) (string, error) {
	var user User

	if err := GetUserDB(ctx, a.DB).Where("`username` = ?", username).Select("avatar").Take(&user).Error; err != nil {
		logger.WithContext(ctx).Errorf("获取用户头像失败: %v", err)
		return "", err
	}

	return user.Avatar, nil
}

// GetUserId 获取用户Id
func (a *UserRepo) GetUserId(ctx context.Context, username string) (uint64, error) {
	var user User

	if err := GetUserDB(ctx, a.DB).Where("`username` = ?", username).Take(&user).Error; err != nil {
		logger.WithContext(ctx).Errorf("获取用户ID失败: %v", err)
		return 0, err
	}

	return user.ID, nil
}

// GetUsers 获取用户信息
func (a *UserRepo) GetUsers(ctx context.Context, pageNum int, pageSize int, name string) (*schema.UserQueryResult, error) {
	var list Users
	var total int64

	db := GetUserDB(ctx, a.DB)
	if name != "" {
		db = db.Where("`username` LIKE ?", "%"+name+"%")
	}

	total, err := util.GetPages(db, &list, pageNum, pageSize)
	if err != nil {
		logger.WithContext(ctx).Errorf("获取用户信息失败: %v", err)
		return nil, err
	}
	var result schema.UserQueryResult
	result.List = list.ToSchemaUsers()
	result.Total = total
	return &result, nil
}

// AddUser 新建用户
func (a *UserRepo) AddUser(ctx context.Context, user *User) error {
	if err := GetUserDB(ctx, a.DB).Create(user).Error; err != nil {
		logger.WithContext(ctx).Errorf("新建用户失败: %v", err)
		return err
	}
	return nil
}

// ExistUserByName 判断用户名是否已存在
func (a *UserRepo) ExistUserByName(ctx context.Context, username string) bool {
	var user User

	err := GetUserDB(ctx, a.DB).Where("`username` = ?", username).Select("id").Take(&user).Error
	//记录不存在错误(RecordNotFound)，返回false
	if err == gorm.ErrRecordNotFound {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		logger.WithContext(ctx).Errorf("判断用户名是否已存在失败: %v", err)
		return false
	}
	return true
}

// UpdateUser 更新用户
func (a *UserRepo) UpdateUser(ctx context.Context, user *User) error {
	if err := GetUserDB(ctx, a.DB).Where("`id` = ?", user.ID).Updates(user).Error; err != nil {
		logger.WithContext(ctx).Errorf("更新用户失败: %v", err)
		return err
	}
	return nil
}

// DeleteUser 删除用户
func (a *UserRepo) DeleteUser(ctx context.Context, id uint64) error {
	if err := GetUserDB(ctx, a.DB).Where("`id` = ?", id).Delete(new(User)).Error; err != nil {
		logger.WithContext(ctx).Errorf("删除用户失败: %v", err)
		return err
	}
	return nil
}

// GetUserById 获取用户
func (a *UserRepo) GetUserById(ctx context.Context, id uint64) (*User, error) {
	var user User
	if err := GetUserDB(ctx, a.DB).Where("`id` = ?", id).Take(&user).Error; err != nil {
		logger.WithContext(ctx).Errorf("获取用户失败: %v", err)
		return &user, err
	}
	return &user, nil
}
