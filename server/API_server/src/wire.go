//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"orenotorero/handler"
	"orenotorero/infrastructure/mysqlDB"
	"orenotorero/service"
)

func InitUserAPI(db *gorm.DB) handler.UserHandler {
	wire.Build(mysqlDB.NewUserRepoImpl, service.NewUserService, handler.NewUserHandler)

	return handler.UserHandler{}
}
