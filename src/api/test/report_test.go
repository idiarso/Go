package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/idiarso/belajar-git/src/api/routes"
)

func TestGetAttendanceReport(t *testing.T) {
	router := gin.Default()
	routes.ReportRoutes(router)

	req, _ := http.NewRequest("GET", "/reports/attendance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}

func TestGetStudentPerformanceReport(t *testing.T) {
	router := gin.Default()
	routes.ReportRoutes(router)

	req, _ := http.NewRequest("GET", "/reports/students", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
