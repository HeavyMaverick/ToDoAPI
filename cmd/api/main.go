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
	// log.Println("Trying to load config")
	// cfg, err := config.LoadConfig(".")
	// // cfg, err := config.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("Error loading config:", err)
	// }
	log.Println("Loading configuration from environment variables")
	cfg := config.LoadConfig()
	log.Printf("Config loaded: DB=%s:%s", cfg.DBHost, cfg.DBPort)

	log.Println("Trying connect to db")
	db, err := database.ConnectDB(&cfg)
	if err != nil {
		log.Fatal("Error connecting database", err)
	}
	log.Println("DB connected")

	log.Println("Trying to migrate db")
	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed")

	rep := repository.NewPostgresTaskRepository(db)
	// rep := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(rep)
	h.SetTaskService(taskService)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"service": "ToDo API",
			"version": "1.0.0",
			"docs":    "https://github.com/heavymaverick/ToDoApi",
			"endpoints": []string{
				"GET    /health",
				"GET    /api/v1/tasks",
				"GET    /api/v1/tasks/:id",
				"POST   /api/v1/tasks",
				"PUT    /api/v1/tasks/:id",
				"DELETE /api/v1/tasks/:id",
			},
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
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Server starting error", err)
	}
}
