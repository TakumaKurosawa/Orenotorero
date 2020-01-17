package mysqlDB

import (
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
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

func (p *CardRepositoryImpliment) InsertCard(title string, kanbanId, position int) error {
	// カード新規作成機能
	return nil
}

func (p *CardRepositoryImpliment) UpdateCardTitle(id int, title string) error {
	// カードタイトル更新機能
	return nil
}

func (p *CardRepositoryImpliment) UpdateCardDeadLine(id int, deadline time.Time) error {
	// カード期限更新機能
	return nil
}

func (p *CardRepositoryImpliment) SelectAll(kanbanId int) ([]model.Card, error) {
	// カンバンに紐づくカード全件取得機能
	var cards []model.Card
	p.DB.Find(&cards)

	return cards, nil
}
