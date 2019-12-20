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

func (userSvc *UserService) CreateNewUser(id, name, email, password string) error {

	user := model.User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return userSvc.UserRepository.InsertUser(user)

}

func (userSvc *UserService) SelectAll() ([]model.User, error) {
	return userSvc.UserRepository.SelectAll()
}
