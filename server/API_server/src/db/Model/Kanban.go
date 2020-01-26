package model

type Kanban struct {
	Id       int    `json:"id"       gorm:"PRIMARY_KEY"`
	BoardId  int    `json:"board_id" gorm:"type:int; NOT NULL"`
	Position int    `json:"position" gorm:"type:int; NOT NULL"`
	Title    string `json:"title"    gorm:"type:varchar(255); NOT NULL"`
	Cards    []Card `json:"card"	 gorm:"foreignkey:KanbanId"`
	Board    Board  `json:"-"        gorm:"foreignkey:Id;association_foreignkey:BoardId"`
}
