package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/middleware"
	"github.com/idiarso/belajar-git/src/api/routes"
	"github.com/idiarso/belajar-git/src/api/services"
	"github.com/idiarso/belajar-git/src/config"
)

type Post struct {
	Title   string
	Content string
}

type PageData struct {
	Posts []Post
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize router
	router := gin.Default()

	// Initialize services
	srv := &services.Services{
		Auth:       services.NewAuthService(cfg),
		User:       services.NewUserService(cfg),
		Attendance: services.NewAttendanceService(cfg),
	}

	// Apply middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.APIRoutes(router, srv)

	// Render index template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.gohtml"))
		data := PageData{
			Posts: []Post{
				{Title: "Post 1", Content: "Content for post 1"},
				{Title: "Post 2", Content: "Content for post 2"},
			},
		}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
