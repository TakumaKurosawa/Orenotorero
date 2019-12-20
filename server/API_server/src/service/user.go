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

func (userSvc *UserService) Login(email, password string) (*model.User, error) {
	return userSvc.UserRepository.Login(email, password)
}

func (userSvc *UserService) GetUser(id string) (model.User, error) {
	return userSvc.UserRepository.SelectByUserId(id)
}

func (userSvc *UserService) CreateNewUser(name, email, password string) (string, error) {
	var token string

	user := model.User{
		Name: name,
		Email: email,
		Password: password,
	}


	err := userSvc.UserRepository.InsertUser(user)

	// token作成処理
	return token, err
}

func (userSvc *UserService) SelectAll() ([]model.User, error) {
	return userSvc.UserRepository.SelectAll()
}
