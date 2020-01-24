package handler

import (
	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/handler/requestHeader"
	"orenotorero/service"
)

type KanbanHandler struct {
	KanbanService service.KanbanService
}

func NewKanbanHandler(service service.KanbanService) KanbanHandler {
	return KanbanHandler{KanbanService: service}
}

func (handler *KanbanHandler) GetKanban(context *gin.Context) {
	var reqHeader requestHeader.KanbanGet

	err := context.BindHeader(reqHeader)
	if err != nil {
		context.Error(err)
	}

	kanbans, err := handler.KanbanService.GetKanban(reqHeader.Token, reqHeader.BoardId)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, kanbans)
}

func (handler *KanbanHandler) CreateNewKanban(context *gin.Context) {
	var reqBody requestBody.KanbanCreate

	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.KanbanService.CreateNewKanban(id, reqBody.Title, reqBody.BoardId, reqBody.Position)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *KanbanHandler) DeleteKanban(context *gin.Context) {
	var reqBody requestBody.KanbanDelete

	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.KanbanService.DeleteKanban(id, reqBody.KanbanId)
	if err != nil {
		context.Error(err)
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}

func (handler *KanbanHandler) ChangeKanbanTitle(context *gin.Context) {
	var token string
	var reqBody requestBody.KanbanChangeTitle

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.KanbanService.ChangeKanbanTitle(reqBody.KanbanId, token, reqBody.Title)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}
