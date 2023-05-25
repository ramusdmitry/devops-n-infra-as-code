package service

import (
	postApp "posts-app-service/pkg/model"
	"posts-app-service/pkg/repository"
)

type Authorization interface {
	ParseToken(token string) (int, error)
}

type PostsList interface {
	GetAllPosts() ([]postApp.Post, error)
	CreatePost(userId int, post postApp.Post) (int, error)
	DeletePostById(userId, postId int) error
	UpdatePostById(userId, postId int, post postApp.UpdatePostInput) error
}

type Service struct {
	Authorization
	PostsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		PostsList:     NewPostService(repos.PostsList),
		Authorization: NewAuthService(repos.Authorization),
	}
}
