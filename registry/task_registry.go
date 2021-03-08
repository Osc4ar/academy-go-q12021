package registry

import (
	"taskmanager/interface/controllers"
	ipresenter "taskmanager/interface/presenter"
	irepository "taskmanager/interface/repository"
	"taskmanager/usecase/interactor"
	upresenter "taskmanager/usecase/presenter"
	urepository "taskmanager/usecase/repository"
)

func (r *registry) NewTaskController() controllers.TaskController {
	return controllers.NewTaskController(r.NewTaskInteractor())
}

func (r *registry) NewTaskInteractor() interactor.TaskInteractor {
	return interactor.NewTaskInteractor(r.NewTaskRepository(), r.NewTaskPresenter())
}

func (r *registry) NewTaskRepository() urepository.TaskRepository {
	return irepository.NewTaskRepository(r.database)
}

func (r *registry) NewTaskPresenter() upresenter.TaskPresenter {
	return ipresenter.NewTaskPresenter()
}
