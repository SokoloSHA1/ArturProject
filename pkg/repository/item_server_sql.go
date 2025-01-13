package repository

import (
	"database/sql"
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

func (r *ItemServerSql) CheckItems(item arturproject.Item) (bool, error) {
	var ok string
	query := fmt.Sprintf("SELECT Title FROM \"%s\" WHERE Id = '%s' AND CategoryId = '%s'", itemsTable, item.Id, item.CategoryId)
	err := r.db.QueryRow(query).Scan(&ok)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}
	return true, nil
}

func (r *ItemServerSql) CreateItem(item arturproject.Item) error {

	query := fmt.Sprintf("INSERT INTO \"%s\" (Id, CategoryId, Title, Description, Rating, Rank, CreatedAt, UpdatedAt) VALUES ('%s', '%s', '%s', '%s', '%d', '%d', '%s', '%s')",
		itemsTable, item.Id, item.CategoryId, item.Title, item.Description, item.Rating, item.Rank, item.CreatedAt, item.UpdatedAt)
	fmt.Println(query)
	_, err := r.db.Exec(query)

	return err
}
