package services

import (
	. "fix-my-driveway/domain"
)

// UserService will hold the repo and its functions
type UserService struct {
	repo Repo
}

// NewUserService creates the user service
func NewUserService(repo Repo) UserService {
	return UserService{repo}
}

// AddUser adds a user to our system
func (s UserService) AddUser(u User) (*User, error) {
	user, err := s.repo.AddUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser is for making any changes to the user information
func (s UserService) UpdateUser(u User) (*User, error) {
	return nil, nil
}
