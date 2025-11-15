package entity

type Marker struct {
	ID     int64   `gorm:"column:id"`
	ChatID int64   `gorm:"column:chat_id"`
	Lat    float64 `gorm:"column:lat"`
	Lon    float64 `gorm:"column:lon"`
	Type   string  `gorm:"column:type"`
	Title  string  `gorm:"column:title"`
	Chat   Chat
}

func (Marker) TableName() string { return "markers" }
