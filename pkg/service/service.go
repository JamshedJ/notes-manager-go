package service

import (
	"github.com/JamshedJ/todo/models"
	"github.com/JamshedJ/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Task interface {
}

type Service struct {
	Authorization
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
