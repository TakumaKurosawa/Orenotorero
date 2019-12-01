package main

import (
	"fmt"
	"orenotorero/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.GormConnect()
	defer db.Close()

	fmt.Println(db)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	// ポートを設定しています。
	r.Run(":3000")
}
