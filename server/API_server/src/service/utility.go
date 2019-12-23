package service

import (
	"orenotorero/repository"
)

type UtilityService struct {
	UserRepository repository.UserRepository
}

func NewUtilityService(repository repository.UserRepository) UtilityService {
	return UtilityService{UserRepository: repository}
}

func (UtilitySvc *UtilityService) CheckEmail(email string) (bool, error) {
	user, err := UtilitySvc.UserRepository.SelectByEmail(email)
	if err != nil {
		return true, err
	}

	if user.Email != "" {
		return true, nil
	}

	return false, nil
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
