package mysqlDB

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/handler/requestBody"
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

func (p *KanbanRepositoryImpliment) SelectByBoardId(userId string, boardId int) ([]model.Kanban, error) {
	// カンバン&カード取得
	var board model.Board
	var kanbans []model.Kanban
	p.DB.Where("id = ?", boardId).Find(&board)
	if board.Id == 0 {
		return nil, errors.New("Boardが見つかりません")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, boardId)

	if isMyBoard {
		p.DB.Model(&board).Preload("Cards", func(db *gorm.DB) *gorm.DB {
			return db.Order("position")
		}).Order("position").Preload("Cards.Files").Related(&kanbans)
		return kanbans, nil
	} else {
		return nil, errors.New("ボードへの権限がありません")
	}
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
		p.DB.Delete(&kanban)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}

func (p *KanbanRepositoryImpliment) UpdateKanbanTitle(userId string, kanbanId int, newTitle string) error {
	// カンバンタイトル変更機能
	var kanban model.Kanban
	p.DB.Where("id = ?", kanbanId).Find(&kanban)
	if kanban.Id == 0 {
		return errors.New("カンバンが見つかりません")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, kanban.BoardId)

	if isMyBoard {
		p.DB.Model(&kanban).Update("title", newTitle)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}

func (p *KanbanRepositoryImpliment) UpdatePosition(userId string, positions []requestBody.UpdatePosition) error {
	//Kanban & Cardのポジション変更機能
	var board model.Board
	p.DB.Find(&model.Kanban{}, positions[0].KanbanId).Related(&board)
	if board.Id == 0 {
		return errors.New("Boardが見つかりませんでした")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, board.Id)

	if isMyBoard {
		var kanbans []model.Kanban
		//対象のBoardにあるカンバン以下全てをしゅとくしてくる
		p.DB.Model(&board).Preload("Cards.Files").Related(&kanbans)

		for num, _ := range positions {
			//Kanban Positionの更新
			kanbans[positions[num].KanbanId-1].Id = num

			////Card Positionの更新
			//for i,_ := range positions[num].CardArray {
			//	kanbans[num].Cards[i].Position = i
			//}
		}
		//最後に変更分をぜんぶ一斉にてきおうさせる
		p.DB.Update(&kanbans)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
	return nil
}
