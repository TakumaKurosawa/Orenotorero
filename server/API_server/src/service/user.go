package service

import (
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return UserService{UserRepository: repository}
}

func (userSvc *UserService) SelectAll() []model.User {
	return userSvc.UserRepository.SelectAll()
}
