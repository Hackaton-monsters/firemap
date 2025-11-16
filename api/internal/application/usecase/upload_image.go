package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/service"
	"firemap/internal/domain/entity"
	"mime/multipart"
)

type imageUploader struct {
	imageService service.ImageService
	userService  service.UserService
}

func NewImageUploader(
	imageService service.ImageService,
	userService service.UserService,
) contract.ImageUploader {
	return &imageUploader{
		imageService: imageService,
		userService:  userService,
	}
}

func (u *imageUploader) UploadImage(token string, file multipart.File, filename string) (int64, error) {
	_, err := u.userService.FindByToken(token)
	if err != nil {
		return 0, err
	}

	fileURL, err := u.imageService.UploadImage(file, filename)

	if err != nil {
		return 0, err
	}

	image := entity.Image{
		URL: fileURL,
	}

	savedImage, err := u.imageService.Save(image)
	if err != nil {
		return 0, err
	}

	return savedImage.ID, nil
}
