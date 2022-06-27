// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"GameAdmin/api/v1"
	"GameAdmin/internal/middleware/jwt"
	"GameAdmin/internal/models/file"
	"GameAdmin/internal/models/role"
	"GameAdmin/internal/models/user"
	"GameAdmin/internal/models/util"
	"GameAdmin/internal/routers"
	"GameAdmin/internal/service"
	"GameAdmin/pkg/fileStore"
)

import (
	_ "GameAdmin/docs"
)

// Injectors from wire.go:

func BuildInjector(local *fileStore.Local) (*Injector, func(), error) {
	db, cleanup, err := InitGormDB()
	if err != nil {
		return nil, nil, err
	}
	trans := &util.Trans{
		DB: db,
	}
	userRepo := &user.UserRepo{
		DB: db,
	}
	roleRepo := &role.RoleRepo{
		DB: db,
	}
	userSrv := &service.UserSrv{
		TransRepo: trans,
		UserRepo:  userRepo,
		RoleRepo:  roleRepo,
	}
	roleSrv := &service.RoleSrv{
		RoleRepo: roleRepo,
	}
	jwtJWT := &jwt.JWT{
		UserSrv: userSrv,
		RoleSrv: roleSrv,
	}
	userApi := &v1.UserApi{
		UserSrv: userSrv,
	}
	systemSrv := &service.SystemSrv{}
	systemApi := &v1.SystemApi{
		SystemSrv: systemSrv,
	}
	fileRepo := &file.FileRepo{
		DB: db,
	}
	fileSrv := &service.FileSrv{
		FileRepo: fileRepo,
		Local:    local,
	}
	fileApi := &v1.FileApi{
		FileSrv: fileSrv,
	}
	router := &routers.Router{
		MyJwt:     jwtJWT,
		UserApi:   userApi,
		SystemApi: systemApi,
		FileApi:   fileApi,
	}
	engine := InitGinEngine(router)
	injector := &Injector{
		Engine: engine,
	}
	return injector, func() {
		cleanup()
	}, nil
}
