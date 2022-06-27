//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	v1 "GameAdmin/api/v1"
	jwt "GameAdmin/internal/middleware/jwt"
	"GameAdmin/internal/models"
	"GameAdmin/internal/routers"
	"GameAdmin/internal/service"
	"GameAdmin/pkg/fileStore"
	"github.com/google/wire"
)

func BuildInjector(*fileStore.Local) (*Injector, func(), error) {
	wire.Build(InitGormDB, InitGinEngine, models.RepoSet, service.ProviderSet, jwt.JWTSet, v1.ProviderSet, routers.RouterSet, InjectorSet)
	return new(Injector), nil, nil
}
