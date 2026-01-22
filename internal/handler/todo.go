package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/HeavyMaverick/todo-api-go/internal/model"
	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	var tasks []model.Task = []model.Task{}
	// remove
	tasks = append(tasks, model.Task{
		ID:          1,
		Title:       "Sample Title",
		Description: "Sample Desc",
		Completed:   false,
		CreatedAt:   time.Now(),
	})
	ctx.JSON(http.StatusOK, tasks)
}
func CreateTask(ctx *gin.Context) {
	log.Println("Task created")
	ctx.JSON(http.StatusCreated, gin.H{"status": "201 - created"})
}
