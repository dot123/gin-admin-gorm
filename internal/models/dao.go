package models

import (
	"GameAdmin/internal/models/file"
	"GameAdmin/internal/models/role"
	"GameAdmin/internal/models/user"
	"GameAdmin/internal/models/util"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RepoSet = wire.NewSet(
	util.TransSet,
	role.RoleSet,
	user.UserSet,
	file.FileSet,
)

type (
	TransRepo = util.Trans
	RoleRepo  = role.RoleRepo
	UserRepo  = user.UserRepo
	FileRepo  = file.FileRepo
)

func AutoMigrate(db *gorm.DB) error {
	existUser := db.Migrator().HasTable(new(user.User))
	existRole := db.Migrator().HasTable(new(role.Role))

	err := db.AutoMigrate(
		new(role.Role),
		new(user.User),
		new(file.File),
	)

	if !existUser {
		db.Create(&user.User{ID: 1, Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Avatar: "", UserType: 1, State: 1})
		db.Create(&user.User{ID: 2, Username: "test", Password: "e10adc3949ba59abbe56e057f20f883e", Avatar: "", UserType: 2, State: 1})
	}

	if !existRole {
		db.Create(&role.Role{ID: 1, UserID: 1, UserName: "admin", Value: "admin"})
		db.Create(&role.Role{ID: 2, UserID: 1, UserName: "admin", Value: "test"})
		db.Create(&role.Role{ID: 3, UserID: 2, UserName: "test", Value: "test"})
	}

	return err
}
