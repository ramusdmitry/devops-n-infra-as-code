package service

import (
	"auth-app-service/pkg/metrics"
	authApp "auth-app-service/pkg/model"
	"auth-app-service/pkg/repository"
)

type Authorization interface {
	CreateUser(user authApp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	GetUser(username, password string) (authApp.Profile, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, metrics *metrics.Metrics) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, metrics),
	}
}
