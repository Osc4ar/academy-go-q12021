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
	GetAll(t []*model.Task) ([]*model.Task, error)
	GetOne(t *model.Task, id uint) (*model.Task, error)
}

func NewTaskInteractor(r repository.TaskRepository, p presenter.TaskPresenter) TaskInteractor {
	return &taskInteractor{r, p}
}

func (ti *taskInteractor) GetAll(t []*model.Task) ([]*model.Task, error) {
	t, err := ti.TaskRepository.FindAll(t)
	if err != nil {
		return nil, err
	}

	return ti.TaskPresenter.ResponseTasks(t), nil
}

func (ti *taskInteractor) GetOne(t *model.Task, id uint) (*model.Task, error) {
	t, err := ti.TaskRepository.Find(t, id)
	if err != nil {
		return nil, err
	}

	return ti.TaskPresenter.ResponseTask(t), nil
}
