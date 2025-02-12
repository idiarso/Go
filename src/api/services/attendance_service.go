package services

import (
	"github.com/idiarso/belajar-git/src/config"
)

type AttendanceService struct {
	config *config.Config
}

func NewAttendanceService(cfg *config.Config) *AttendanceService {
	return &AttendanceService{
		config: cfg,
	}
}
