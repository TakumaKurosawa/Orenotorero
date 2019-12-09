package repository

import (
	"orenotorero/db/Model"
)

type UserRepository interface {
	InsertUser() error
	Login() (string, error)
	UpdateImg(s3Url string) error
	SelectByUserId() (model.User, error)
	SelectAll() ([]model.User, error)
	SelectByEmail(email string) (model.User, error)
}
