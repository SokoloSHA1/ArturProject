package repository

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	CreateUser(user arturproject.User) error
	GetUser(id string) (arturproject.User, error)
	DeleteUser(id string) error
}

type TodoCategory interface {
	UpdateCategories(category arturproject.Category) error
	DeleteCategories(user arturproject.User, categories []string) error
	CreateCategory(category arturproject.Category) error
	CheckCategories(category arturproject.Category) (bool, error)
	GetCategories(userId string) ([]arturproject.Category, error)
}

type TodoItem interface {
	DeleteItems(items []string) error
	CheckItems(item arturproject.Item) (bool, error)
	CreateItem(item arturproject.Item) error
	UpdateItem(item arturproject.Item) error
	GetItems(categories []arturproject.Category) ([]arturproject.Item, error)
}

type TodoTag interface {
	DeleteTags(tags []string) error
	CheckTags(tag arturproject.Tag) (bool, error)
	CreateTag(tag arturproject.Tag) error
	UpdateTag(tag arturproject.Tag) error
	GetTags(categories []arturproject.Category) ([]arturproject.Tag, error)
}

type TodoItemTag interface {
	DeleteItemTags(itemTags []string) error
	CreateItemTags(itemTags []arturproject.ItemTag) error
	GetItemTags(items []arturproject.Item) ([]arturproject.ItemTag, error)
}

type Repository struct {
	TodoUser
	TodoCategory
	TodoItem
	TodoTag
	TodoItemTag
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser:     NewAuthServerSql(db),
		TodoCategory: NewCategoryServerSql(db),
		TodoItem:     NewItemServerSql(db),
		TodoTag:      NewTagServerSql(db),
		TodoItemTag:  NewItemTagServerSql(db),
	}
}
