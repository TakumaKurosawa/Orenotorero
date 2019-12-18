package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{UserService: service}
}

func (handler *UserHandler) Login(context *gin.Context) {
	var reqBody requestBody.UserLogin

	err := context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	token, err := handler.UserService.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func (handler *UserHandler) GetUser(context *gin.Context) {
	var token string

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	user, err := handler.UserService.GetUser(token)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, user)
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
