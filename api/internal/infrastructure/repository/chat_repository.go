package repository

import (
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(
	db *gorm.DB,
) contract.ChatRepository {
	return &chatRepository{
		db: db,
	}
}

func (r *chatRepository) Add(chat entity.Chat) (entity.Chat, error) {
	tx := r.db.Create(&chat)
	if tx.Error != nil {
		return chat, tx.Error
	}
	return chat, nil
}

func (r *chatRepository) GetById(Id int) (entity.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (r *chatRepository) GetAll() ([]entity.Chat, error) {
	//TODO implement me
	panic("implement me")
}
