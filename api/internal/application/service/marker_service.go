package service

import (
	"firemap/internal/application/command"
	"math"

	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

type MarkerService interface {
	Create(marker command.CreateMarker, chatID int64, title string) (entity.Marker, error)
	GetAll() ([]entity.Marker, error)
	GetByChatId(chatID int64) (entity.Marker, error)

	FindNearestWithinRadius(lat, lon, radiusMeters float64) (entity.Marker, bool, error)
}

type markerService struct {
	repository contract.MarkerRepository
}

func NewMarkerService(repository contract.MarkerRepository) MarkerService {
	return &markerService{
		repository: repository,
	}
}

func (s *markerService) Create(marker command.CreateMarker, chatID int64, title string) (entity.Marker, error) {
	return s.repository.Add(entity.Marker{
		ChatID: chatID,
		Lat:    marker.Lat,
		Lon:    marker.Lon,
		Type:   marker.Type,
		Title:  title,
	})
}

func (s *markerService) GetAll() ([]entity.Marker, error) {
	return s.repository.GetAll()
}

func (s *markerService) GetByChatId(chatID int64) (entity.Marker, error) {
	return s.repository.GetByChatID(chatID)
}

func (s *markerService) FindNearestWithinRadius(lat, lon, radiusMeters float64) (entity.Marker, bool, error) {
	markers, err := s.repository.GetAll()
	if err != nil {
		return entity.Marker{}, false, err
	}

	const earthRadius = 6371000.0 // м
	toRad := func(deg float64) float64 {
		return deg * math.Pi / 180
	}

	var (
		closest entity.Marker
		minDist = math.MaxFloat64
	)

	for _, m := range markers {
		dLat := toRad(m.Lat - lat)
		dLon := toRad(m.Lon - lon)
		lat1 := toRad(lat)
		lat2 := toRad(m.Lat)

		a := math.Sin(dLat/2)*math.Sin(dLat/2) +
			math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
		c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		dist := earthRadius * c

		if dist < radiusMeters && dist < minDist {
			minDist = dist
			closest = m
		}
	}

	if minDist == math.MaxFloat64 {
		return entity.Marker{}, false, nil
	}

	return closest, true, nil
}
