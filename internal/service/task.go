package service

import (
	"go-project-manage-server/api/task"
)

type ITask interface {
	SelectDateTask(uid string, v *taskApi.SelectDateTaskResponse) (data []taskApi.SelectDateTaskRequest, total int64, err error)
}

var serviceTask ITask

func RegisterServiceTask(s ITask) {
	serviceTask = s
}

func Task() ITask {
	if serviceTask == nil {
		panic("implement not found for interface ITask, forgot register?")
	}
	return serviceTask
}
