package handler

import (
	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/service"
	"orenotorero/utility"
	"os"
	"time"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{UserService: service}
}

func (handler *UserHandler) Login(context *gin.Context) (interface{}, error) {
	var reqBody requestBody.UserLogin

	err := context.BindJSON(&reqBody)
	if err != nil {
		return "", ginJwt.ErrMissingLoginValues
	}

	user, err := handler.UserService.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, ginJwt.ErrFailedAuthentication
	}

	return user, nil
}

func (handler *UserHandler) GetUser(context *gin.Context) {
	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	user, err := handler.UserService.GetUser(id)
	if err != nil {
		context.Error(err)
	}

	context.JSON(200, user)
}

func (handler *UserHandler) CreateNewUser(context *gin.Context) {
	var reqBody requestBody.UserCreate

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	token, err := handler.UserService.CreateNewUser(reqBody.Name, reqBody.Email, reqBody.Password)
	if err != nil {
		context.Error(err)
	}

	token = "hoge"
	context.JSON(http.StatusOK, gin.H{"token": token})
}

func (handler *UserHandler) GetAllUsers(context *gin.Context) {
	users, err := handler.UserService.SelectAll()
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"testUsers": users})
}
