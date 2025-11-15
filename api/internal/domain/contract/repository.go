package contract

import "firemap/internal/domain/entity"

type ChatRepository interface {
	Add(chat entity.Chat) (entity.Chat, error)
	GetById(Id int) (entity.Chat, error)
	GetAll() ([]entity.Chat, error)
}

type MessageRepository interface {
	Add(message entity.Message) error
}
