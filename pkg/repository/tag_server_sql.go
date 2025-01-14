package repository

import (
	"database/sql"
	"fmt"
	"strings"

	arturproject "github.com/SokoloSHA/ArturProject"
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
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE TagId = '%s'", itemTagsTable, id)

		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'",
			tagsTable, id)

		_, err = r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *TagServerSql) CheckTags(tag arturproject.Tag) (bool, error) {
	var ok string
	query := fmt.Sprintf("SELECT Title FROM \"%s\" WHERE Id = '%s' AND CategoryId = '%s'", tagsTable, tag.Id, tag.CategoryId)
	err := r.db.QueryRow(query).Scan(&ok)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}
	return true, nil
}

func (r *TagServerSql) CreateTag(tag arturproject.Tag) error {

	query := fmt.Sprintf("INSERT INTO \"%s\" (Id, CategoryId, Title, CreatedAt, UpdatedAt) VALUES ('%s', '%s', '%s', '%s', '%s')",
		tagsTable, tag.Id, tag.CategoryId, tag.Title, tag.CreatedAt, tag.UpdatedAt)
	_, err := r.db.Exec(query)

	return err
}

func (r *TagServerSql) UpdateTag(tag arturproject.Tag) error {
	setValues := make([]string, 0)
	if tag.Title != "" {
		setValues = append(setValues, fmt.Sprintf("Title = '%s'", tag.Title))
	}

	setValues = append(setValues, fmt.Sprintf("UpdatedAt = '%s'", tag.UpdatedAt))

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE \"%s\" SET %s WHERE Id = '%s' AND CategoryId = '%s'",
		tagsTable, setQuery, tag.Id, tag.CategoryId)

	_, err := r.db.Exec(query)

	return err
}

func (r *TagServerSql) GetTags(categories []arturproject.Category) ([]arturproject.Tag, error) {
	var allTags []arturproject.Tag

	for _, category := range categories {
		var categoryTags []arturproject.Tag

		query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE CategoryId = '%s'", tagsTable, category.Id)
		err := r.db.Select(&categoryTags, query)
		if err != nil {
			return nil, err
		}

		allTags = append(allTags, categoryTags...)
	}

	return allTags, nil

}
