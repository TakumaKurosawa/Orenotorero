package service

import (
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type KanbanService struct {
	KanbanRepository repository.KanbanRepository
}

func NewKanbanService(repository repository.KanbanRepository) KanbanService {
	return KanbanService{KanbanRepository: repository}
}

func (KanbanSvc *KanbanService) GetKanban(userId string, boardId int) []model.Kanban {
	// UserIdでボードへのアクセス権限があるかをチェックする
	return KanbanSvc.KanbanRepository.SelectByBoardId(boardId)
}

func (KanbanSvc *KanbanService) CreateNewKanban(token, title string, boardId, position int) error {
	// tokenでボードへのアクセス権限があるかをチェックする

	return KanbanSvc.KanbanRepository.InsertKanban(boardId, position, title)
}

func (KanbanSvc *KanbanService) DeleteKanban(kanbanId int, token string) error {
	// tokenでボードへのアクセス権限があるかをチェックする

	return KanbanSvc.KanbanRepository.DeleteKanban(kanbanId)
}

func (KanbanSvc *KanbanService) ChangeKanbanTitle(kanbanId int, token, title string) error {
	// tokenでボードへのアクセス権限があるかをチェックする

	return KanbanSvc.KanbanRepository.UpdateKanbanTitle(kanbanId, title)
}
