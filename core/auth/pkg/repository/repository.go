package repository

import (
	authApp "auth-app-service/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user authApp.User) (int, error)
	GetUser(username, password string) (authApp.Profile, error)
	GetUserById(userId int) (authApp.Profile, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
