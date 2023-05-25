package repository

import (
	"github.com/jmoiron/sqlx"
	"group-app-service/pkg/model"
)

type Administration interface {
	GetAllUsers() ([]model.User, error)
	UpdateUsers(users model.UpdateUsers) error
	DeleteUsers() error
}

type Repository struct {
	Administration
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Administration: NewAdminGroupPostgres(db),
	}
}
