package client

import (
	"gin-choes-server/api/v1/client/task"
	"gin-choes-server/internal/service/client"
)

type STask struct{}

func init() {
	client.RegisterServiceTask(&STask{})
}

// SelectDateTask 根据日期获取任务
func (t *STask) SelectDateTask(v *task.SelectDateTaskResponse) (data []task.SelectDateTaskRequest, err error) {

	return
}
