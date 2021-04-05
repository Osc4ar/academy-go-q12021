package interactor

import (
	"fmt"
	"math"
	"sync"
	"taskmanager/domain/model"
	"taskmanager/usecase/presenter"
	"taskmanager/usecase/repository"
)

type taskInteractor struct {
	TaskRepository repository.TaskRepository
	TaskPresenter  presenter.TaskPresenter
}

type TaskInteractor interface {
	GetAll(t []*model.Task) ([]*model.Task, error)
	GetOne(t *model.Task, id uint) (*model.Task, error)
	GetConcurrently(t []*model.Task, isEven bool, items, itemsPerWorker int) ([]*model.Task, error)
}

var wg sync.WaitGroup

func NewTaskInteractor(r repository.TaskRepository, p presenter.TaskPresenter) TaskInteractor {
	return &taskInteractor{r, p}
}

func (ti *taskInteractor) GetAll(t []*model.Task) ([]*model.Task, error) {
	t, err := ti.TaskRepository.FindAll(t)
	if err != nil {
		return nil, err
	}

	return ti.TaskPresenter.ResponseTasks(t), nil
}

func (ti *taskInteractor) GetOne(t *model.Task, id uint) (*model.Task, error) {
	t, err := ti.TaskRepository.Find(t, id)
	if err != nil {
		return nil, err
	}

	return ti.TaskPresenter.ResponseTask(t), nil
}

func (ti *taskInteractor) GetConcurrently(t []*model.Task, onlyEven bool, items, itemsPerWorker int) ([]*model.Task, error) {
	allTasks, err := ti.TaskRepository.FindAll(t) // Load all the Tasks in Memory
	if err != nil {
		return nil, err
	}

	t = []*model.Task{}                            // Slice which will contain filtered tasks
	tasks := make(chan *model.Task, len(allTasks)) // Channel with size of all tasks
	results := make(chan *model.Task, items)       // Channel of results with size of maximum number of items

	// Calculate number of workers and run them
	numberOfWorkers := int(math.Ceil(float64(items) / float64(itemsPerWorker)))

	wg.Add(numberOfWorkers) // Add all the workers to WaitGroup

	for w := 1; w <= numberOfWorkers; w++ {
		go worker(w, itemsPerWorker, onlyEven, tasks, results)
	}
	fmt.Println(numberOfWorkers, "workers started")

	for _, task := range allTasks { // Send all tasks to Channel
		tasks <- task
	}
	close(tasks)
	fmt.Printf("All tasks were send to channel\n")

	go func() { // Wait all the workers finish to close results channels
		wg.Wait()
		close(results)
	}()

	for task := range results { // Read Result channel
		t = append(t, task)

		if len(t) == items {
			break
		}
	}

	return ti.TaskPresenter.ResponseTasks(t), nil
}

func worker(id, maxItems int, onlyEven bool, tasks <-chan *model.Task, results chan<- *model.Task) {
	processedItems := 0

	for task := range tasks {
		fmt.Printf("Worker %v took task with ID: %v\n", id, task.ID)

		if onlyEven {
			if task.ID%2 == 0 {
				results <- task
			}
		} else {
			if task.ID%2 != 0 {
				results <- task
			}
		}

		fmt.Printf("Worker %v processed task with ID: %v\n", id, task.ID)

		processedItems += 1
		fmt.Printf("Worker %v has processed %v tasks\n", id, processedItems)
		if processedItems == maxItems {
			fmt.Printf("Worker %v has reached its maximun number of tasks\n", id)
			break
		}

	}
	wg.Done()
	fmt.Println("Worker", id, "has been closed")
}
