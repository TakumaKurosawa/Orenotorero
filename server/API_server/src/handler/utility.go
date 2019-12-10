package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orenotorero/handler/requestBody"
	"orenotorero/service"
)

type UtilityHandler struct {
	UtilityService service.UtilityService
}

func NewUtilityHandler(service service.UtilityService) UtilityHandler {
	return UtilityHandler{UtilityService: service}
}

func (handler *UtilityHandler) EmailCheck(context *gin.Context) {
	var token string
	var reqBody requestBody.EmailCheck

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	isExist, err := handler.UtilityService.CheckEmail(token, reqBody.Email)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"is_exist": isExist})
}

func (handler *UtilityHandler) FileUpload(context *gin.Context) {
	var token string
	var reqBody requestBody.FileUpload

	err := context.BindHeader(token)
	if err != nil {
		context.Error(err)
	}

	err = context.BindJSON(reqBody)
	if err != nil {
		context.Error(err)
	}

	s3Url, err := handler.UtilityService.FileUpload(token, reqBody.Img)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"url": s3Url})
}
