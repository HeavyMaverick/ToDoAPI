package main

import (
	"log"
	"net/http"

	"ToDoApi/internal/config"
	"ToDoApi/internal/database"
	h "ToDoApi/internal/handler"
	"ToDoApi/internal/repository"
	"ToDoApi/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Panicln("Error loading config:", err)
	}
	db, err := database.ConnectDB(&cfg)
	if err != nil {
		log.Panicln("Error connecting database", err)
	}
	err = database.AutoMigrate(db)
	if err != nil {
		log.Println("Migration failed:", err)
	}

	rep := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(rep)
	h.SetTaskService(taskService)

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/", func(ctx *gin.Context) {
		endpoints := map[string]string{
			"health":      "GET  /health",
			"get_tasks":   "GET  /api/v1/tasks",
			"get_task":    "GET  /api/v1/tasks/:id",
			"create_task": "POST /api/v1/tasks",
			"update_task": "PUT  /api/v1/tasks/:id",
			"delete_task": "DELETE /api/v1/tasks/:id",
		}

		ctx.JSON(http.StatusOK, gin.H{
			"service":   "ToDo API",
			"version":   "1.0",
			"endpoints": endpoints,
			"docs":      "http://" + ctx.Request.Host + "/",
		})
	})

	{
		v1 := r.Group("/api/v1")
		v1.GET("/tasks", h.GetTasks)
		v1.GET("/tasks/:id", h.GetTask)
		v1.POST("/tasks", h.CreateTask)
		v1.PUT("/tasks/:id", h.UpdateTask)
		v1.DELETE("/tasks/:id", h.DeleteTask)

	}

	log.Println("Server starting on :8080")
	log.Println("📄 Documentation: http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
