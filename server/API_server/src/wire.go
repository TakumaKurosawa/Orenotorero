package wire

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"orenotorero/handler"
	"orenotorero/repository"
	"orenotorero/service"
)

func InitUserAPI(db *gorm.DB) handler.UserHandler {
	wire.Build(repository.NewUserRepostiory, service.NewUserService, handler.NewUserHandler)

	return handler.UserHandler{}
}
