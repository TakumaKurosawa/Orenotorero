package service

import (
	"orenotorero/repository"
	"time"
)

type CardService struct {
	CardRepository repository.CardRepository
	FileRepository repository.FileRepository
}

func NewCardService(repository repository.CardRepository) CardService {
	return CardService{CardRepository: repository}
}

func (CardSvc *CardService) CreateCard(token, title string, kanbanId, position int) error {
	// tokenを使ってアクセス権限の有無
	return CardSvc.CardRepository.InsertCard(title, kanbanId, position)
}

func (CardSvc *CardService) ChangeCardTitle(cardId int, token, title string) error {
	return CardSvc.CardRepository.UpdateCardTitle(cardId, title)
}

func (CardSvc *CardService) ChangeCardDeadline(cardId int, token string, deadline time.Time) error {
	return CardSvc.CardRepository.UpdateCardDeadLine(cardId, deadline)
}

func (CardSvc *CardService) InsertFileData(userId string, cardId int, s3Url, fileName string) error {
	// FileテーブルにS3に保存したファイルへのURLなどの情報を格納する処理
	return CardSvc.FileRepository.AttachFile(userId, cardId, s3Url, fileName)
}
