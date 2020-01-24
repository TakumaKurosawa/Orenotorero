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

func (KanbanSvc *KanbanService) GetKanban(token string, boardId int) ([]model.Kanban, error) {
	// tokenでボードへのアクセス権限があるかをチェックする

	return KanbanSvc.KanbanRepository.SelectByBoardId(boardId)
}

func (KanbanSvc *KanbanService) CreateNewKanban(userId, title string, boardId, position int) error {
	//UserがBoardを所有しているかを確認する
	return KanbanSvc.KanbanRepository.InsertKanban(userId, boardId, position, title)
}

func (KanbanSvc *KanbanService) DeleteKanban(kanbanId int, token string) error {
	// tokenでボードへのアクセス権限があるかをチェックする

	return KanbanSvc.KanbanRepository.DeleteKanban(kanbanId)
}

func (KanbanSvc *KanbanService) ChangeKanbanTitle(userId string, kanbanId int, title string) error {
	// User確認をする
	return KanbanSvc.KanbanRepository.UpdateKanbanTitle(kanbanId, title)
}
