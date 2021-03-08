package presenter

import (
	"taskmanager/domain/model"
)

type taskPresenter struct {
}

type TaskPresenter interface {
	ResponseTasks(t []*model.Task) []*model.Task
}

func NewTaskPresenter() TaskPresenter {
	return &taskPresenter{}
}

func (tp *taskPresenter) ResponseTasks(t []*model.Task) []*model.Task {
	for _, task := range t {
		if task.Completed {
			task.Content = "ToDo: " + task.Content
		} else {
			task.Content = "Done: " + task.Content
		}
	}

	return t
}
