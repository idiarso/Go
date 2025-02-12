package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/routes"
)

func TestRegister(t *testing.T) {
	router := gin.Default()
	routes.AuthRoutes(router)

	json := `{"username":"testuser","password":"testpass","role":"student"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	routes.AuthRoutes(router)

	json := `{"username":"testuser","password":"testpass"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
