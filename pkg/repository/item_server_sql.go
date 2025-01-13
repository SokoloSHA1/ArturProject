package repository

import (
	"fmt"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type ItemServerSql struct {
	db *sqlx.DB
}

func NewItemServerSql(db *sqlx.DB) *ItemServerSql {
	return &ItemServerSql{db: db}
}

func (r *ItemServerSql) DeleteItems(user arturproject.User, items []string) error {
	for _, id := range items {
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s' AND UserId = '%s'", tagsTable, id, user.Id)
		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CategoryServerSql) CreateItem(item arturproject.Item) error {
	query := fmt.Sprintf("INSERT INTO \"%s\" (Id, UserId, Title, OrderMode, CreatedAt, UpdatedAt) VALUES ('%s', '%s', '%s', '%d', '%s', '%s')",
		categoriesTable, category.Id, category.UserId, category.Title, category.OrderMode, category.CreatedAt, category.UpdatedAt)
	fmt.Println(query)
	_, err := r.db.Exec(query)

	return err
}
