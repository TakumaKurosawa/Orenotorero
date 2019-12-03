package model

type Files struct {
	Id     int    `json:"id"      gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	CardId int    `json:"card_id" gorm:"type:int;NOT NULL"`
	Name   string `json:"name"    gorm:"type:varchar(255);NOT NULL"`
	URL    string `json:"url"     gorm:"type:varchar(255);NOT NULL"`
	Card   Card   `json:"card"    gorm:"foreignkey:CardId"`
}
