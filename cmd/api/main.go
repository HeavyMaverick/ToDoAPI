package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	//Подгружаем конфиг
	log.Println("Loading configuration from environment variables")
	cfg := config.LoadConfig()
	log.Printf("Config loaded: DB=%s:%s", cfg.DBHost, cfg.DBPort)

	//Коннектимся к базе данных
	log.Println("Trying connect to db")
	db, err := database.ConnectDB(&cfg)
	if err != nil {
		log.Fatal("Error connecting database", err)
	}
	log.Println("DB connected")

	//Мигрируем
	log.Println("Trying to migrate db")
	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed")

	rep := repository.NewPostgresTaskRepository(db)
	userRepo := repository.NewPostgresUserRepository(db)
	// для хранения в памяти
	// rep := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(rep, userRepo)
	h.SetTaskService(taskService)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Получаем абсолютный путь к корню проекта
	// Определяем путь к папке templates относительно main.go
	execPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory:", err)
	}

	// Для cmd/api/main.go, templates находятся в ../../templates
	templatesPath := filepath.Join(execPath, "../../templates")
	staticPath := filepath.Join(execPath, "../../static")

	// Посмотреть откуда подтягивает файлы
	log.Printf("Looking for templates in: %s", templatesPath)
	log.Printf("Looking for static files in: %s", staticPath)

	// Проверяем существование папок
	if _, err := os.Stat(templatesPath); os.IsNotExist(err) {
		log.Printf("Templates folder not found at %s", templatesPath)
		templatesPath = filepath.Join(execPath, "templates")
		if err := os.MkdirAll(templatesPath, 0755); err != nil {
			log.Printf("Failed to create templates directory: %v", err)
		}
	}

	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		log.Printf("Static folder not found at %s", staticPath)
		staticPath = filepath.Join(execPath, "static")
		if err := os.MkdirAll(staticPath, 0755); err != nil {
			log.Printf("Failed to create static directory: %v", err)
		}
	}

	// Подгружаем HTML шаблоны
	r.LoadHTMLGlob(filepath.Join(templatesPath, "*.html"))
	// Настраиваем статические файлы
	r.Static("/static", staticPath)
	r.StaticFile("/favicon.ico", filepath.Join(staticPath, "favicon.ico"))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"service": "ToDo API",
	// 		"version": "1.0.0",
	// 		"docs":    "https://github.com/heavymaverick/ToDoApi",
	// 		"endpoints": []string{
	// 			"GET    /health",
	// 			"GET    /api/v1/tasks",
	// 			"GET    /api/v1/tasks/:id",
	// 			"POST   /api/v1/tasks",
	// 			"PUT    /api/v1/tasks/:id",
	// 			"DELETE /api/v1/tasks/:id",
	// 		},
	// 	})
	// })

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "ToDoApi",
			"version": "1.0.0",
		})
	})

	//html test controller
	r.GET("/testpage", h.TestpageGET)

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
