package repository

import (
	"orenotorero/db/Model"
)

type KanbanRepository interface {
	InsertKanban(boardId, position int, title string) error
	SelectByBoardId(boardId int) ([]model.Kanban, error)
	DeleteKanban(kanbanId int) error
	UpdateKanbanTitle(kanbanId int, newTitle string) error
	UpdatePosition(position []int) error
}
