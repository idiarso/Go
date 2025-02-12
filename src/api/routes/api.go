package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/handlers"
	"github.com/idiarso/belajar-git/src/api/middleware"
	"github.com/idiarso/belajar-git/src/api/services"
)

// APIRoutes sets up all API routes
func APIRoutes(router *gin.Engine, services *services.Services) {
	api := router.Group("/api/v1")
	{
		// Auth routes
		authHandler := handlers.NewAuthHandler(services.Auth)
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// User routes
			user := protected.Group("/users")
			{
				user.GET("/profile", authHandler.GetProfile)
				user.PUT("/profile", authHandler.UpdateProfile)
			}

			// Attendance routes
			attendance := protected.Group("/attendance")
			{
				attendance.POST("/check-in", authHandler.CheckIn)
				attendance.GET("/history", authHandler.GetAttendanceHistory)
			}
		}
	}
}
