package model

import "time"

type Card struct {
	Id       int       `json:"id"        gorm:"PRIMARY_KEY"`
	KanbanId int       `json:"kanban_id" gorm:"type:int; NOT NULL"`
	Position int       `json:"position"  gorm:"type:int; NOT NULL"`
	Title    string    `json:"title"     gorm:"type:varchar(255); NOT NULL"`
	Describe string    `json:"describe"  gorm:"type:varchar(255)"`
	DeadLine time.Time `json:"dead_line" gorm:"type:datetime;"`
	Files    []File    `json:"files"     gorm:"foreignkey:CardId"`
}
