package repository

import (
	"orenotorero/db/Model"
	"orenotorero/handler/requestBody"
)

type KanbanRepository interface {
	InsertKanban(userId string, boardId, position int, title string) error
	SelectByBoardId(userId string, boardId int) ([]model.Kanban, error)
	DeleteKanban(userId string, kanbanId int) error
	UpdateKanbanTitle(userId string, kanbanId int, newTitle string) error
	UpdatePosition(userId string, position []requestBody.UpdatePosition) error
}
