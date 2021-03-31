package model

type Task struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	Completed   bool   `json:"completed"`
	DueDate     string
	WorkingTime uint
}
