package client

import (
	"gin-choes-server/api/v1/client/task"
	"gin-choes-server/internal/consts"
	"gin-choes-server/internal/service/client"
	"gin-choes-server/utility"
	"github.com/gin-gonic/gin"
)

type TaskControllerV1 struct{}

func TaskNewV1() task.ITask {
	return &TaskControllerV1{}
}

// SelectDateTask 根据日期获取任务
func (t *TaskControllerV1) SelectDateTask(c *gin.Context) {
	uid, ok := c.Get("uid")
	if !ok {
		utility.Result(c, consts.StatusError, "获取失败")
		return
	}

	v := &task.SelectDateTaskResponse{}
	if err := c.ShouldBindJSON(&v); err != nil {
		utility.Result(c, consts.StatusError, err.Error())
		return
	}

	if data, count, err := client.Task().SelectDateTask(uid.(string), v); err != nil {
		utility.Result(c, consts.StatusError, err.Error())
	} else {
		utility.ResultDataList(c, consts.StatusOK, "获取成功", data, count)
	}
}
