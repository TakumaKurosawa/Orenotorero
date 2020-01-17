package utility

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
)

func IsMyBoard(dbInstance *gorm.DB, userId string, boardId int) bool {
	//モデル変数の宣言
	var user model.User
	var board model.Board

	dbInstance.Where("id=?", userId).Find(&user)
	dbInstance.Model(&user).Where("board_id=?", boardId).Related(&board, "Boards")

	if board.Id != 0 {
		return true
	} else {
		return false
	}
}
