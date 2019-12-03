package repository

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepostiory(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (p *UserRepository) SelectAll() []model.User {
	var users []model.User
	p.DB.Find(&users)

	return users
}
