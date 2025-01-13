package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

type Config struct {
	Server   string
	Port     string
	User     string
	Password string
	DBName   string
}

const (
	usersTable      = "User"
	categoriesTable = "Category"
	itemsTable      = "Item"
	tagsTable       = "Tag"
	itemTagsTable   = "ItemTag"
	orderMode       = "OrderMode"
)

func NewSQLServerDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlserver", fmt.Sprintf("server=%s;port=%s;database=%s",
		cfg.Server, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
