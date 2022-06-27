package schema

// UserData 用户登录信息结构体
type UserData struct {
	Roles        []string
	Introduction string
	Avatar       string
	Name         string
}

// User 用户管理结构体
type User struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	UserType  string `json:"user_type"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserQueryResult struct {
	List  []*User `json:"list"`
	Total int64   `json:"total"`
}

type UserDataParam struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserType int    `json:"user_type" validate:"required"`
	Avatar   string `json:"avatar"`
}
