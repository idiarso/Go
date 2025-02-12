package routes

import (
    "github.com/gin-gonic/gin"
    "your_project/src/api/services"
)

func ReportRoutes(router *gin.Engine) {
    router.GET("/reports/attendance", func(c *gin.Context) {
        report := services.GenerateAttendanceReport()
        c.JSON(200, report)
    })

    router.GET("/reports/students", func(c *gin.Context) {
        report := services.GenerateStudentPerformanceReport()
        c.JSON(200, report)
    })
}
