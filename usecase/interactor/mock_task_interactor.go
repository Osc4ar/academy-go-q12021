package interactor

import (
	"fmt"
	"taskmanager/domain/model"
)

type mockTaskInteractor struct {
	tasks []*model.Task
}

func NewMockTaskInteractor(tasks []*model.Task) TaskInteractor {
	return &mockTaskInteractor{tasks: tasks}
}

func (mti *mockTaskInteractor) GetAll(t []*model.Task) ([]*model.Task, error) {
	return mti.tasks, nil
}

func (mti *mockTaskInteractor) GetOne(t *model.Task, id uint) (*model.Task, error) {
	for _, task := range mti.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, fmt.Errorf("Could not find ID: %v", id)
}
