package mysqlDB

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
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
	// ユーザIDにひもづくボードを全件取得
	//var authors []model.Author
	//p.DB.Where("user_id = ?", userId).Find(&authors)
	//
	//var boardIds []int
	//for _, author := range authors {
	//	boardIds = append(boardIds, author.BoardId)
	//}
	//
	//var Board []model.Board
	//p.DB.Where("id IN (?)", boardIds).Find(&Board)

	var authors []model.Author
	var Boards []model.Board

	p.DB.Where("user_id = ?", userId).Find(&authors).Related(&Boards)
	return Board
}

func (p *BoardRepositoryImpliment) InsertBoard(userId int, title, img string) error {
	// ボード新規作成
	return nil
}

func (p *BoardRepositoryImpliment) UpdateBoardPublish(id int, publish bool) error {
	// ボードの公開・非公開情報更新
	return nil
}
