package handler

import (
	"fmt"
	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/service"
	"time"
)

type CardHandler struct {
	CardService service.CardService
}

func NewCardHandler(service service.CardService) CardHandler {
	return CardHandler{CardService: service}
}

func (handler *CardHandler) CreateNewCard(context *gin.Context) {
	var reqBody requestBody.CardCreate

	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.CreateCard(id, reqBody.Title, reqBody.KanbanId, reqBody.Position)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) ChangeCardTitle(context *gin.Context) {
	var reqBody requestBody.CardChangeTitle

	claims := ginJwt.ExtractClaims(context)
	userId, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	fmt.Println(reqBody)

	err = handler.CardService.ChangeCardTitle(userId, reqBody.Id, reqBody.Title)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusBadRequest)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) ChangeCardDeadline(context *gin.Context) {
	var reqBody requestBody.CardChangeDeadline

	claims := ginJwt.ExtractClaims(context)
	userId, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	var deadline time.Time
	deadline, err = time.Parse("2006-01-02 15:04:05", reqBody.Deadline)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.ChangeCardDeadline(userId, reqBody.Id, deadline)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusBadRequest)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) AddFile(context *gin.Context) {
	var reqBody requestBody.FileAdd

	claims := ginJwt.ExtractClaims(context)
	userId, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.AttachFile(userId, reqBody.Id, reqBody.FileData, reqBody.FileName)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusBadRequest)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) DeleteCard(context *gin.Context) {
	var reqBody requestBody.CardDelete

	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.DeleteCard(id, reqBody.Id)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusBadRequest)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) DeleteFile(context *gin.Context) {
	var reqBody requestBody.FileDelete

	claims := ginJwt.ExtractClaims(context)
	userId, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.DeleteFile(userId, reqBody.Id)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusBadRequest)
		return
	}

	context.Status(http.StatusOK)
}
