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
	tasksToPresent := []*model.Task{}
	for _, task := range t {
		tasksToPresent = append(tasksToPresent, tp.ResponseTask(task))
	}

	return tasksToPresent
}

func (tp *taskPresenter) ResponseTask(t *model.Task) *model.Task {
	newTask := *t
	if t.Completed {
		newTask.Content = "Done: " + t.Content
	} else {
		newTask.Content = "ToDo: " + t.Content
	}

	return &newTask
}
