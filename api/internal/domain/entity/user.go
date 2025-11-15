package entity

type User struct {
	ID       int64  `gorm:"column:id"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Nickname string `gorm:"column:nickname"`
	Role     string `gorm:"column:role"`
	Token    string `gorm:"column:token"`
	Chats    []Chat `gorm:"many2many:chat_users;"`
}

func (User) TableName() string { return "users" }
