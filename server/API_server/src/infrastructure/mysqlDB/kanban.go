package mysqlDB

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
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

func (p *KanbanRepositoryImpliment) DeleteKanban(kanbanId int) error {

	// カンバン削除機能
	return nil
}

func (p *KanbanRepositoryImpliment) UpdateKanbanTitle(kanbanId int, newTitle string) error {

	// カンバンタイトル変更機能
	return nil
}

func (p *KanbanRepositoryImpliment) UpdatePosition(position []int) error {
	//Kanban & Cardのポジション変更機能
	return nil
}