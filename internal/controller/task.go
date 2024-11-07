package controller

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/api/task"
	"go-project-manage-server/internal/constants"
	"go-project-manage-server/internal/service"
	"go-project-manage-server/utils"
)

type TaskController struct{}

func TaskNew() taskApi.ITask {
	return &TaskController{}
}

// SelectDateTask 根据日期获取任务
func (t *TaskController) SelectDateTask(c *gin.Context) {
	uid, ok := c.Get("uid")
	if !ok {
		utils.Result(c, constants.StatusError, "获取失败")
		return
	}

	v := &taskApi.SelectDateTaskResponse{}
	if err := c.ShouldBindJSON(&v); err != nil {
		utils.Result(c, constants.StatusError, err.Error())
		return
	}

	if data, count, err := service.Task().SelectDateTask(uid.(string), v); err != nil {
		utils.Result(c, constants.StatusError, err.Error())
	} else {
		utils.ResultDataList(c, constants.StatusOK, "获取成功", data, count)
	}
}
