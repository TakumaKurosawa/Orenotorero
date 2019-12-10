package main

import (
	"github.com/gin-gonic/gin"
	"orenotorero/db"
)

func main() {
	dbInstance := db.GormCreate()
	defer dbInstance.Close()

	userAPI := InitUserAPI(dbInstance)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// userAPI
	r.POST("/login", userAPI.Login)
	r.POST("/user/create", userAPI.CreateNewUser)
	r.GET("/user/get", userAPI.GetUser)
	r.GET("/users", userAPI.GetAllUsers)

	// boardAPI

	// kanbanAPI

	// cardAPI

	// utilityAPI


	// ポートを設定しています。
	r.Run(":3000")
}
