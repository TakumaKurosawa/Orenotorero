package main

import (
	"github.com/gin-gonic/gin"
	"orenotorero/db"
)

func main() {
	dbInstance := db.GormCreate()
	defer dbInstance.Close()

	//dbInstance.Create(&user)
	//dbInstance.Find(&user)
	//fmt.Println()
	//fmt.Printf("%v\n", user)
	//fmt.Println()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	// ポートを設定しています。
	r.Run(":3000")
}
