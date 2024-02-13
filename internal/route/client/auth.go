package client

import (
	"gin-choes-server/internal/controller/client"
	"github.com/gin-gonic/gin"
)

func AuthRoute(api *gin.RouterGroup) {
	sign := api.Group("auth")
	{
		sign.GET("code", client.AuthNewV1().Code)
		sign.POST("login", client.AuthNewV1().Login)
		sign.POST("register", client.AuthNewV1().Register)
	}
}
