package repository

import (
	"taskmanager/domain/model"
	"taskmanager/infrastructure/client"
	"taskmanager/infrastructure/datastore"
)

type taskRepository struct {
	database datastore.DB
	client   client.ApiClient
}

// TaskRepository contains the Methods used with the Tasks data
type TaskRepository interface {
	FindAll(t []*model.Task) ([]*model.Task, error)
	Find(t *model.Task, id uint) (*model.Task, error)
}

// NewTaskRepository creates a new TaskRepository with the given Database
func NewTaskRepository(database datastore.DB, client client.ApiClient) TaskRepository {
	return &taskRepository{database: database, client: client}
}

func (tr *taskRepository) FindAll(t []*model.Task) ([]*model.Task, error) {
	t, err := tr.client.RequestAllTasks(t)
	if err != nil {
		return nil, err
	}

	go tr.database.SaveRecords(t)

	return t, nil
}

func (tr *taskRepository) Find(t *model.Task, id uint) (*model.Task, error) {
	return tr.database.FindByID(id)
}
