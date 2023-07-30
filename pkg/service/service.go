package service

import "github.com/JamshedJ/todo/pkg/repository"

type Authorization interface {
}

type Task interface {
}

type Service struct {
	Authorization
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
