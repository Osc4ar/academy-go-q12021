package controllers

import (
	"encoding/json"
	"fmt"
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

	fmt.Printf("Tasks: %v\n", t)

	return json.NewEncoder(w).Encode(t)
}

func (tc *taskController) GetTask(id uint, w http.ResponseWriter) error {
	var t *model.Task

	fmt.Printf("ID: %v\n", id)

	t, err := tc.taskInteractor.GetOne(t, id)
	if err != nil {
		return err
	}

	fmt.Printf("Task: %v\n", t)

	return json.NewEncoder(w).Encode(t)
}
