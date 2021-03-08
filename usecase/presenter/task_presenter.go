package presenter

import (
	"taskmanager/domain/model"
)

// TaskPresenter is the interface needed to behave as an OutputPort
type TaskPresenter interface {
	ResponseTasks(t []*model.Task) []*model.Task
	ResponseTask(t *model.Task) *model.Task
}
