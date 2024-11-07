package handler

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/internal/controller"
	"go-project-manage-server/internal/middleware"
)

func User(api *gin.RouterGroup) {
	user := api.Group("user").Use(middleware.AuthToken())
	userNew := controller.UserNew()
	{
		user.GET("userinfo", userNew.Userinfo)
	}
}
