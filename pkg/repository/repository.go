package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Task interface {
}

type Repository struct {
	Authorization
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
