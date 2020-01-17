package repository

import (
	"orenotorero/db/Model"
	"time"
)

type CardRepository interface {
	InsertCard(title string, kanbanId, position int) error
	UpdateCardTitle(id int, title string) error
	UpdateCardDeadLine(userId string, cardId int, deadline time.Time) error
	SelectAll(kanbanId int) ([]model.Card, error)
}
