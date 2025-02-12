package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/routes"
)

func TestMarkAttendance(t *testing.T) {
	router := gin.Default()
	routes.AttendanceRoutes(router)

	json := `{"user_id":1,"status":"present"}`
	req, _ := http.NewRequest("POST", "/attendance", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}

func TestGetAttendance(t *testing.T) {
	router := gin.Default()
	routes.AttendanceRoutes(router)

	req, _ := http.NewRequest("GET", "/attendance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
