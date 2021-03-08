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
}

// NewTaskController creates a new Controller with the given TaskInteractor
func NewTaskController(ti interactor.TaskInteractor) TaskController {
	return &taskController{ti}
}

func (tc *taskController) GetTasks(w http.ResponseWriter) error {
	var t []*model.Task

	t, err := tc.taskInteractor.Get(t)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(t)
}
