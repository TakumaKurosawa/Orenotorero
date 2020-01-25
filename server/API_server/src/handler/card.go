package handler

import (
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
	var token, s3Url string
	var reqBody requestBody.CardAddFile

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.CardService.InsertFileData(reqBody.Id, token, s3Url, reqBody.FileName)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}
