package client

import (
	"gin-choes-server/internal/controller/client"
	"github.com/gin-gonic/gin"
)

func AuthRoute(api *gin.RouterGroup) {
	auth := api.Group("auth")
	{
		auth.GET("code", client.AuthNewV1().Code)
		auth.POST("login", client.AuthNewV1().Login)
		auth.POST("register", client.AuthNewV1().Register)
	}
}
