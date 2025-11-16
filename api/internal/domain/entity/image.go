package entity

type Image struct {
	ID  int64  `gorm:"column:id"`
	URL string `gorm:"column:url"`
}

func (Image) TableName() string { return "image" }
