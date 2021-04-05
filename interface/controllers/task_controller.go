package controllers

import (
	"encoding/json"
	"net/http"

	"taskmanager/domain/model"
	"taskmanager/usecase/interactor"
)

type taskController struct {
	taskInteractor interactor.TaskInteractor
}

// TaskController has the structure for a HTTP controller
type TaskController interface {
	GetTasks(w http.ResponseWriter) error
	GetTask(id uint, w http.ResponseWriter) error
	GetTaskConcurrently(isEven bool, items, itemsPerWorker int, w http.ResponseWriter) error
}

// NewTaskController creates a new Controller with the given TaskInteractor
func NewTaskController(ti interactor.TaskInteractor) TaskController {
	return &taskController{ti}
}

func (tc *taskController) GetTasks(w http.ResponseWriter) error {
	var t []*model.Task

	t, err := tc.taskInteractor.GetAll(t)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(t)
}

func (tc *taskController) GetTask(id uint, w http.ResponseWriter) error {
	var t *model.Task

	t, err := tc.taskInteractor.GetOne(t, id)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(t)
}

func (tc *taskController) GetTaskConcurrently(isEven bool, items, itemsPerWorker int, w http.ResponseWriter) error {
	var t []*model.Task

	t, err := tc.taskInteractor.GetConcurrently(t, isEven, items, itemsPerWorker)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(t)
}
