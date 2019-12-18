package mysqlDB

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type UserRepositoryImpliment struct {
	DB *gorm.DB
}

func NewUserRepoImpl(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpliment{
		DB: DB,
	}
}

func (p *UserRepositoryImpliment) Login() (string, error) {
	var token string

	// ユーザ認証機能

	return token, nil
}

func (p *UserRepositoryImpliment) UpdateImg(s3Url string) error {
	// s3Urlをユーザのimgカラムに挿入・更新
	return nil
}

func (p *UserRepositoryImpliment) SelectByUserId() (model.User, error) {
	var user model.User

	// ユーザ取得機能

	return user, nil
}

func (p *UserRepositoryImpliment) InsertUser(user model.User) error {
	// ユーザ作成機能
	p.DB.Create(&user)

	return nil
}

func (p *UserRepositoryImpliment) SelectAll() ([]model.User, error) {
	var users []model.User
	p.DB.Find(&users)

	return users, nil
}

func (p *UserRepositoryImpliment) SelectByEmail(email string) (model.User, error) {
	// DB.Whereを使って受け取ったEmailがデータベースに存在するかチェックする
	return model.User{}, nil
}
