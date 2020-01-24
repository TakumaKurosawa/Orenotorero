package mysqlDB

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
	"orenotorero/utility"
	"time"
)

type CardRepositoryImpliment struct {
	DB *gorm.DB
}

func NewCardRepoImpl(DB *gorm.DB) repository.CardRepository {
	return &CardRepositoryImpliment{
		DB: DB,
	}
}

func (p *CardRepositoryImpliment) InsertCard(userId string, title string, kanbanId, position int) error {
	var kanban model.Kanban
	p.DB.Find(&kanban, kanbanId)
	if kanban.Id == 0 {
		return errors.New("挿入先のカンバンが存在しません")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, kanban.BoardId)

	if isMyBoard {
		// カード新規作成機能
		newCard := model.Card{
			KanbanId: kanbanId,
			Position: position,
			Title:    title,
			Describe: nil,
			DeadLine: nil,
		}

		p.DB.Create(&newCard)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}

func (p *CardRepositoryImpliment) UpdateCardTitle(id int, title string) error {
	// カードタイトル更新機能
	return nil
}

func (p *CardRepositoryImpliment) UpdateCardDeadLine(userId string, cardId int, deadline time.Time) error {
	// カード期限更新機能
	var card model.Card
	p.DB.Find(&card, cardId).Related(&card.Kanban)
	if card.Id == 0 {
		return errors.New("Cardが見つかりませんでした")
	}
	if card.Kanban.Id == 0 {
		return errors.New("Kanbanが見つかりませんでした")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, card.Kanban.BoardId)

	if isMyBoard {
		p.DB.Model(&card).Update("dead_line", deadline)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}

func (p *CardRepositoryImpliment) SelectAll(kanbanId int) ([]model.Card, error) {
	// カンバンに紐づくカード全件取得機能
	var cards []model.Card
	p.DB.Find(&cards)

	return cards, nil
}
