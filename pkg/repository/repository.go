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
	DeleteItems(user arturproject.User, items []string) error
	CreateItem(item arturproject.Item) error
}

type TodoTag interface {
	DeleteTags(tags []string) error
}

type Repository struct {
	TodoUser
	TodoCategory
	TodoItem
	TodoTag
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser:     NewAuthServerSql(db),
		TodoCategory: NewCategoryServerSql(db),
		TodoItem:     NewItemServerSql(db),
		TodoTag:      NewTagServerSql(db),
	}
}
