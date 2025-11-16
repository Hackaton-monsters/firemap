package service

import (
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
	"mime/multipart"
)

type ImageService interface {
	UploadImage(file multipart.File, filename string) (string, error)
	DeleteImage(imageKey string) error
	Save(image entity.Image) (entity.Image, error)
	GetByID(ID int64) (entity.Image, error)
}

type imageService struct {
	imageRepository contract.ImageRepository
}

func NewImageService(
	imageRepository contract.ImageRepository,
) ImageService {
	return &imageService{
		imageRepository: imageRepository,
	}
}

func (s *imageService) UploadImage(file multipart.File, filename string) (string, error) {
	return s.imageRepository.UploadImage(file, filename)
}

func (s *imageService) DeleteImage(imageKey string) error {
	return s.imageRepository.DeleteImage(imageKey)
}

func (s *imageService) Save(image entity.Image) (entity.Image, error) {
	return s.imageRepository.Save(image)
}

func (s *imageService) GetByID(ID int64) (entity.Image, error) {
	return s.imageRepository.GetByID(ID)
}
