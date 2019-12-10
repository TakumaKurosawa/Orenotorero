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

func (userSvc *UserService) Login(email, password string) (string, error) {
	return userSvc.UserRepository.Login()
}

func (userSvc *UserService) GetUser(token string) (model.User, error) {
	return userSvc.UserRepository.SelectByUserId()
}

func (userSvc *UserService) CreateNewUser(name, email, password string) (string, error) {
	var token string

	err := userSvc.UserRepository.InsertUser()

	// token作成処理
	return token, err
}

func (userSvc *UserService) SelectAll() ([]model.User, error) {
	return userSvc.UserRepository.SelectAll()
}