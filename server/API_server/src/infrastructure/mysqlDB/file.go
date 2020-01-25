package mysqlDB

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orenotorero/db/Model"
	"orenotorero/repository"
	"orenotorero/utility"
)

type FileRepositoryImpliment struct {
	DB *gorm.DB
}

func NewFileRepoImpl(DB *gorm.DB) repository.FileRepository {
	return &FileRepositoryImpliment{
		DB: DB,
	}
}

func (p *FileRepositoryImpliment) AttachFile(userId string, cardId int, s3Url, fileName string) error {
	var card model.Card
	p.DB.Find(&card, cardId).Related(&card.Kanban)
	if card.Id == 0 {
		return errors.New("Cardが見つかりませんでした")
	}
	if card.Kanban.Id == 0 {
		return errors.New("Kanbanが見つかりませんでした")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, card.Kanban.BoardId)

	if isMyBoard {
		newFile := model.File{
			CardId: cardId,
			Name:   fileName,
			URL:    s3Url,
		}
		p.DB.Create(&newFile)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}

}

func (p *FileRepositoryImpliment) DeleteFile(userId string, fileId int) error {
	var file model.File
	p.DB.Preload("Card.Kanban").Find(&file, fileId)
	if file.Id == 0 {
		return errors.New("Fileが見つかりませんでした")
	}
	if file.Card.Id == 0 {
		return errors.New("Cardが見つかりませんでした")
	}
	if file.Card.Kanban.Id == 0 {
		return errors.New("Kanbanが見つかりませんでした")
	}

	isMyBoard := utility.IsMyBoard(p.DB, userId, file.Card.Kanban.BoardId)

	if isMyBoard {
		p.DB.Delete(&file)
		return nil
	} else {
		return errors.New("ボードへの権限がありません")
	}
}
