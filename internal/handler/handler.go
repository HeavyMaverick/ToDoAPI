package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"ToDoApi/internal/model"

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
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// addTask(task)
	log.Println("Task created")
	ctx.JSON(http.StatusCreated, gin.H{"status": "201 - created"})
}
