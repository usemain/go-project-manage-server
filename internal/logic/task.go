package logic

import (
	"go-project-manage-server/api/task"
	"go-project-manage-server/internal/global"
	"go-project-manage-server/internal/model"
	"go-project-manage-server/internal/service"
)

type ITask struct{}

func init() {
	service.RegisterServiceTask(&ITask{})
}

// SelectDateTask 根据日期获取任务
func (s *ITask) SelectDateTask(uid string, v *taskApi.SelectDateTaskResponse) (data []taskApi.SelectDateTaskRequest, count int64, err error) {
	var tasks []model.TASK

	global.GVA_DB.Model(&model.TASK{}).Where("uid = ? and create_time", uid, v.Date).Count(&count).Offset((v.Page - 1) * v.PageSize).Limit(v.PageSize).Find(&tasks)

	for _, taskItem := range tasks {
		item := taskApi.SelectDateTaskRequest{
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
