package model

// USER 用户表
type USER struct {
	UID    string `json:"uid" gorm:"primaryKey;comment:用户ID"`
	Email  string `json:"email" gorm:"comment:账号"`
	Pwd    string `json:"pwd" gorm:"comment:密码"`
	Head   string `json:"head" gorm:"comment:头像"`
	Name   string `json:"name" gorm:"comment:昵称"`
	Sex    uint8  `json:"sex" gorm:"comment:性别"`
	Level  uint8  `json:"level" gorm:"comment:等级"`
	Status uint8  `json:"status" gorm:"comment:账户状态"`
}

func (table *USER) TableName() string {
	return "user"
}
