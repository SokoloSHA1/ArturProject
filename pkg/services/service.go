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
	DeleteItems(items []string) error
	UpdateItems(items []arturproject.Item) error
	GetItems(categories []arturproject.Category) ([]arturproject.Item, error)
}

type TodoTag interface {
	DeleteTags(tags []string) error
	UpdateTags(tags []arturproject.Tag) error
	GetTags(categories []arturproject.Category) ([]arturproject.Tag, error)
}

type TodoItemTag interface {
	DeleteItemTags(itemTags []string) error
	UpdateItemTags(itemTags []arturproject.ItemTag) error
	GetItemTags(items []arturproject.Item) ([]arturproject.ItemTag, error)
}

type Service struct {
	TodoUser
	TodoCategory
	TodoItem
	TodoTag
	TodoItemTag
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoUser:     NewAuthService(repos.TodoUser),
		TodoCategory: NewCategoryService(repos.TodoCategory),
		TodoItem:     NewItemService(repos.TodoItem),
		TodoTag:      NewTagService(repos.TodoTag),
		TodoItemTag:  NewItemTagService(repos.TodoItemTag),
	}
}
