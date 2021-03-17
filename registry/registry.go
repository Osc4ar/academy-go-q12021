package registry

import (
	"taskmanager/infrastructure/datastore"
	"taskmanager/interface/controllers"
)

type registry struct {
	database datastore.DB
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(database datastore.DB) Registry {
	return &registry{database}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewTaskController()
}
