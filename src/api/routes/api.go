package routes

import (
	"github.com/gin-gonic/gin"
	"belajar-git/src/api/handlers"
	"belajar-git/src/api/middleware"
	"belajar-git/src/api/services"
)

// APIRoutes sets up all API routes
func APIRoutes(router *gin.Engine, services *services.Services) {
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login(services.Auth))
			auth.POST("/register", handlers.Register(services.Auth))
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// User routes
			user := protected.Group("/users")
			{
				user.GET("/profile", handlers.GetProfile(services.User))
				user.PUT("/profile", handlers.UpdateProfile(services.User))
			}

			// Attendance routes
			attendance := protected.Group("/attendance")
			{
				attendance.POST("/check-in", handlers.CheckIn(services.Attendance))
				attendance.GET("/history", handlers.GetAttendanceHistory(services.Attendance))
			}
		}
	}
}
