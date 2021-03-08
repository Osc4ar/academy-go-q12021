package repository

import (
	"taskmanager/domain/model"
	"taskmanager/infrastructure/datastore"
)

type taskRepository struct {
	database datastore.DB
}

// TaskRepository contains the Methods used with the Tasks data
type TaskRepository interface {
	FindAll(t []*model.Task) ([]*model.Task, error)
	Find(t *model.Task, id uint) (*model.Task, error)
}

// NewTaskRepository creates a new TaskRepository with the given Database
func NewTaskRepository(database datastore.DB) TaskRepository {
	return &taskRepository{database: database}
}

func (tr *taskRepository) FindAll(t []*model.Task) ([]*model.Task, error) {
	return tr.database.FindAll(), nil
}

func (tr *taskRepository) Find(t *model.Task, id uint) (*model.Task, error) {
	return tr.database.FindByID(id)
}
