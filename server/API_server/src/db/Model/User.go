package model

type User struct {
	Id       string  `json:"id"       gorm:"type:varchar(255); NOT NULL; PRIMARY_KEY:true;"`
	Name     string  `json:"name"     gorm:"type:varchar(255); NOT NULL;"`
	Email    string  `json:"email"    gorm:"type:varchar(255); NOT NULL;"`
	Password string  `json:"password" gorm:"type:varchar(255); NOT NULL;"`
	Img      string  `json:"img"      gorm:"type:varchar(255); NOT NULL;"`
	Boards   []*Board `json:"boards"	  gorm:"many2many:user_boards;"`
}
