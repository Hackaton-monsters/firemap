package repository

import (
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(
	db *gorm.DB,
) contract.ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (r reportRepository) Add(report entity.Report) (entity.Report, error) {
	tx := r.db.Create(&report)
	if tx.Error != nil {
		return report, tx.Error
	}

	if report.MarkerID != 0 {
		if err := r.db.First(&report.Marker, report.MarkerID).Error; err != nil {
			return report, err
		}
	}
	return report, nil
}
