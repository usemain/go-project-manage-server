package client

import (
	"gin-choes-server/internal/controller/client"
	"gin-choes-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(api *gin.RouterGroup) {
	user := api.Group("user").Use(middleware.AuthToken())
	{
		user.GET("userinfo", client.UserNewV1().Userinfo)
	}
}
