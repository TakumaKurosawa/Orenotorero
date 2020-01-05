package service

import (
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type BoardService struct {
	BoardRepository repository.BoardRepository
}

func NewBoardService(repository repository.BoardRepository) BoardService {
	return BoardService{BoardRepository: repository}
}

func (boardSvc *BoardService) GetBoard(id string) ([]model.Board) {
	// 受け取ったTokenを元にuserIdを取得してボードへのアクセス権限があるかをチェックする
	return boardSvc.BoardRepository.SelectByUserId(id)
}

func (boardSvc *BoardService) CreateNewBoard(token string, title, img string) error {
	var userId int
	// 受け取ったTokenを元にuserIdを取得する

	// ボード新規作成機能
	err := boardSvc.BoardRepository.InsertBoard(userId, title, img)
	if err != nil {
		return err
	}

	return nil
}

func (boardSvc *BoardService) ChangePublishInfo(token string, boardId int, publish bool) error {
	//var userId int//
	// 受け取ったTokenを元にuserIdを取得し、ボードへのアクセス権限があるかを調べる

	// ボードの公開・非公開情報更新機能
	return boardSvc.BoardRepository.UpdateBoardPublish(boardId, publish)
}

func (boardSvc *BoardService) SendInviteMail(boardId int, token, email, message string) error {
	//var userId int//
	// 受け取ったTokenを元にuserIdを取得し、ボードへのアクセス権限があるかを調べる

	// 招待メール送信機能
	return nil
}
