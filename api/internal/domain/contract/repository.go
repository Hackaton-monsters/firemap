package contract

import "firemap/internal/domain/entity"

type ChatRepository interface {
	Add(chat entity.Chat) (entity.Chat, error)
	GetById(ID int64) (entity.Chat, error)
	GetAll() ([]entity.Chat, error)
}

type UserRepository interface {
	Add(user entity.User) (entity.User, error)
	FindByParams(params entity.User) (entity.User, error)
	FindByToken(token string) (entity.User, error)
}

type MessageRepository interface {
	Add(message entity.Message) (entity.Message, error)
	GetAll(chatID int64) ([]entity.Message, error)
	GetById(messageID int64) (*entity.Message, error)
}

type MarkerRepository interface {
	Add(marker entity.Marker) (entity.Marker, error)
	GetAll() ([]entity.Marker, error)
	GetByChatID(chatId int64) (entity.Marker, error)
}

type ReportRepository interface {
	Add(report entity.Report) (entity.Report, error)
}
