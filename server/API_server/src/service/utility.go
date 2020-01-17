package service

import (
	"orenotorero/repository"
)

type UtilityService struct {
	UserRepository repository.UserRepository
	KanbanRepository repository.KanbanRepository
}

func NewUtilityService(repository repository.UserRepository) UtilityService {
	return UtilityService{UserRepository: repository}
}

func (UtilitySvc *UtilityService) CheckEmail(email string) (bool, error) {
	flag, err := UtilitySvc.UserRepository.IsExistEmail(email)
	if err != nil {
		return true, err
	}
	return flag, nil
}

func (UtilitySvc *UtilityService) FileUpload(token, img string) (string, error) {
	var s3Url string
	// S3へ受け取った画像を保存後、保存先URLをDBへ保存する
	err := UtilitySvc.UserRepository.UpdateImg(s3Url)
	if err != nil {
		return "", nil
	}

	return s3Url, nil
}

func (UtilitySvc *UtilityService) UpdatePosition(userId string, position []int) error {
	//ここでリクエストを投げたUserがちゃんと該当Boardを所持しているかを確認する

	//Update
	err := UtilitySvc.KanbanRepository.UpdatePosition(position)
	if err != nil {
		return err
	}
	return nil
}
