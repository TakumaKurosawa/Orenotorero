package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"orenotorero/db/Model"
	"os"
)

func GormCreate() *gorm.DB {
	ip := os.Getenv("ORENOTORERO_DB_SERVICE_HOST")
	if ip == "" {
		ip = "127.0.0.1"
	}

	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(" + ip + ":30002)"
	DBNAME := "orenotorero"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	db = migration(db)

	fmt.Printf("db connected: %+v\n", db)
	return db
}

func migration(db *gorm.DB) *gorm.DB {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Board{})
	db.AutoMigrate(&model.Kanban{}).AddForeignKey("board_id", "board(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&model.Card{}).AddForeignKey("kanban_id", "kanban(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&model.File{}).AddForeignKey("card_id", "card(id)", "CASCADE", "CASCADE")

	return db
}
