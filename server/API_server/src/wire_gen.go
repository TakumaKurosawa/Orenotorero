// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"orenotorero/handler"
	"orenotorero/infrastructure/mysqlDB"
	"orenotorero/service"
)

// Injectors from wire.go:

func InitUserAPI(db *gorm.DB) handler.UserHandler {
	userRepository := mysqlDB.NewUserRepoImpl(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}

func InitUtilityAPI(db *gorm.DB) handler.UtilityHandler {
	userRepository := mysqlDB.NewUserRepoImpl(db)
	kanbanRepository := mysqlDB.NewKanbanRepoImpl(db)
	utilityService := service.NewUtilityService(userRepository, kanbanRepository)
	utilityHandler := handler.NewUtilityHandler(utilityService)
	return utilityHandler
}

func InitBoardAPI(db *gorm.DB) handler.BoardHandler {
	boardRepository := mysqlDB.NewBoardRepoImpl(db)
	boardService := service.NewBoardService(boardRepository)
	boardHandler := handler.NewBoardHandler(boardService)
	return boardHandler
}

func InitCardAPI(db *gorm.DB) handler.CardHandler {
	cardRepository := mysqlDB.NewCardRepoImpl(db)
	fileRepository := mysqlDB.NewFileRepoImpl(db)
	cardService := service.NewCardService(cardRepository, fileRepository)
	cardHandler := handler.NewCardHandler(cardService)
	return cardHandler
}

func InitKanbanAPI(db *gorm.DB) handler.KanbanHandler {
	kanbanRepository := mysqlDB.NewKanbanRepoImpl(db)
	kanbanService := service.NewKanbanService(kanbanRepository)
	kanbanHandler := handler.NewKanbanHandler(kanbanService)
	return kanbanHandler
}
