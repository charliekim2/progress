package models

type Task struct {
	Id    string
	Title string
	Desc  string
}

func NewTask(id string, title string, desc string) *Task {
	return &Task{
		Title: title,
		Desc:  desc,
		Id:    id,
	}
}
