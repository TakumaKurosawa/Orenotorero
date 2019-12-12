package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"orenotorero/db/Model"
)

func GormCreate() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(192.168.33.10:3333)"
	DBNAME := "orenotorero"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Printf("hoge")
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
