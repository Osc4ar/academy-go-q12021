package model

type Task struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	Completed   bool   `json:"completed"`
	DueDate     string `json:"due_date"`
	WorkingTime uint
}
