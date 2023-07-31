package service

import (
	"github.com/JamshedJ/todo/models"
	"github.com/JamshedJ/todo/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return s.repo.CreateUser(user)
}
