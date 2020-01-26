package main

import (
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

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
		authByJwt.PUT("/board/publish", boardAPI.ChangeBoardPublish)
		authByJwt.POST("/board", boardAPI.CreateNewBoard)
		authByJwt.POST("/board/invite", boardAPI.SendInviteMail)
	}

	// kanbanAPI
	authByJwt.Use(jwtAuth.MiddlewareFunc())
	{
		authByJwt.GET("/kanban", kanbanAPI.GetKanban)
		authByJwt.POST("/kanban", kanbanAPI.CreateNewKanban)
		authByJwt.DELETE("/kanban", kanbanAPI.DeleteKanban)
		authByJwt.PUT("/kanban", kanbanAPI.ChangeKanbanTitle)
	}

	// cardAPI
	authByJwt.Use(jwtAuth.MiddlewareFunc())
	{
		authByJwt.POST("/card", cardAPI.CreateNewCard)
		authByJwt.POST("/card/file", cardAPI.AddFile)
		authByJwt.DELETE("/card/file", cardAPI.DeleteFile)
		authByJwt.PUT("/card", cardAPI.ChangeCardTitle)
		authByJwt.PUT("/card/deadline", cardAPI.ChangeCardDeadline)
		authByJwt.DELETE("/card", cardAPI.DeleteCard)
	}

	// utilityAPI
	r.GET("/email/check", utilityAPI.EmailCheck)
	authByJwt.Use(jwtAuth.MiddlewareFunc())
	{
		authByJwt.POST("/img", utilityAPI.FileUpload)
		authByJwt.PUT("/position", utilityAPI.UpdatePosition)
	}

	// ポートを設定しています。
	r.Run(":8080")
}
