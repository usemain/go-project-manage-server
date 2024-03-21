package model

// GROUP 小组表
type GROUP struct {
	Gid        int64  `json:"gid" gorm:"primaryKey;comment:小组ID"`
	Uid        string `json:"uid"  gorm:"comment:小组所属用户ID"`
	CreateTime string `json:"createTime" gorm:"comment:创建时间"`
	GroupName  string `json:"groupName" gorm:"comment:小组名称"`
	GroupDesc  string `json:"groupDesc" gorm:"comment:小组描述"`
	GroupType  uint8  `json:"groupType" gorm:"comment:小组类型"`
	GroupUsers string `json:"groupUsers" gorm:"comment:小组成员"`
	Status     bool   `json:"status" gorm:"comment:小组状态"`
}

func (table *GROUP) TableName() string {
	return "group"
}
