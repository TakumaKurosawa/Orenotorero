package repository

import (
	"orenotorero/db/Model"
	"time"
)

type CardRepository interface {
	InsertCard(userId string, title string, kanbanId, position int) error
	UpdateCardDeadLine(userId string, cardId int, deadline time.Time) error
	UpdateCardTitle(userId string, id int, title string) error
	UpdateCardDescribe(userId string, id int, describe string) error
	SelectAll(kanbanId int) ([]model.Card, error)
	DeleteCard(userId string, cardId int) error
}
