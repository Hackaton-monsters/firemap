package entity

import "time"

type Chat struct {
	ID        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Messages  []Message `gorm:"foreignKey:ChatID;references:ID"`
}

type Message struct {
	ID        int64     `gorm:"column:id"`
	ChatID    int64     `gorm:"column:chat_id"`
	UserID    int64     `gorm:"column:user_id"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      *User
}

func (Chat) TableName() string    { return "chats" }
func (Message) TableName() string { return "messages" }
