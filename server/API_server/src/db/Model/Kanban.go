package model

type Kanban struct {
	Id       int    `json:"id"       gorm:"type:int AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"`
	Position int    `json:"position" gorm:"type:int; NOT NULL"`
	Title    string `json:"title"    gorm:"type:varchar(255); NOT NULL"`
	cards    []Card `json:"" gorm:""`
}
