package task

import "time"

type SelectDateTaskResponse struct {
	Date       string `json:"date" binding:"required"`
	Page       int    `json:"page" binging:"required,numeric"`
	PageSize   int    `json:"pageSize" binding:"required,numeric"`
	SearchName string `json:"searchName"`
}

type SelectDateTaskRequest struct {
	Tid        int64     `json:"tid"`
	CreateTime time.Time `json:"createTime"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Type       uint8     `json:"type"`
	Status     bool      `json:"status"`
}
