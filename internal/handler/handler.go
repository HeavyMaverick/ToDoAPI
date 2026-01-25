package handler

import (
	"log"
	"net/http"
	"strconv"

	"ToDoApi/internal/model"
	"ToDoApi/internal/service"

	"github.com/gin-gonic/gin"
)

var taskService service.TaskService

func GetTasks(ctx *gin.Context) {
	tasks, err := taskService.GetAllTasks()
	if err != nil {
		log.Println("Error", err)
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := taskService.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
	param, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = taskService.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func UpdateTask(ctx *gin.Context) {
	var task model.Task
	param, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(param)
	err = ctx.ShouldBindJSON(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	taskService.UpdateTask(id, &task)
}

func SetTaskService(ts service.TaskService) {
	taskService = ts
}
