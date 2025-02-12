package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/idiarso/belajar-git/src/api/services"
    "github.com/sirupsen/logrus"
)

type AuthHandler struct {
    authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
    var input services.RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    user, err := h.authService.Register(input)
    if err != nil {
        logrus.WithError(err).Error("Registration failed")
        if services.IsValidationError(err) {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        }
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input services.LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    accessToken, refreshToken, err := h.authService.Login(input, c.ClientIP())
    if err != nil {
        logrus.WithError(err).Error("Login failed")
        if services.IsAuthenticationError(err) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken, "token_type": "Bearer"})
}

func TestAuthHandler_Register(t *testing.T) {
    // Setup
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name         string
        input        map[string]interface{}
        expectedCode int
        expectedError bool
    }{
        {
            name: "Valid Registration",
            input: map[string]interface{}{
                "username": "testuser",
                "password": "Test@123",
                "role":     "user",
            },
            expectedCode: http.StatusCreated,
            expectedError: false,
        },
        {
            name: "Invalid Input - Missing Password",
            input: map[string]interface{}{
                "username": "testuser",
                "role":     "user",
            },
            expectedCode: http.StatusBadRequest,
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            router := gin.New()
            authService := services.NewAuthService(nil) // Mock DB
            handler := NewAuthHandler(authService)

            router.POST("/register", handler.Register)

            body, _ := json.Marshal(tt.input)
            req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")

            w := httptest.NewRecorder()
            router.ServeHTTP(w, req)

            assert.Equal(t, tt.expectedCode, w.Code)

            if tt.expectedError {
                var response map[string]string
                err := json.Unmarshal(w.Body.Bytes(), &response)
                assert.NoError(t, err)
                assert.Contains(t, response, "error")
            }
        })
    }
}

// Additional tests for Login and RefreshToken can be added similarly
