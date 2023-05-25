package service

import (
	commsApp "comms-app-service/pkg/model"
	"comms-app-service/pkg/repository"
)

type Authorization interface {
	ParseToken(token string) (int, error)
}

type Comms interface {
	CreateComment(userId int, comment commsApp.Comment) (int, error)
	GetCommentsByPostId(postId int) ([]commsApp.Comment, error)
	DeleteCommentById(userId int, commentId int) error
	UpdateCommentById(userId int, commentId int, comment commsApp.UpdateCommentInput) error
	GetAllComments() ([]commsApp.Comment, error)
}

type Service struct {
	Authorization
	Comms
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Comms:         NewCommsService(repos.Comms),
		Authorization: NewAuthService(repos.Authorization),
	}
}
