package mysqlDB

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/infrastructure/mysqlDB/"
	"orenotorero/repository"
	"orenotorero/utility"
)

type KanbanRepositoryImpliment struct {
	DB *gorm.DB
}

func NewKanbanRepoImpl(DB *gorm.DB) repository.KanbanRepository {
	return &KanbanRepositoryImpliment{
		DB: DB,
	}
}

func (p *KanbanRepositoryImpliment) InsertKanban(boardId, position int, title string) error {

	// カンバン作成機能
	return nil
}

func (p *KanbanRepositoryImpliment) SelectByBoardId(boardId int) ([]model.Kanban, error) {
	var kanban []model.Kanban

	// カンバン取得機能

	return kanban, nil
}

func (p *KanbanRepositoryImpliment) DeleteKanban(userId string, kanbanId int) error {
	// カンバン削除機能
	var kanban model.Kanban
	p.DB.Where("id = ?", kanbanId).Find(&kanban)
	if kanban.Id == 0 {
		return errors.New("カンバンが存在しません")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, kanban.BoardId)

	if isMyBoard {
		var cards []model.Card
		p.DB.Model(&kanban).Related(&cards)

		//該当データの削除
		p.DB.Delete(&kanban)
		//Kanbanに紐づくCardがあれば削除する
		if len(cards) == 0 {
			p.DB.Delete(&cards)
		}
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}

func (p *KanbanRepositoryImpliment) UpdateKanbanTitle(kanbanId int, newTitle string) error {

	// カンバンタイトル変更機能
	return nil
}

func (p *KanbanRepositoryImpliment) UpdatePosition(position []int) error {
	//Kanban & Cardのポジション変更機能
	return nil
}
