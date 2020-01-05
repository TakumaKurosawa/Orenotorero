package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/service"
)

type BoardHandler struct {
	BoardService service.BoardService
	service.UserService
}

func NewBoardHandler(service service.BoardService) BoardHandler {
	return BoardHandler{BoardService: service}
}

func (handler *BoardHandler) ChangeBoardPublish(context *gin.Context) {
	var token string
	var reqBody requestBody.BoardChangePublish

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.BoardService.ChangePublishInfo(token, reqBody.Id, reqBody.Publish)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (handler *BoardHandler) SendInviteMail(context *gin.Context) {
	var token string
	var reqBody requestBody.BoardSendInviteMail

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.BoardService.SendInviteMail(reqBody.Id, token, reqBody.Email, reqBody.Message)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (handler *BoardHandler) CreateNewBoard(context *gin.Context) {
	var token string
	var reqBody requestBody.BoardCreate

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.BoardService.CreateNewBoard(token, reqBody.Title, reqBody.Img)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (handler *BoardHandler) GetBoard(context *gin.Context) {

	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}
	Boards := handler.BoardService.GetBoard(id)

	context.JSON(http.StatusOK, Boards)

}