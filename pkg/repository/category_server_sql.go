package repository

import (
	"fmt"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type CategoryServerSql struct {
	db *sqlx.DB
}

func NewCategoryServerSql(db *sqlx.DB) *CategoryServerSql {
	return &CategoryServerSql{db: db}
}

func (r *CategoryServerSql) UpdateCategories(categories []arturproject.Category) error {
	//TODO update categories
	return nil
}

func (r *CategoryServerSql) DeleteCategories(cagories []string) error {
	for _, id := range cagories {
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'", usersTable, id)
		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
