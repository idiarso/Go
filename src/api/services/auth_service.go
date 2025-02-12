package services

import (
	"github.com/idiarso/belajar-git/src/config"
)

type AuthService struct {
	config *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		config: cfg,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	// TODO: Implement login logic
	return "", nil
}

func (s *AuthService) Register(username, password, email string) error {
	// TODO: Implement registration logic
	return nil
}
