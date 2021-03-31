package datastore

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"taskmanager/domain/model"
)

type dbType struct {
	tasks []*model.Task
}

// DB interface contains the methods needed for query data
type DB interface {
	FindByID(id uint) (*model.Task, error)
	FindAll() []*model.Task
	SaveRecords([]*model.Task) error
}

// NewDB returns a new DB instance
func NewDB() (DB, error) {
	csvfile, err := os.Open("tasks.csv")
	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	tasks, err := populateTasks(csv.NewReader(csvfile))
	if err != nil {
		return nil, fmt.Errorf("Could not open DB")
	}

	populatedDB := &dbType{tasks: tasks}

	return populatedDB, nil
}

func (d *dbType) RefreshDB() error {
	csvfile, err := os.Open("tasks.csv")
	if err != nil {
		return err
	}
	defer csvfile.Close()

	tasks, err := populateTasks(csv.NewReader(csvfile))
	if err != nil {
		return fmt.Errorf("Could not open DB")
	}

	d.tasks = tasks
	return nil
}

func (d *dbType) FindByID(id uint) (*model.Task, error) {
	for _, task := range d.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("Could not find ID: %v", id)
}

func (d *dbType) FindAll() []*model.Task {
	return d.tasks
}

func (d *dbType) SaveRecords(tasks []*model.Task) error {
	defer d.RefreshDB()

	file, err := os.Create("tasks.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		err := writer.Write(convertTaskToSlice(task))
		if err != nil {
			return err
		}
	}

	return nil
}

func convertTaskToSlice(t *model.Task) []string {
	record := []string{
		strconv.Itoa(int(t.ID)),
		t.Content,
		strconv.FormatBool(t.Completed),
		t.DueDate,
		strconv.Itoa(int(t.WorkingTime)),
	}

	return record
}

func populateTasks(reader *csv.Reader) ([]*model.Task, error) {
	tasks := []*model.Task{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
			continue
		}

		task, err := recordToTask(record)
		if err != nil {
			log.Print(err)
			continue
		}

		tasks = append(tasks, &task)
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("CSV file could not be parsed")
	}

	return tasks, nil
}

func recordToTask(record []string) (model.Task, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		log.Print(err)
		return model.Task{}, fmt.Errorf("Invalid value for ID: %v", record[0])
	}

	completed, err := strconv.ParseBool(record[2])
	if err != nil {
		log.Print(err)
		return model.Task{}, fmt.Errorf("Invalid value for Completed: %v", record[2])
	}

	workingTime, err := strconv.Atoi(record[4])
	if err != nil {
		log.Print(err)
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
