package handler

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/internal/controller"
)

func Auth(api *gin.RouterGroup) {
	auth := api.Group("auth")
	authNew := controller.AuthNew()
	{
		auth.GET("code", authNew.Code)
		auth.POST("login", authNew.Login)
		auth.POST("register", authNew.Register)
	}
}
