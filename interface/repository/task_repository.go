package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"taskmanager/domain/model"
)

type taskRepository struct {
	reader *csv.Reader
}

// TaskRepository contains the Methods used with the Tasks data
type TaskRepository interface {
	FindAll(t []*model.Task) ([]*model.Task, error)
	Find(t *model.Task, id uint) (*model.Task, error)
}

// NewTaskRepository creates a new TaskRepository with the given csv.Reader
func NewTaskRepository(reader *csv.Reader) TaskRepository {
	return &taskRepository{reader}
}

func (tr *taskRepository) FindAll(t []*model.Task) ([]*model.Task, error) {
	tasks := []*model.Task{}

	for {
		record, err := tr.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}

		task, err := recordToTask(record)
		if err != nil {
			log.Fatal(err)
			continue
		}

		tasks = append(tasks, &task)
	}

	if len(tasks) < 0 {
		return nil, fmt.Errorf("CSV file could not be parsed")
	}

	return tasks, nil
}

func (tr *taskRepository) Find(t *model.Task, id uint) (*model.Task, error) {
	for {
		record, err := tr.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}

		task, err := recordToTask(record)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if task.ID == id {
			t = &task
			return t, nil
		}
	}

	return &model.Task{}, fmt.Errorf("Could not find record with ID: %v", id)
}

func recordToTask(record []string) (model.Task, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		log.Fatal(err)
		return model.Task{}, fmt.Errorf("Invalid value for ID: %v", record[0])
	}

	completed, err := strconv.ParseBool(record[2])
	if err != nil {
		log.Fatal(err)
		return model.Task{}, fmt.Errorf("Invalid value for Completed: %v", record[2])
	}

	workingTime, err := strconv.Atoi(record[4])
	if err != nil {
		log.Fatal(err)
		return model.Task{}, fmt.Errorf("Invalid value for WorkingTime: %v", record[4])
	}

	task := model.Task{
		ID:          uint(id),
		Content:     record[1],
		Completed:   completed,
		DueDate:     record[3],
		WorkingTime: uint(workingTime),
	}
	return task, nil
}
