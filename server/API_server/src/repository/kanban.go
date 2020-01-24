package repository

import (
	"orenotorero/db/Model"
)

type KanbanRepository interface {
	InsertKanban(userId string, boardId, position int, title string) error
	SelectByBoardId(boardId int) ([]model.Kanban, error)
	DeleteKanban(kanbanId int) error
	UpdateKanbanTitle(userId string, kanbanId int, newTitle string) error
	UpdatePosition(position []int) error
}
