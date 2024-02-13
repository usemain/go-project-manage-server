package auth

type LoginRequest struct {
	Email string `form:"email" binding:"required,email"`
	Pwd   string `form:"pwd" binding:"required"`
}

type LoginResponse struct {
	Email  string `json:"email"`
	Head   string `json:"head"`
	Name   string `json:"name"`
	Sex    uint8  `json:"sex"`
	Level  uint8  `json:"level"`
	Status uint8  `json:"status"`
	Token  string `json:"token"`
}
