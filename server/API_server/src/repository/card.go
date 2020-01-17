package repository

import (
	"orenotorero/db/Model"
	"time"
)

type CardRepository interface {
	InsertCard(title string, kanbanId, position int) error
	UpdateCardTitle(cardId int, title string) error
	UpdateCardDeadLine(id int, deadline time.Time) error
	SelectAll(kanbanId int) ([]model.Card, error)
}
