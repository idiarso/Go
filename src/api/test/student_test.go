package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/routes"
)

func TestAddStudent(t *testing.T) {
	router := gin.Default()
	routes.StudentRoutes(router)

	json := `{"name":"John Doe","age":20,"class":"10A"}`
	req, _ := http.NewRequest("POST", "/students", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}

func TestGetStudents(t *testing.T) {
	router := gin.Default()
	routes.StudentRoutes(router)

	req, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
