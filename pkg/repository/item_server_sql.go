package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ItemServerSql struct {
	db *sqlx.DB
}

func NewItemServerSql(db *sqlx.DB) *ItemServerSql {
	return &ItemServerSql{db: db}
}

func (r *ItemServerSql) DeleteItems(items []string) error {
	for _, id := range items {
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'", usersTable, id)
		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
