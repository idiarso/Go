package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "github.com/idiarso/belajar-git/src/api/services"
)

func AuthRoutes(router *gin.Engine) {
    router.POST("/register", func(c *gin.Context) {
        var json struct {
            Username string `json:"username"`
            Password string `json:"password"`
            Role     string `json:"role"`
        }
        if err := c.ShouldBindJSON(&json); err == nil {
            user, err := services.Register(json.Username, json.Password, json.Role)
            if err == nil {
                c.JSON(200, user)
            } else {
                c.JSON(400, gin.H{"error": err.Error()})
            }
        } else {
            c.JSON(400, gin.H{"error": "Invalid input"})
        }
    })

    router.POST("/login", func(c *gin.Context) {
        var json struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&json); err == nil {
            token, err := services.Login(json.Username, json.Password)
            if err == nil {
                c.JSON(200, gin.H{"token": token})
            } else {
                c.JSON(400, gin.H{"error": err.Error()})
            }
        } else {
            c.JSON(400, gin.H{"error": "Invalid input"})
        }
    })
}
