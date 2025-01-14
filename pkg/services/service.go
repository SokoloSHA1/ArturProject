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
	DeleteCategories(user arturproject.User, categories []string) error
	GetCategories(userId string) ([]arturproject.Category, error)
}

type TodoItem interface {
	DeleteItems(user arturproject.User, items []string) error
	UpdateItems(items []arturproject.Item) error
	GetItems(userId string) ([]arturproject.Item, error)
}

type TodoTag interface {
	DeleteTags(user arturproject.User, tags []string) error
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
