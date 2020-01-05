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

	db.SingularTable(true)
	db.AutoMigrate(
		&model.User{},
		&model.Board{},
		&model.Author{},
		&model.Card{},
		&model.Kanban{},
		&model.File{},
	)

	return db
}
