package models

type Goal struct {
	Id    string
	Title string
	Desc  string
	Tasks []*Task
}

func NewGoal(id string, title string, desc string) *Goal {
	return &Goal{
		Id:    id,
		Title: title,
		Desc:  desc,
		Tasks: []*Task{},
	}
}

func (g *Goal) AddTask(task *Task) {
	g.Tasks = append(g.Tasks, task)
}
