package auth

type LoginRequest struct {
	Email string `form:"email" binding:"required,email"`
	Pwd   string `form:"pwd" binding:"required"`
}

type LoginResponse struct {
	Gid               int64  `json:"gid"`
	IsGroupLeader     bool   `json:"isGroupLeader"`
	IsGroupCreateTask bool   `json:"isGroupCreateTask"`
	LastTime          string `json:"lastTime"`
	Email             string `json:"email"`
	Head              string `json:"head"`
	Name              string `json:"name"`
	Sex               uint8  `json:"sex"`
	Level             uint8  `json:"level"`
	Token             string `json:"token"`
}
