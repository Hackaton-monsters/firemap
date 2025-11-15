package service

import (
	"firemap/internal/application/command"

	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

type MarkerService interface {
	Create(marker command.CreateMarker, chatID int64) (entity.Marker, error)
	GetAll() ([]entity.Marker, error)
	GetByChatId(chatID int64) (entity.Marker, error)
}

type markerService struct {
	repository contract.MarkerRepository
}

func NewMarkerService(repository contract.MarkerRepository) MarkerService {
	return &markerService{
		repository: repository,
	}
}

func (s *markerService) Create(marker command.CreateMarker, chatID int64) (entity.Marker, error) {
	return s.repository.Add(entity.Marker{
		ChatID: chatID,
		Lat:    marker.Lat,
		Lon:    marker.Lon,
		Type:   marker.Type,
		Title:  "test",
	})
}

func (s *markerService) GetAll() ([]entity.Marker, error) {
	return s.repository.GetAll()
}

func (s *markerService) GetByChatId(chatID int64) (entity.Marker, error) {
	return s.repository.GetByChatID(chatID)
}
