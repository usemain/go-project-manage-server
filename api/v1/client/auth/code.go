package auth

type CodeRequest struct {
	Email string `form:"email" binding:"required"`
}
