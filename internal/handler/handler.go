package handler

import (
	"net/http"
	"strconv"

	"ToDoApi/internal/model"
	"ToDoApi/internal/service"

	"github.com/gin-gonic/gin"
)

var taskService service.TaskService

func SetTaskService(ts service.TaskService) {
	taskService = ts
}

// var (
// 	ErrInvalidId = errors.New("Invalid task ID").Error()
// 	ErrNotFound  = errors.New("Task not found").Error()
// )

func GetTasks(ctx *gin.Context) {
	tasks, err := taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := taskService.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

func GetTask(ctx *gin.Context) {
	param, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidId})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := taskService.GetTask(id)
	if err != nil {
		// ctx.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound})
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
	param, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidId})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = taskService.DeleteTask(id)
	if err != nil {
		// ctx.JSON(http.StatusNotFound, gin.H{"error 1": ErrNotFound})
		ctx.JSON(http.StatusNotFound, gin.H{"error 1": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func UpdateTask(ctx *gin.Context) {
	param, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidId})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := taskService.UpdateTask(id, &task); err != nil {
		if err == service.ErrEmptyTitle || err == service.ErrTitleTooLong {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			// ctx.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound})
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		}
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func TestpageGET(ctx *gin.Context) {
	tasks, err := taskService.GetAllTasks()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError,
			"error.html",
			gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK,
		"testpage.html",
		gin.H{
			"title":    "Test page",
			"allTasks": tasks,
		})
}
