package registry

import (
	"taskmanager/infrastructure/client"
	"taskmanager/infrastructure/datastore"
	"taskmanager/interface/controllers"
)

type registry struct {
	database datastore.DB
	client   client.ApiClient
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(database datastore.DB, client client.ApiClient) Registry {
	return &registry{database, client}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewTaskController()
}
