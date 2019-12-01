package model

type Kanban struct {
	Id       int    `json:"id"       gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	BoardId  int    `json:"board_id" gorm:"type:int;NOT NULL"`
	Position int    `json:"position" gorm:"type:int;NOT NULL"`
	Title    string `json:"title"    gorm:"type:varchar(255);NOT NULL"`
	Board    Board  `json:"board"    gorm:"foreignkey:BoardId"`
}
