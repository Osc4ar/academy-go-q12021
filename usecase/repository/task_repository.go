package repository

import (
	"taskmanager/domain/model"
)

// TaskRepository is an interface which matches the behaviour of a Repository
type TaskRepository interface {
	FindAll(t []*model.Task) ([]*model.Task, error)
	Find(t *model.Task, id uint) (*model.Task, error)
}
