package model

// USER 用户表
type USER struct {
	Uid           string `json:"uid" gorm:"primaryKey;comment:用户ID"`
	Gid           int64  `json:"gid" gorm:"comment:小组"`
	CreateTime    string `json:"createTime" gorm:"comment:创建时间"`
	LastTime      string `json:"lastTime" gorm:"comment:最后登录时间"`
	Email         string `json:"email" gorm:"comment:邮箱账号"`
	Pwd           string `json:"pwd" gorm:"comment:密码"`
	Head          string `json:"head" gorm:"comment:头像"`
	Name          string `json:"name" gorm:"comment:昵称"`
	Sex           uint8  `json:"sex" gorm:"comment:性别"`
	VipLevel      uint8  `json:"vipLevel" gorm:"comment:等级"`
	VipExpireTime string `json:"vipExpireTime" gorm:"comment:VIP到期时间"`
	Status        bool   `json:"status" gorm:"comment:用户状态"`
}

func (table *USER) TableName() string {
	return "user"
}
