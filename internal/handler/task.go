package handler

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/internal/controller"
	"go-project-manage-server/internal/middleware"
)

func Task(api *gin.RouterGroup) {
	task := api.Group("task").Use(middleware.AuthToken())
	taskNew := controller.TaskNew()
	{
		task.GET("selectDateTask", taskNew.SelectDateTask)
	}
}
