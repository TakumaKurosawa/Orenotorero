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

func (p *BoardRepositoryImpliment) SelectByUserId(userId int) ([]model.Board, error) {
	// ユーザIDにひもづくボードを全件取得
	return nil, nil
}

func (p *BoardRepositoryImpliment) InsertBoard(userId int, title, img string) error {
	// ボード新規作成
	return nil
}

func (p *BoardRepositoryImpliment) UpdateBoardPublish(id int, publish bool) error {
	// ボードの公開・非公開情報更新
	return nil
}
