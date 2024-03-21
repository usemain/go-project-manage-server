package client

import (
	"gin-choes-server/api/v1/client/task"
	"gin-choes-server/internal/global"
	"gin-choes-server/internal/model"
	"gin-choes-server/internal/service/client"
)

type STask struct{}

func init() {
	client.RegisterServiceTask(&STask{})
}

// SelectDateTask 根据日期获取任务
func (s *STask) SelectDateTask(uid string, v *task.SelectDateTaskResponse) (data []task.SelectDateTaskRequest, count int64, err error) {
	var tasks []model.TASK

	global.GVA_DB.Model(&model.TASK{}).Where("uid = ? and create_time", uid, v.Date).Count(&count).Offset((v.Page - 1) * v.PageSize).Limit(v.PageSize).Find(&tasks)

	for _, taskItem := range tasks {
		item := task.SelectDateTaskRequest{
			Tid:        taskItem.Tid,
			CreateTime: taskItem.CreateTime,
			Title:      taskItem.Title,
			Content:    taskItem.Content,
			Type:       taskItem.Type,
			Status:     taskItem.Status,
		}
		data = append(data, item)
	}

	return
}
