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

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(400, Response{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	// Call the service layer to register the user
	err := h.authService.Register(registerRequest.Username, registerRequest.Password, registerRequest.Email)
	if err != nil {
		c.JSON(500, Response{
			Status:  "error",
			Message: "Registration failed",
		})
		return
	}

	c.JSON(201, Response{
		Status:  "success",
		Message: "User registered successfully",
	})
}

// GetProfile retrieves the user's profile information
func (h *AuthHandler) GetProfile(c *gin.Context) {
	// Assume we have a way to get the user ID from the context
	userID := c.Param("id") // Example: get user ID from URL parameters

	profile, err := h.authService.GetProfile(userID)
	if err != nil {
		c.JSON(404, Response{
			Status:  "error",
			Message: "User not found",
		})
		return
	}

	c.JSON(200, Response{
		Status: "success",
		Data:   profile,
	})
}

// UpdateProfile updates the user's profile information
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var updateRequest struct {
		Username string `json:"username"`
		// Add other fields as necessary
	}

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(400, Response{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	// Assume we have a way to get the user ID from the context
	userID := c.Param("id")

	err := h.authService.UpdateProfile(userID, updateRequest.Username)
	if err != nil {
		c.JSON(500, Response{
			Status:  "error",
			Message: "Update failed",
		})
		return
	}

	c.JSON(200, Response{
		Status:  "success",
		Message: "Profile updated successfully",
	})
}

// CheckIn handles user check-in functionality
func (h *AuthHandler) CheckIn(c *gin.Context) {
	// Logic for handling check-in
	c.JSON(200, Response{
		Status:  "success",
		Message: "Checked in successfully",
	})
}

// GetAttendanceHistory retrieves the user's attendance history
func (h *AuthHandler) GetAttendanceHistory(c *gin.Context) {
	// Logic for retrieving attendance history
	c.JSON(200, Response{
		Status:  "success",
		Data:    []string{"Attendance record 1", "Attendance record 2"}, // Example data
	})
}
