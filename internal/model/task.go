package model

import "time"

// TASK 任务表
type TASK struct {
	TID        int64     `json:"tid" gorm:"primaryKey;comment:任务ID"`
	UID        string    `json:"uid"  gorm:"comment:任务所属用户ID"`
	CreateTime time.Time `json:"createTime" gorm:"comment:创建时间"`
	Title      string    `json:"title" gorm:"comment:任务标题"`
	Content    string    `json:"content" gorm:"comment:任务内容"`
	Type       uint8     `json:"type" gorm:"comment:任务类型"`
	Status     bool      `json:"status" gorm:"comment:任务状态"`
}
