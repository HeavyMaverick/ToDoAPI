package main

import (
	"log"
	"net/http"

	h "ToDoApi/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/", func(ctx *gin.Context) {
		endpoints := map[string]string{
			"health":    "GET  /health",
			"get_tasks": "GET  /api/v1/tasks",
			// "create_task":   "POST /api/v1/tasks",
			// "get_task":      "GET  /api/v1/tasks/:id",
			// "update_task":   "PUT  /api/v1/tasks/:id",
			// "delete_task":   "DELETE /api/v1/tasks/:id",
			// "complete_task": "PATCH /api/v1/tasks/:id/complete",
		}

		ctx.JSON(http.StatusOK, gin.H{
			"service":   "ToDo API",
			"version":   "1.0",
			"endpoints": endpoints,
			"docs":      "http://" + ctx.Request.Host + "/",
			// "docs":      "http://" + ctx.Request.Host + "/docs",
		})
	})

	{
		v1 := r.Group("/api/v1")
		v1.GET("/tasks", h.GetTasks)
		v1.POST("/tasks", h.CreateTask)
	}

	log.Println("Server starting on :8080")
	log.Println("📄 Documentation: http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
