package repository

import (
	"database/sql"
	"fmt"
	"strings"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type CategoryServerSql struct {
	db *sqlx.DB
}

func NewCategoryServerSql(db *sqlx.DB) *CategoryServerSql {
	return &CategoryServerSql{db: db}
}

func (r *CategoryServerSql) UpdateCategories(category arturproject.Category) error {
	setValues := make([]string, 0)
	if category.Title != "" {
		setValues = append(setValues, fmt.Sprintf("Title = '%s'", category.Title))
	}

	if category.OrderMode != 0 {
		setValues = append(setValues, fmt.Sprintf("OrderMode = '%d'", category.OrderMode))
	}

	setValues = append(setValues, fmt.Sprintf("UpdatedAt = '%s'", category.UpdatedAt))

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE \"%s\" SET %s WHERE Id = '%s' AND UserId = '%s'",
		categoriesTable, setQuery, category.Id, category.UserId)

	_, err := r.db.Exec(query)

	return err
}

func (r *CategoryServerSql) DeleteCategories(user arturproject.User, categories []string) error {
	for _, id := range categories {
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s' AND UserId = '%s'", categoriesTable, id, user.Id)
		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CategoryServerSql) CreateCategory(category arturproject.Category) error {
	query := fmt.Sprintf("INSERT INTO \"%s\" (Id, UserId, Title, OrderMode, CreatedAt, UpdatedAt) VALUES ('%s', '%s', '%s', '%d', '%s', '%s')",
		categoriesTable, category.Id, category.UserId, category.Title, category.OrderMode, category.CreatedAt, category.UpdatedAt)
	fmt.Println(query)
	_, err := r.db.Exec(query)

	return err
}

func (r *CategoryServerSql) CheckCategories(category arturproject.Category) (bool, error) {
	var ok string
	query := fmt.Sprintf("SELECT Title FROM \"%s\" WHERE Id = '%s' AND UserId = '%s'", categoriesTable, category.Id, category.UserId)
	err := r.db.QueryRow(query).Scan(&ok)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}
	return true, nil
}

func (r *CategoryServerSql) GetCategories(userId string) ([]arturproject.Category, error) {
	var categories []arturproject.Category
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE UserId = '%s'", categoriesTable, userId)
	err := r.db.Select(&categories, query)
	return categories, err
}
