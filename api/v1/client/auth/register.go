package auth

type RegisterRequest struct {
	Email string `form:"email" binding:"required,email"`
	Pwd   string `form:"password" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
