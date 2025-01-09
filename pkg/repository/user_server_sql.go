package repository

import (
	"fmt"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type AuthServerSql struct {
	db *sqlx.DB
}

// func NewAuthServerSql(db *sqlx.DB) *AuthServerSql {
// 	return &AuthServerSql{db: db}
// }

func NewAuthServerSql() *AuthServerSql {
	return &AuthServerSql{}
}

func (r *AuthServerSql) CreateUser(user arturproject.User) error {
	query := fmt.Sprintf("INSERT INTO %s (Id, DeviceId, Locale, AppVersion, CreatedAt, LastSeen) values ($1, $2, $3, $4, $5, $6)", usersTable)

	_, err := r.db.Exec(query, user.Id, user.DeviceId, user.Locale, user.AppVersion, user.CreatedAt, user.LastSeen)

	return err
}

func (r *AuthServerSql) GetUser(id string) (arturproject.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (Id, DeviceId, Locale, AppVersion, CreatedAt, LastSeen) values ($1, $2, $3, $4, $5, $6)", usersTable)

	row := r.db.QueryRow(query, user.Id, user.DeviceId, user.Locale, user.AppVersion, user.CreatedAt, user.LastSeen)

	return err
}
