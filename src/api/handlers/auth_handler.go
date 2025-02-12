package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, Response{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	token, err := h.authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(401, Response{
			Status:  "error",
			Message: "Invalid credentials",
		})
		return
	}

	c.JSON(200, Response{
		Status: "success",
		Data:   gin.H{"token": token},
	})
}
