package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{UserService: service}
}

func (handler *UserHandler) SelectAll(context *gin.Context) {
	users := handler.UserService.SelectAll()

	context.JSON(http.StatusOK, gin.H{"testUsers": users})
}
