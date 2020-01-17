package handler

import (
	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/service"
)

type CardHandler struct {
	CardService service.CardService
}

func NewCardHandler(service service.CardService) CardHandler {
	return CardHandler{CardService: service}
}

func (handler *CardHandler) CreateNewCard(context *gin.Context) {
	var token string
	var reqBody requestBody.CardCreate

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.CreateCard(token, reqBody.Title, reqBody.KanbanId, reqBody.Position)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) ChangeCardTitle(context *gin.Context) {
	var token string
	var reqBody requestBody.CardChangeTitle

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.ChangeCardTitle(reqBody.Id, token, reqBody.Title)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) ChangeCardDeadline(context *gin.Context) {
	var token string
	var reqBody requestBody.CardChangeDeadline

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.ChangeCardDeadline(reqBody.Id, token, reqBody.Deadline)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (handler *CardHandler) AddFile(context *gin.Context) {
	var s3Url string
	var reqBody requestBody.CardAddFile

	claims := ginJwt.ExtractClaims(context)
	userId, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}
	err = handler.CardService.InsertFileData(userId, reqBody.Id, s3Url, reqBody.FileName)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}
