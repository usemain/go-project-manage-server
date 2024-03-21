package auth

type LoginRequest struct {
	Email string `form:"email" binding:"required,email"`
	Pwd   string `form:"pwd" binding:"required"`
}

type LoginResponse struct {
	Gid               int64  `json:"gid,omitempty"`
	IsGroupLeader     bool   `json:"isGroupLeader,omitempty"`
	IsGroupCreateTask bool   `json:"isGroupCreateTask,omitempty"`
	LastTime          string `json:"lastTime"`
	Email             string `json:"email"`
	Head              string `json:"head"`
	Name              string `json:"name"`
	Sex               uint8  `json:"sex"`
	VipLevel          uint8  `json:"vipLevel"`
	VipExpireTime     string `json:"VipExpireTime"`
	Token             string `json:"token"`
}
