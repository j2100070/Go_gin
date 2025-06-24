package services

import (
	"Go_gin/dto"
	"Go_gin/models"
	"Go_gin/repositories"
)

type IUserService interface {
	Create(createUserInput dto.CreateUserInput) (*models.User, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

func (s *UserService) Create(createUserInput dto.CreateUserInput) (*models.User, error) {
	newUser := models.User{
		Username: createUserInput.Username,
		Email:    createUserInput.Email,
		Password: createUserInput.Password,
	}
	return s.repository.Create(newUser)
}
