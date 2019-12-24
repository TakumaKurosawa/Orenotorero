package model

type File struct {
	Id     int    `json:"id"      gorm:"type:int AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"`
	CardId int    `json:"card_id" gorm:"type:int; NOT NULL"`
	Name   string `json:"name"    gorm:"type:varchar(255); NOT NULL"`
	URL    string `json:"url"     gorm:"type:varchar(255); NOT NULL"`
	files  []File `json:"" gorm:""`
}
