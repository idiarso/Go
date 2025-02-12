package routes

import (
    "github.com/gin-gonic/gin"
    "your_project/src/api/services"
)

func StudentRoutes(router *gin.Engine) {
    router.POST("/students", func(c *gin.Context) {
        var json struct {
            Name  string `json:"name"`
            Age   int    `json:"age"`
            Class string `json:"class"`
        }
        if err := c.ShouldBindJSON(&json); err == nil {
            student := services.AddStudent(json.Name, json.Age, json.Class)
            c.JSON(200, student)
        } else {
            c.JSON(400, gin.H{"error": "Invalid input"})
        }
    })

    router.GET("/students", func(c *gin.Context) {
        students := services.GetStudents()
        c.JSON(200, students)
    })
}
