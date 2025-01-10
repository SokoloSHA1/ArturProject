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
	UpdateCategories(categories []arturproject.Category) error
	DeleteCategories(cagories []string) error
}

type TodoItem interface {
	DeleteItems(items []string) error
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
