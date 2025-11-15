package contract

import "firemap/internal/domain/entity"

type ChatRepository interface {
	Add(chat entity.Chat) (entity.Chat, error)
	GetById(Id int) (entity.Chat, error)
	GetAll() ([]entity.Chat, error)
}

type UserRepository interface {
	Add(user entity.User) (entity.User, error)
	FindByParams(params entity.User) (entity.User, error)
	FindByToken(token string) (entity.User, error)
}

type MessageRepository interface {
	Add(message entity.Message) error
}
