package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/idiarso/belajar-git/src/api/services"
)

func AttendanceRoutes(router *gin.Engine) {
    router.POST("/attendance", func(c *gin.Context) {
        var json struct {
            UserID int    `json:"user_id"`
            Status  string `json:"status"`
        }
        if err := c.ShouldBindJSON(&json); err == nil {
            record := services.MarkAttendance(json.UserID, json.Status)
            c.JSON(200, record)
        } else {
            c.JSON(400, gin.H{"error": "Invalid input"})
        }
    })

    router.GET("/attendance", func(c *gin.Context) {
        records := services.GetAttendance()
        c.JSON(200, records)
    })
}
