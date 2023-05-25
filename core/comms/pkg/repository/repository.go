package repository

import (
	commsApp "comms-app-service/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Comms interface {
	CreateComment(userId int, comment commsApp.Comment) (int, error)
	GetCommentsByPostId(postId int) ([]commsApp.Comment, error)
	DeleteCommentById(userId int, commentId int) error
	UpdateCommentById(userId int, commentId int, comment commsApp.UpdateCommentInput) error
	GetAllComments() ([]commsApp.Comment, error)
}

type Authorization interface {
}

type Repository struct {
	Comms
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Comms:         NewCommsPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
