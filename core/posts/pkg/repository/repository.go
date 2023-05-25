package repository

import (
	"github.com/jmoiron/sqlx"
	postApp "posts-app-service/pkg/model"
)

type PostsList interface {
	GetAllPosts() ([]postApp.Post, error)
	CreatePost(userId int, post postApp.Post) (int, error)
	UpdatePostById(userId, postId int, post postApp.UpdatePostInput) error
	DeletePostById(userId, postId int) error
	GetAllComments(postId int) ([]postApp.Comment, error)
}

type Authorization interface{}

type Repository struct {
	PostsList
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PostsList:     NewPostListPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
