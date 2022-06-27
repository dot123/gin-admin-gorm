package service

import (
	"GameAdmin/internal/models/role"
	"context"
	"github.com/google/wire"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleSrv), "*"))

type RoleSrv struct {
	RoleRepo *role.RoleRepo
}

// GetUserRoles 获取用户身份信息
func (a *RoleSrv) GetUserRoles(ctx context.Context, userName string) (*[]*role.Role, error) {
	return a.RoleRepo.GetUserRoles(ctx, userName)
}
