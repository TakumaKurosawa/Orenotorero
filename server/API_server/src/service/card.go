package service

import (
	"orenotorero/repository"
	"time"
)

type CardService struct {
	CardRepository repository.CardRepository
}

func NewCardService(repository repository.CardRepository) CardService {
	return CardService{CardRepository: repository}
}

func (CardSvc *CardService) CreateCard(userId, title string, kanbanId, position int) error {
	return CardSvc.CardRepository.InsertCard(userId, title, kanbanId, position)
}

func (CardSvc *CardService) ChangeCardTitle(cardId int, token, title string) error {
	return CardSvc.CardRepository.UpdateCardTitle(cardId, title)
}

func (CardSvc *CardService) ChangeCardDeadline(userId string, cardId int, deadline time.Time) error {
	return CardSvc.CardRepository.UpdateCardDeadLine(userId, cardId, deadline)
}

func (CardSvc *CardService) InsertFileData(cardId int, token, s3Url, fileName string) error {
	// FileテーブルにS3に保存したファイルへのURLなどの情報を格納する処理
	return nil
}
