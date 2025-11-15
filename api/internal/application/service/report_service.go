package service

import (
	"firemap/internal/application/command"

	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

type ReportService interface {
	Create(marker command.CreateMarker, markerID int64) (entity.Report, error)
}

type reportService struct {
	repository contract.ReportRepository
}

func NewReportService(repository contract.ReportRepository) ReportService {
	return &reportService{
		repository: repository,
	}
}

func (s *reportService) Create(marker command.CreateMarker, markerID int64) (entity.Report, error) {
	return s.repository.Add(entity.Report{
		MarkerID: markerID,
		Comment:  marker.Comment,
		Photos:   marker.Photos,
	})
}
