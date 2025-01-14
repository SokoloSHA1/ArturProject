package repository

import (
	"fmt"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type ItemTagServerSql struct {
	db *sqlx.DB
}

func NewItemTagServerSql(db *sqlx.DB) *ItemTagServerSql {
	return &ItemTagServerSql{db: db}
}

func (r *ItemTagServerSql) DeleteItemTags(itemTags []string) error {
	for _, id := range itemTags {
		query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'", itemTagsTable, id)

		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ItemTagServerSql) CreateItemTags(itemTags []arturproject.ItemTag) error {
	for _, itemTag := range itemTags {
		query := fmt.Sprintf("INSERT INTO \"%s\" (Id, ItemId, TagId, CreatedAt) VALUES ('%s', '%s', '%s', '%s')",
			itemTagsTable, itemTag.Id, itemTag.ItemId, itemTag.TagId, itemTag.CreatedAt)
		_, err := r.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ItemTagServerSql) GetItemTags(items []arturproject.Item) ([]arturproject.ItemTag, error) {
	var allItemTags []arturproject.ItemTag

	for _, item := range items {
		var itemTags []arturproject.ItemTag

		query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE ItemId = '%s'", itemTagsTable, item.Id)
		err := r.db.Select(&itemTags, query)
		if err != nil {
			return nil, err
		}

		allItemTags = append(allItemTags, itemTags...)
	}

	return allItemTags, nil

}
