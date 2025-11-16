package repository

import (
	"context"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
	"firemap/internal/infrastructure/config"
	"firemap/internal/infrastructure/s3"
	"fmt"

	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type imageRepository struct {
	db               *gorm.DB
	cloudinaryClient s3.CloudinaryClient
	ctx              context.Context
}

func NewImageRepository(db *gorm.DB, cloudinaryClient s3.CloudinaryClient, config *config.Config) contract.ImageRepository {
	ctx := context.Background()

	service := &imageRepository{
		db:               db,
		cloudinaryClient: cloudinaryClient,
		ctx:              ctx,
	}

	return service
}

func (s *imageRepository) UploadImage(file multipart.File, filename string) (string, error) {
	publicID := fmt.Sprintf("firemap/images/%s_%s", uuid.New().String()[:8], filename)

	url, err := s.cloudinaryClient.UploadImage(s.ctx, file, publicID)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *imageRepository) DeleteImage(publicID string) error {
	return s.cloudinaryClient.DeleteImage(s.ctx, publicID)
}

func (s *imageRepository) Save(image entity.Image) (entity.Image, error) {
	tx := s.db.Create(&image)

	if tx.Error != nil {
		return image, tx.Error
	}

	return image, nil
}

func (s *imageRepository) GetByID(ID int64) (entity.Image, error) {
	var image entity.Image

	tx := s.db.First(&image, ID)
	if tx.Error != nil {
		return image, tx.Error
	}

	return image, nil
}
