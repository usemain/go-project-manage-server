package client

import (
	"gin-choes-server/internal/controller/client"
	"gin-choes-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRoute(api *gin.RouterGroup) {
	task := api.Group("task").Use(middleware.AuthToken())
	{
		task.GET("selectDateTask", client.TaskNewV1().SelectDateTask)
	}
}
