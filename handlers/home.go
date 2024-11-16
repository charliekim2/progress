package handlers

import (
	"net/http"

	"github.com/charliekim2/progress/models"
	"github.com/charliekim2/progress/views/layouts"
)

func Home(w http.ResponseWriter, r *http.Request) {
	layout := layouts.Home(testData())
	layout.Render(r.Context(), w)
}

func testData() []*models.Goal {
	goal1 := models.NewGoal("1", "Read 10 books", "")
	task1 := models.NewTask("1", "Read LOTR", "Read all three books in LOTR")
	task2 := models.NewTask("2", "Read The Hobbit", "Read this before starting LOTR")
	goal1.AddTask(task1)
	goal1.AddTask(task2)

	goal2 := models.NewGoal("2", "Running", "Run 5 miles every day")
	goal3 := models.NewGoal("3", "Progress app", "Create a progress app and deploy it online")
	return []*models.Goal{goal1, goal2, goal3}
}
