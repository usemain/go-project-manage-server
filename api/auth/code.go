package authApi

type CodeRequest struct {
	Email string `form:"email" binding:"required"`
}
