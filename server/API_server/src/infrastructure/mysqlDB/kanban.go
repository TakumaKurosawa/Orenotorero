package mysqlDB

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
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

func (p *KanbanRepositoryImpliment) InsertKanban(userId string, boardId, position int, title string) error {
	var board model.Board
	p.DB.Find(&board, boardId)
	if board.Id == 0 {
		return errors.New("挿入先のボードが存在しません")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, boardId)

	if isMyBoard {
		// カンバン作成
		newKanban := model.Kanban{
			BoardId:  boardId,
			Position: position,
			Title:    title,
		}

		p.DB.Create(&newKanban)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}

func (p *KanbanRepositoryImpliment) SelectByBoardId(boardId int) ([]model.Kanban, error) {
	var kanban []model.Kanban

	// カンバン取得機能

	return kanban, nil
}

func (p *KanbanRepositoryImpliment) DeleteKanban(kanbanId int) error {

	// カンバン削除機能
	return nil
}

func (p *KanbanRepositoryImpliment) UpdateKanbanTitle(kanbanId int, newTitle string) error {
	// カンバンタイトル変更機能
	p.DB.Find(&model.Kanban{}, kanbanId).Update("Title", newTitle)
	return nil
}

func (p *KanbanRepositoryImpliment) UpdatePosition(position []int) error {
	//Kanban & Cardのポジション変更機能

	return nil
}
