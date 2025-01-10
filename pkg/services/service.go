package services

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
)

type TodoUser interface {
	CreateUser(user arturproject.User) error
	GetUser(id string) (arturproject.User, error)
	DeleteUser(id string) error
}

type TodoCategory interface {
}

type TodoItem interface {
}

type Service struct {
	TodoUser
	TodoCategory
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoUser: NewAuthService(repos.TodoUser),
	}
}
