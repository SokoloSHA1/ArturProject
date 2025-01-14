package repository

import (
	"database/sql"
	"fmt"
	"strings"

	arturproject "github.com/SokoloSHA/ArturProject"
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
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE ItemId = '%s'", itemTagsTable, id)

		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'",
			itemsTable, id)

		_, err = r.db.Exec(query)
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

	query := fmt.Sprintf("INSERT INTO \"%s\" (Id, CategoryId, Title, Description, Rating, Rank, CreatedAt, UpdatedAt) VALUES ('%s', '%s', '%s', '%s', '%s', '%d', '%s', '%s')",
		itemsTable, item.Id, item.CategoryId, item.Title, item.Description, item.Rating, item.Rank, item.CreatedAt, item.UpdatedAt)
	fmt.Println(query)
	_, err := r.db.Exec(query)

	return err
}

func (r *ItemServerSql) UpdateItem(item arturproject.Item) error {
	setValues := make([]string, 0)
	if item.Title != "" {
		setValues = append(setValues, fmt.Sprintf("Title = '%s'", item.Title))
	}

	if item.Description != "" {
		setValues = append(setValues, fmt.Sprintf("Description = '%s'", item.Description))
	}

	if item.Rating != "" {
		setValues = append(setValues, fmt.Sprintf("Rating = '%s'", item.Rating))
	}

	if item.Rank != 0 {
		setValues = append(setValues, fmt.Sprintf("Rank = '%d'", item.Rank))
	}

	setValues = append(setValues, fmt.Sprintf("UpdatedAt = '%s'", item.UpdatedAt))

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE \"%s\" SET %s WHERE Id = '%s' AND CategoryId = '%s'",
		itemsTable, setQuery, item.Id, item.CategoryId)

	_, err := r.db.Exec(query)

	return err
}

func (r *ItemServerSql) GetItems(categories []arturproject.Category) ([]arturproject.Item, error) {
	var allItems []arturproject.Item

	for _, category := range categories {
		var categoryItems []arturproject.Item

		query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE CategoryId = '%s'", itemsTable, category.Id)
		err := r.db.Select(&categoryItems, query)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, categoryItems...)
	}

	return allItems, nil

}
