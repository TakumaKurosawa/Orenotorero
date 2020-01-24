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

func (KanbanSvc *KanbanService) GetKanban(userId string, boardId int) ([]model.Kanban, error) {
	return KanbanSvc.KanbanRepository.SelectByBoardId(userId, boardId)
}

func (KanbanSvc *KanbanService) CreateNewKanban(userId, title string, boardId, position int) error {
	//UserがBoardを所有しているかを確認する
	return KanbanSvc.KanbanRepository.InsertKanban(userId, boardId, position, title)
}

func (KanbanSvc *KanbanService) DeleteKanban(userId string, kanbanId int) error {
	return KanbanSvc.KanbanRepository.DeleteKanban(userId, kanbanId)
}

func (KanbanSvc *KanbanService) ChangeKanbanTitle(userId string, kanbanId int, title string) error {
	return KanbanSvc.KanbanRepository.UpdateKanbanTitle(userId, kanbanId, title)
}
