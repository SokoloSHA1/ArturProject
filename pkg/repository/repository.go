package repository

import arturproject "github.com/SokoloSHA/ArturProject"

type TodoUser interface {
	CreateUser(user arturproject.User) error
}

type TodoCategory interface {
}

type TodoItem interface {
}

type Repository struct {
	TodoUser
	TodoCategory
	TodoItem
}

// func NewRepository(db *sqlx.DB) *Repository
func NewRepository() *Repository {
	return &Repository{
		TodoUser: NewAuthServerSql(),
	}
}
