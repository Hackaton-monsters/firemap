package entity

import "time"

type Report struct {
	ID        int64     `gorm:"column:id"`
	MarkerID  int64     `gorm:"column:marker_id"`
	Comment   string    `gorm:"column:comment"`
	Photos    []int64   `gorm:"column:photos;serializer:json"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Marker    Marker
}

func (Report) TableName() string { return "reports" }
