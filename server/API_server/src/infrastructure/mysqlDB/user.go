package mysqlDB

import (
	"github.com/appleboy/gin-jwt/v2"
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

func (p *UserRepositoryImpliment) Login(email, password string) (*model.User, error) {
	user, err := p.SelectByEmail(email)
	if err != nil {
		return nil, err
	}

	// ユーザ認証機能
	if user.Password != password {
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}

func (p *UserRepositoryImpliment) UpdateImg(s3Url string) error {
	// s3Urlをユーザのimgカラムに挿入・更新
	return nil
}

func (p *UserRepositoryImpliment) SelectByUserId(id string) (model.User, error) {
	var user model.User

	// ユーザ取得機能
	p.DB.Where("id = ?", id).Find(&user)

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
	var user model.User
	p.DB.Where("email = ?", email).Find(&user)
	// DB.Whereを使って受け取ったEmailがデータベースに存在するかチェックする
	return user, nil
}
