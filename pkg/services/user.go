package services

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
)

type AuthService struct {
	repo repository.TodoUser
}

func NewAuthService(repo repository.TodoUser) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user arturproject.User) error {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(id string) (arturproject.User, error) {
	return s.repo.GetUser(id)
}

func (s *AuthService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
