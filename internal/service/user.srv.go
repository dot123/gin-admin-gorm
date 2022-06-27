package service

import (
	"GameAdmin/internal/errors"
	"GameAdmin/internal/models"
	"GameAdmin/internal/models/role"
	"GameAdmin/internal/models/user"
	"GameAdmin/internal/schema"
	"GameAdmin/pkg/logger"
	"context"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSrv), "*"))

type UserSrv struct {
	TransRepo *models.TransRepo
	UserRepo  *user.UserRepo
	RoleRepo  *role.RoleRepo
}

// CheckUser 身份验证
func (a *UserSrv) CheckUser(ctx context.Context, username string, password string) (uint64, error) {
	return a.UserRepo.CheckUser(ctx, username, password)
}

// GetUserAvatar 获取用户头像
func (a *UserSrv) GetUserAvatar(ctx context.Context, username string) (string, error) {
	return a.UserRepo.GetUserAvatar(ctx, username)
}

// GetRoles 获取用户角色
func (a *UserSrv) GetRoles(ctx context.Context, username string) (*[]string, error) {
	userId, err := a.UserRepo.GetUserId(ctx, username)
	if err != nil {
		return nil, err
	}
	return a.RoleRepo.GetRoles(ctx, userId)
}

// GetUsers 获取用户信息
func (a *UserSrv) GetUsers(ctx context.Context, pageNum, pageSize int, name string) (*schema.UserQueryResult, error) {
	return a.UserRepo.GetUsers(ctx, pageNum, pageSize, name)
}

// AddUser 新建用户，同时新建用户角色
func (a *UserSrv) AddUser(ctx context.Context, params *schema.UserDataParam, createdBy string) error {
	user := user.User{}
	user.Username = params.Username
	user.Password = params.Password
	user.UserType = params.UserType
	user.Avatar = params.Avatar

	user.CreatedBy = createdBy
	user.State = 1
	if user.Avatar == "" {
		user.Avatar = "https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG"
	}

	// 此处不能使用事务同时创建用户和角色，因为Role表中需要UserId，而UserId需要插入用户数据后才生成，所以不能用事务，否则会报错
	// 用业务逻辑实现事务效果
	if err := a.UserRepo.AddUser(ctx, &user); err != nil {
		return err
	}
	// 当成功插入User数据后，user为指针地址，可以获取到ID的值。省去了查数据库拿ID的值步骤
	var role role.Role
	role.UserID = user.ID
	role.UserName = user.Username
	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	if err := a.RoleRepo.AddRole(ctx, &role); err != nil {
		// 插入role失败后，删除新插入的用户信息，达到事务处理效果
		return a.UserRepo.DeleteUser(ctx, user.ID)
	}
	return nil
}

// ExistUserByName 判断用户名是否已存在
func (a *UserSrv) ExistUserByName(ctx context.Context, username string) bool {
	return a.UserRepo.ExistUserByName(ctx, username)
}

// UpdateUser 更新用户
func (a *UserSrv) UpdateUser(ctx context.Context, params *schema.UserDataParam, modifiedBy string) error {
	user, err := a.UserRepo.GetUserById(ctx, params.ID)
	if err != nil {
		return err
	}

	user.Password = params.Password
	user.ModifiedBy = modifiedBy
	user.UserType = params.UserType
	user.Avatar = params.Avatar
	role, err := a.RoleRepo.GetRole(ctx, user.ID)
	if err != nil {
		return err
	}

	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}

	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		if err := a.RoleRepo.UpdateRole(ctx, role); err != nil {
			return err
		}
		return a.UserRepo.UpdateUser(ctx, user)
	})
	return err
}

// DeleteUser 删除用户
func (a *UserSrv) DeleteUser(ctx context.Context, id uint64) error {
	user, err := a.UserRepo.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	if user.Username == "admin" {
		err = errors.New("删除用户失败:不能删除admin账号")
		logger.WithContext(ctx).Error(err)
		return err
	}

	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		if err := a.RoleRepo.DeleteRole(ctx, id); err != nil {
			return err
		}
		return a.UserRepo.DeleteUser(ctx, id)
	})

	return err
}
