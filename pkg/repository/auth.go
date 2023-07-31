package repository

import (
	"fmt"
	"github.com/JamshedJ/todo/models"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (p *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", userTable)

	row := p.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
