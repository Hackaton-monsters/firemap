package repository

import (
	"errors"
	"firemap/internal/application/service"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"

	"gorm.io/gorm"
)

type markerRepository struct {
	db *gorm.DB
}

func NewMarkerRepository(
	db *gorm.DB,
) contract.MarkerRepository {
	return &markerRepository{
		db: db,
	}
}

func (r *markerRepository) Add(marker entity.Marker) (entity.Marker, error) {
	tx := r.db.Create(&marker)
	if tx.Error != nil {
		return marker, tx.Error
	}

	if marker.ChatID != 0 {
		if err := r.db.First(&marker.Chat, marker.ChatID).Error; err != nil {
			return marker, err
		}
	}
	return marker, nil
}

func (r *markerRepository) GetAll() ([]entity.Marker, error) {
	var markers []entity.Marker

	if err := r.db.
		Preload("Chat").
		Preload("Reports").
		Find(&markers).Error; err != nil {
		return nil, err
	}

	return markers, nil
}

func (r *markerRepository) GetByChatID(chatId int64) (entity.Marker, error) {
	var marker entity.Marker

	err := r.db.
		Preload("Chat.Messages").
		Preload("Chat.Messages.User").
		Where("chat_id = ?", chatId).
		First(&marker).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return marker, service.ErrChatNotFound
		}
		return marker, err
	}

	return marker, nil
}
