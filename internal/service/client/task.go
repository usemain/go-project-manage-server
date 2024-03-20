package client

import (
	"gin-choes-server/api/v1/client/task"
)

type ITask interface {
	SelectDateTask(v *task.SelectDateTaskResponse) (data []task.SelectDateTaskRequest, err error)
}

var serviceTask ITask

func RegisterServiceTask(s ITask) {
	serviceTask = s
}

func Task() ITask {
	if serviceTask == nil {
		panic("implement not found for interface STask, forgot register?")
	}
	return serviceTask
}
