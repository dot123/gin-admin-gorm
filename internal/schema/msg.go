package schema

import "GameAdmin/pkg/types"

type Notice struct {
	ID        uint64     `json:"id" form:"id"`
	CreatedAt types.Time `json:"created_at" form:"created_at"`
	StartTime types.Time `json:"start_time" form:"start_time" binding:"required"`
	EndTime   types.Time `json:"end_time" form:"end_time" binding:"required"`
	Title     string     `json:"title" form:"title" binding:"required"`
	Content   string     `json:"content" form:"content" binding:"required"`
	Operator  string     `json:"operator" form:"operator" binding:"required"`
}

type Notices []*Notice

type NoticeQueryResult struct {
	List  Notices `json:"list"`
	Total int64   `json:"total"`
}
