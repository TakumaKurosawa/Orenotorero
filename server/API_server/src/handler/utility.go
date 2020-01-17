package handler

import (
	ginJwt "github.com/appleboy/gin-jwt/v2"
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
	var reqBody requestBody.EmailCheck

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	isExist, err := handler.UtilityService.CheckEmail(reqBody.Email)
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

func (handler *UtilityHandler) UpdatePosition(context *gin.Context) {
	var reqBody requestBody.UpdatePosition

	claims := ginJwt.ExtractClaims(context)
	userId, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	err = handler.UtilityService.UpdatePosition(userId, reqBody.Position)
	if err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}