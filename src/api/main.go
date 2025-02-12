package main

import (
    "github.com/gin-gonic/gin"
    "github.com/idiarso/belajar-git/src/api/routes"
    "github.com/idiarso/belajar-git/src/api/middleware"
    "github.com/sirupsen/logrus"
    "github.com/go-redis/redis/v8"
    "github.com/idiarso/belajar-git/src/migrations"
    "github.com/idiarso/belajar-git/src/health"
    "os"
)

func main() {
    // Set environment variables
    os.Setenv("JWT_SECRET", "your-secure-secret-key")
    os.Setenv("DATABASE_URL", "your-database-connection-string")
    os.Setenv("REDIS_URL", "localhost:6379")

    // Initialize database connection
    // db := // your database initialization

    // Run migrations
    // if err := migrations.AutoMigrate(db); err != nil {
    //     log.Fatalf("Failed to run migrations: %v", err)
    // }

    // Initialize Redis
    redisClient := redis.NewClient(&redis.Options{
        Addr: os.Getenv("REDIS_URL"),
    })

    // Initialize health checker
    // healthChecker := health.NewHealthChecker(db, redisClient) (not implemented)

    // Create a logrus logger
    log := logrus.New()

    // Initialize router
    router := gin.Default()

    // Apply CORS middleware
    router.Use(middleware.CORSMiddleware())

    // Apply request/response logging middleware
    router.Use(Logger(log))

    // Health check endpoint
    // router.GET("/health", func(c *gin.Context) {
    //     health := healthChecker.Check(c.Request.Context())
    //     if health.Status == "healthy" {
    //         c.JSON(200, health)
    //     } else {
    //         c.JSON(503, health)
    //     }
    // })

    // Register routes
    routes.AuthRoutes(router)
    routes.AttendanceRoutes(router)
    routes.StudentRoutes(router)

    // Start the server
    router.Run(":8080")
}

func Logger(log *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path

        // Log request details
        log.Infof("Started %s %s", c.Request.Method, path)

        // Process request
        c.Next()

        // Log response details
        latency := time.Since(start)
        statusCode := c.Writer.Status()
        log.Infof("Completed %s %s with status %d in %v", c.Request.Method, path, statusCode, latency)
    }
}
