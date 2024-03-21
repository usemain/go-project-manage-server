package model

// TASK 任务表
type TASK struct {
	Tid        int64  `json:"tid" gorm:"primaryKey;comment:任务ID"`
	Uid        string `json:"uid"  gorm:"comment:任务所属用户ID"`
	Gid        int64  `json:"gid" gorm:"comment:小组ID"`
	CreateTime string `json:"createTime" gorm:"comment:创建时间"`
	Title      string `json:"title" gorm:"comment:任务标题"`
	Content    string `json:"content" gorm:"comment:任务内容"`
	Type       uint8  `json:"type" gorm:"comment:任务类型"`
	Fixed      bool   `json:"fixed" gorm:"comment:是否固定顶部"`
	Status     bool   `json:"status" gorm:"comment:任务状态"`
	Users      string `json:"users" gorm:"comment:任务参与用户"`
}

func (table *TASK) TableName() string {
	return "task"
}
