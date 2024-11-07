package authApi

import (
	"github.com/gin-gonic/gin"
)

type IAuth interface {
	Code(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
}
