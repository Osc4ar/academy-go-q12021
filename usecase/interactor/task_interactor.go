package interactor

import (
	"taskmanager/domain/model"
	"taskmanager/usecase/presenter"
	"taskmanager/usecase/repository"
)

type taskInteractor struct {
	TaskRepository repository.TaskRepository
	TaskPresenter  presenter.TaskPresenter
}

type TaskInteractor interface {
	Get(t []*model.Task) ([]*model.Task, error)
}

func NewTaskInteractor(r repository.TaskRepository, p presenter.TaskPresenter) TaskInteractor {
	return &taskInteractor{r, p}
}

func (ti *taskInteractor) Get(t []*model.Task) ([]*model.Task, error) {
	t, err := ti.TaskRepository.FindAll(t)
	if err != nil {
		return nil, err
	}

	return ti.TaskPresenter.ResponseTasks(t), nil
}
