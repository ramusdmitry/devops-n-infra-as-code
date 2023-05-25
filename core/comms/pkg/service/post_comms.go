package service

import (
	commsApp "comms-app-service/pkg/model"
	"comms-app-service/pkg/repository"
)

type CommsService struct {
	repo repository.Comms
}

func (s *CommsService) CreateComment(userId int, comment commsApp.Comment) (int, error) {
	return s.repo.CreateComment(userId, comment)
}

func (s *CommsService) GetCommentsByPostId(postId int) ([]commsApp.Comment, error) {
	return s.repo.GetCommentsByPostId(postId)
}

func (s *CommsService) GetAllComments() ([]commsApp.Comment, error) {
	return s.repo.GetAllComments()
}

func (s *CommsService) DeleteCommentById(userId int, commentId int) error {
	return s.repo.DeleteCommentById(userId, commentId)
}

func (s *CommsService) UpdateCommentById(userId int, commentId int, comment commsApp.UpdateCommentInput) error {

	if err := comment.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateCommentById(userId, commentId, comment)
}

func NewCommsService(repo repository.Comms) *CommsService {
	return &CommsService{
		repo: repo,
	}
}
