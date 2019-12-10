package repository

import (
	"orenotorero/db/Model"
)

type BoardRepository interface {
	SelectByUserId(userId int) ([]model.Board, error)
	InsertBoard(userId int, title, img string) error
	UpdateBoardPublish(id int, publish bool) error
}
