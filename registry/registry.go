package registry

import (
	"encoding/csv"
	"taskmanager/interface/controllers"
)

type registry struct {
	reader *csv.Reader
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(reader *csv.Reader) Registry {
	return &registry{reader}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewTaskController()
}
