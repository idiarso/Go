package main

import (
    "github.com/gin-gonic/gin"
    "github.com/idiarso/belajar-git/src/api/routes"
    "github.com/idiarso/belajar-git/src/api/middleware"
)

func main() {
    router := gin.Default()

    // Apply CORS middleware
    router.Use(middleware.CORSMiddleware())

    // Register routes
    routes.AuthRoutes(router)
    routes.AttendanceRoutes(router)
    routes.StudentRoutes(router)

    // Start the server
    router.Run(":8080")
}
