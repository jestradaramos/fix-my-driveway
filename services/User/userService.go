package services

import (
	"fix-my-driveway/domain"
)

// UserService will hold the repo and its functions
type UserService struct {
	repo domain.Repo
}

// NewUserService creates the user service
func NewUserService(repo domain.Repo) UserService {
	return UserService{repo}
}

// AddUser adds a user to our system
func (s UserService) AddUser(u *domain.User) (*domain.User, error) {
	err := s.repo.AddUser(u)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateUser is for making any changes to the user information
func (s UserService) UpdateUser(u domain.User) (*domain.User, error) {
	return nil, nil
}
