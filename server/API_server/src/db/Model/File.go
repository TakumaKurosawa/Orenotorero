package model

type File struct {
	Id     int    `json:"id"      gorm:"PRIMARY_KEY"`
	CardId int    `json:"card_id" gorm:"type:int; NOT NULL"`
	Name   string `json:"name"    gorm:"type:varchar(255); NOT NULL"`
	URL    string `json:"url"     gorm:"type:varchar(255); NOT NULL"`
	Card   Card   `gorm:"foreignkey:Id;association_foreignkey:CardId"`
}
