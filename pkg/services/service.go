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
	UpdateCategories(categories []arturproject.Category) error
	DeleteCategories(cagories []string) error
}

type TodoItem interface {
	DeleteItems(items []string) error
}

type TodoTag interface {
	DeleteTags(tags []string) error
}

type Service struct {
	TodoUser
	TodoCategory
	TodoItem
	TodoTag
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoUser:     NewAuthService(repos.TodoUser),
		TodoCategory: NewCategoryService(repos.TodoCategory),
		TodoItem:     NewItemService(repos.TodoItem),
		TodoTag:      NewTagService(repos.TodoTag),
	}
}
