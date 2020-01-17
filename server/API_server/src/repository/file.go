package repository

type FileRepository interface {
	AttachFile(userId string, cardId int, s3Url, fileName string) error
}
