package service

import (
	postApp "posts-app-service/pkg/model"
	"posts-app-service/pkg/repository"
)

type PostService struct {
	repo repository.PostsList
}

func (s *PostService) CreatePost(userId int, post postApp.Post) (int, error) {
	return s.repo.CreatePost(userId, post)
}

func (s *PostService) GetCommentsInPost(postId int) ([]postApp.Comment, error) {
	return s.repo.GetAllComments(postId)
}

func (s *PostService) GetAllPosts() ([]postApp.Post, error) {

	postsArray, err := s.repo.GetAllPosts()

	tempArr := &postsArray

	if err != nil {
		return nil, err
	}

	for i, post := range *tempArr {

		comments, err := s.GetCommentsInPost(post.Id)
		if err != nil {
			return nil, err
		}

		var commentsArray postApp.Comments
		commentsArray.Data = comments
		post.Comments = commentsArray
		(*tempArr)[i] = post

	}
	return postsArray, err
}

func (s *PostService) DeletePostById(userId, postId int) error {
	return s.repo.DeletePostById(userId, postId)
}

func (s *PostService) UpdatePostById(userId, postId int, post postApp.UpdatePostInput) error {

	if err := post.Validate(); err != nil {
		return err
	}

	return s.repo.UpdatePostById(userId, postId, post)
}

func NewPostService(repo repository.PostsList) *PostService {
	return &PostService{
		repo: repo,
	}
}
