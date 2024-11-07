package userApi

import (
	"github.com/gin-gonic/gin"
)

type IUser interface {
	Userinfo(c *gin.Context)
}
