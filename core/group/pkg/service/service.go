package service

import (
	"group-app-service/pkg/model"
	"group-app-service/pkg/repository"
)

type Authorization interface {
}

type Administration interface {
	GetAllUsers() ([]model.User, error)
	UpdateUsers(users model.UpdateUsers) error
	DeleteUsers() error
}

type Service struct {
	Administration
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Administration: NewAdminService(repos.Administration),
	}
}
