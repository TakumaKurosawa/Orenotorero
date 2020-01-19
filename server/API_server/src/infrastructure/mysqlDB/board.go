package mysqlDB

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
	"orenotorero/utility"
	"time"
)

type BoardRepositoryImpliment struct {
	DB *gorm.DB
}

func NewBoardRepoImpl(DB *gorm.DB) repository.BoardRepository {
	return &BoardRepositoryImpliment{
		DB: DB,
	}
}

func (p *BoardRepositoryImpliment) SelectByUserId(userId string) []model.Board {
	//ユーザIDにひもづくボードを全件取得する

	//モデル変数の宣言
	var boards []model.Board
	var user model.User

	//userIdによるユーザー情報の取得
	p.DB.Where("id=?", userId).Find(&user)

	//取得したユーザーに対応するBoardを取得
	p.DB.Model(&user).Related(&boards, "Boards")

	return boards
}

func (p *BoardRepositoryImpliment) InsertBoard(userId, title, img string) error {
	var user model.User
	user.Id = userId

	// ボード新規作成
	newBoard := model.Board{
		CreatedUserId: userId,
		Title:         title,
		Publish:       false,
		BgImg:         img,
		LastAccess:    time.Now(),
		InviteURL:     "https://orenotorero/invite/1", // TODO:招待URLを捌くチケットで修正
	}

	p.DB.Create(&newBoard)
	p.DB.Model(&user).Association("Boards").Append(&newBoard)
	return nil
}

func (p *BoardRepositoryImpliment) UpdateBoardPublish(userId string, boardId int, publish bool) error {
	// ボードの公開・非公開情報更新
	var board model.Board

	isMyBoard := utility.IsMyBoard(p.DB, userId, boardId)

	if isMyBoard {
		board.Id = boardId
		p.DB.Model(&board).Update("publish", publish)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}
