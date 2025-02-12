package services

import (
	"github.com/idiarso/belajar-git/src/config"
)

type UserService struct {
	config *config.Config
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		config: cfg,
	}
}
