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

func (CardSvc *CardService) CreateCard(userId, title string, kanbanId, position int) error {
	return CardSvc.CardRepository.InsertCard(userId, title, kanbanId, position)
}

func (CardSvc *CardService) ChangeCardTitle(userId string, cardId int, title string) error {
	return CardSvc.CardRepository.UpdateCardTitle(userId, cardId, title)
}

func (CardSvc *CardService) ChangeCardDeadline(cardId int, token string, deadline time.Time) error {
	return CardSvc.CardRepository.UpdateCardDeadLine(cardId, deadline)
}

func (CardSvc *CardService) AttachFile(userId string, cardId int, s3Url, fileName string) error {
	return CardSvc.FileRepository.AttachFile(userId, cardId, s3Url, fileName)
}

func (CardSvc *CardService) DeleteCard(userId string, cardId int) error {
	return CardSvc.CardRepository.DeleteCard(userId, cardId)
}
