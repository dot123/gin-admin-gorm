package schema

import "GameAdmin/pkg/types"

type Notice struct {
	ID        uint64     `json:"id" form:"id"`
	CreatedAt types.Time `json:"created_at" form:"created_at"`
	StartTime types.Time `json:"start_time" form:"start_time" validate:"required"`
	EndTime   types.Time `json:"end_time" form:"end_time" validate:"required"`
	Title     string     `json:"title" form:"title" validate:"required"`
	Content   string     `json:"content" form:"content" validate:"required"`
	Operator  string     `json:"operator" form:"operator" validate:"required"`
}

type Notices []*Notice

type NoticeQueryResult struct {
	List  Notices `json:"list"`
	Total int64   `json:"total"`
}
