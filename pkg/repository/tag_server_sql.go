package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TagServerSql struct {
	db *sqlx.DB
}

func NewTagServerSql(db *sqlx.DB) *TagServerSql {
	return &TagServerSql{db: db}
}

func (r *TagServerSql) DeleteTags(tags []string) error {
	for _, id := range tags {
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'", tagsTable, id)
		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
