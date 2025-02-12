package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/idiarso/belajar-git/src/api/services"
)

func InternshipRoutes(router *gin.Engine) {
    router.POST("/internships", func(c *gin.Context) {
        var json struct {
            CompanyName string `json:"company_name"`
            StudentID   int    `json:"student_id"`
            StartDate   string `json:"start_date"`
            EndDate     string `json:"end_date"`
            Status      string `json:"status"`
        }
        if err := c.ShouldBindJSON(&json); err == nil {
            internship := services.AddInternship(json.CompanyName, json.StudentID, json.StartDate, json.EndDate, json.Status)
            c.JSON(200, internship)
        } else {
            c.JSON(400, gin.H{"error": "Invalid input"})
        }
    })

    router.GET("/internships", func(c *gin.Context) {
        internships := services.GetInternships()
        c.JSON(200, internships)
    })

    router.GET("/internships/:id", func(c *gin.Context) {
        id := c.Param("id")
        // Convert id to int and get internship
        internship, found := services.GetInternshipByID(id)
        if found {
            c.JSON(200, internship)
        } else {
            c.JSON(404, gin.H{"error": "Internship not found"})
        }
    })
}
