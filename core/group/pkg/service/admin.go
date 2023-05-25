package service

import "group-app-service/pkg/repository"
import "group-app-service/pkg/model"

type AdminService struct {
	repo repository.Administration
}

func NewAdminService(repo repository.Administration) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (s *AdminService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *AdminService) UpdateUsers(users model.UpdateUsers) error {
	return s.repo.UpdateUsers(users)
}

func (s *AdminService) DeleteUsers() error {
	return s.repo.DeleteUsers()
}
