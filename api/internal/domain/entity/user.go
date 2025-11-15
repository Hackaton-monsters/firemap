package entity

type User struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (User) TableName() string { return "users" }
