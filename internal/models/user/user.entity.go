package user

import (
	"GameAdmin/internal/models/role"
	"GameAdmin/internal/models/util"
	"GameAdmin/internal/schema"
	"GameAdmin/internal/schema/enum"
	"GameAdmin/pkg/structure"
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	Username   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	Avatar     string    `gorm:"column:avatar;default:https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG"`
	UserType   int       `gorm:"column:user_type;default:0;NOT NULL"`
	State      int       `gorm:"column:state;default:1;NOT NULL"`
	CreatedBy  string    `gorm:"column:created_by"`
	ModifiedBy string    `gorm:"column:modified_by"`
}

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(User))
}

// UserRole 用户身份结构体
type UserRole struct {
	UserName  string
	UserID    uint64
	UserRoles []*role.Role
}

type Users []*User

type SchemaUser schema.User

func (a SchemaUser) ToUser() *User {
	item := new(User)
	structure.Copy(a, item)
	return item
}

func (m User) ToSchemaUser() *schema.User {
	item := new(schema.User)
	structure.Copy(m, item)
	item.UserType = enum.GetUserType(m.UserType)
	item.State = enum.GetStatus(m.State)
	item.CreatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	item.UpdatedAt = m.UpdatedAt.Format("2006-01-02 15:04:05")
	return item
}

func (a Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}
