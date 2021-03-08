package presenter

import (
	"taskmanager/domain/model"
)

type taskPresenter struct {
}

type TaskPresenter interface {
	ResponseTasks(t []*model.Task) []*model.Task
	ResponseTask(t *model.Task) *model.Task
}

func NewTaskPresenter() TaskPresenter {
	return &taskPresenter{}
}

func (tp *taskPresenter) ResponseTasks(t []*model.Task) []*model.Task {
	for _, task := range t {
		tp.ResponseTask(task)
	}

	return t
}

func (tp *taskPresenter) ResponseTask(t *model.Task) *model.Task {
	if t.Completed {
		t.Content = "ToDo: " + t.Content
	} else {
		t.Content = "Done: " + t.Content
	}

	return t
}
