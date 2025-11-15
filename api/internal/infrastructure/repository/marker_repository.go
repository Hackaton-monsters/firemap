package repository

import (
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
