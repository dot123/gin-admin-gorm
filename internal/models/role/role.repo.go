package role

import (
	"GameAdmin/pkg/logger"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleRepo), "*"))

type RoleRepo struct {
	DB *gorm.DB
}

// GetUserRoles 获取用户身份信息
func (a *RoleRepo) GetUserRoles(ctx context.Context, userName string) (*[]*Role, error) {
	var roles []*Role

	if err := GetRoleDB(ctx, a.DB).Where("`user_name` = ?", userName).Find(&roles).Error; err != nil {
		logger.WithContext(ctx).Errorf("获取用户身份信息错误: %v", err)
		return &roles, err
	}
	return &roles, nil
}

// GetRoles 获取用户角色
func (a *RoleRepo) GetRoles(ctx context.Context, userID uint64) (*[]string, error) {
	var arrRole []string
	var roles []*Role

	if err := GetRoleDB(ctx, a.DB).Where("`user_id` = ?", userID).Select("value").Find(&roles).Error; err != nil {
		logger.WithContext(ctx).Errorf("获取用户角色失败: %v", err)
		return &arrRole, err
	}
	for _, role := range roles {
		arrRole = append(arrRole, role.Value)
	}
	return &arrRole, nil
}

// AddRole 添加用户角色
func (a *RoleRepo) AddRole(ctx context.Context, role *Role) error {
	if err := GetRoleDB(ctx, a.DB).Create(role).Error; err != nil {
		logger.WithContext(ctx).Errorf("添加用户角色失败: %v", err)
		return err
	}
	return nil
}

// GetRole 获取角色
func (a *RoleRepo) GetRole(ctx context.Context, userID uint64) (*Role, error) {
	var role Role

	if err := GetRoleDB(ctx, a.DB).Where("`user_id` = ?", userID).Take(&role).Error; err != nil {
		logger.WithContext(ctx).Errorf("获取角色失败: %v", err)
		return nil, err
	}
	return &role, nil
}

// DeleteRole 删除用户角色
func (a *RoleRepo) DeleteRole(ctx context.Context, userID uint64) error {
	if err := GetRoleDB(ctx, a.DB).Where("`user_id` = ?", userID).Delete(new(Role)).Error; err != nil {
		logger.WithContext(ctx).Errorf("删除用户角色失败: %v", err)
		return err
	}
	return nil
}

// UpdateRole 更新用户角色
func (a *RoleRepo) UpdateRole(ctx context.Context, role *Role) error {
	if err := GetRoleDB(ctx, a.DB).Where("`id` = ?", role.ID).Updates(role).Error; err != nil {
		logger.WithContext(ctx).Errorf("更新用户角色失败: %v", err)
		return err
	}
	return nil
}
