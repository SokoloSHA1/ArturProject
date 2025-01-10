package repository

import (
	"fmt"

	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/jmoiron/sqlx"
)

type AuthServerSql struct {
	db *sqlx.DB
}

func NewAuthServerSql(db *sqlx.DB) *AuthServerSql {
	return &AuthServerSql{db: db}
}

func (r *AuthServerSql) CreateUser(user arturproject.User) error {
	query := fmt.Sprintf("INSERT INTO \"%s\" (Id, DevicedID, Locale, CreatedAt, LastSeen, UpVersion) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
		usersTable, user.Id, user.DeviceId, user.Locale, user.CreatedAt, user.LastSeen, user.AppVersion)
	fmt.Println(query)
	_, err := r.db.Exec(query)

	return err
}

func (r *AuthServerSql) GetUser(id string) (arturproject.User, error) {
	var user arturproject.User
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE Id='%s'", usersTable, id)
	err := r.db.Get(&user, query)
	return user, err
}

func (r *AuthServerSql) DeleteUser(id string) error {
	query := fmt.Sprintf("DELETE FROM \"%s\" WHERE Id = '%s'", usersTable, id)
	_, err := r.db.Exec(query)

	return err
}
