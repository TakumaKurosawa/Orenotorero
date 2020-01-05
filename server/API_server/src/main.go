package main

import (
	"github.com/gin-gonic/gin"
	"orenotorero/db"
	"orenotorero/middleware"
)

func main() {
	dbInstance := db.GormCreate()
	defer dbInstance.Close()

	userAPI := InitUserAPI(dbInstance)
	utilityAPI := InitUtilityAPI(dbInstance)
	boardAPI := InitBoardAPI(dbInstance)
	cardAPI := InitCardAPI(dbInstance)
	kanbanAPI := InitKanbanAPI(dbInstance)

	jwtAuth, err := middleware.CreateJwtInstance(userAPI)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// connection testAPI
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// userAPI
	r.POST("/login", jwtAuth.LoginHandler)
	authByJwt := r.Group("")
	authByJwt.Use(jwtAuth.MiddlewareFunc())
	{
		authByJwt.GET("/user/get", userAPI.GetUser)
	}
	r.POST("/user/create", userAPI.CreateNewUser)
	r.GET("/users", userAPI.GetAllUsers)

	// boardAPI
	authByJwt.Use(jwtAuth.MiddlewareFunc())
	{
		authByJwt.GET("/board", boardAPI.GetBoard)
	}
	r.POST("/board", boardAPI.CreateNewBoard)
	r.PUT("/board/publish", boardAPI.ChangeBoardPublish)
	r.POST("/board/invite", boardAPI.SendInviteMail)

	// kanbanAPI
	r.GET("/kanban", kanbanAPI.GetKanban)
	r.POST("/kanban", kanbanAPI.CreateNewKanban)
	r.DELETE("/kanban", kanbanAPI.DeleteKanban)
	r.PUT("/kanban", kanbanAPI.ChangeKanbanTitle)
	r.PUT("/kanban/position", kanbanAPI.ChangeKanbanPosition)

	// cardAPI
	r.POST("/card", cardAPI.CreateNewCard)
	r.POST("/card/file", cardAPI.AddFile)
	r.PUT("/card", cardAPI.ChangeCardTitle)
	r.PUT("/card/deadline", cardAPI.ChangeCardDeadline)
	r.PUT("/card/position", cardAPI.ChangeCardPosition)

	// utilityAPI
	r.PUT("/email/check", utilityAPI.EmailCheck)
	r.PUT("/img", utilityAPI.FileUpload)

	// ポートを設定しています。
	r.Run(":3000")
}
