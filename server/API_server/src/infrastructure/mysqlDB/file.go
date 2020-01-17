package mysqlDB

import (
	"github.com/jinzhu/gorm"
	model "orenotorero/db/Model"
)

type FileRepositoryImpliment struct {
	DB *gorm.DB
}

func (p *CardRepositoryImpliment) AttachFile(userId string, cardId int, s3Url, fileName string) error {
	newFile := model.File{
		CardId: cardId,
		Name:   fileName,
		URL:    s3Url,
	}
	p.DB.Create(&newFile)
	return nil
}
