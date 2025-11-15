package entity

import "time"

type ChatUser struct {
	ChatID    int64     `gorm:"column:chat_id"`
	UserID    int64     `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Chat Chat `gorm:"foreignKey:ChatID;references:ID"`
	User User `gorm:"foreignKey:UserID;references:ID"`
}

func (ChatUser) TableName() string { return "chat_users" }
