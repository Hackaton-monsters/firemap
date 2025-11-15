package repository

import (
	"errors"
	"firemap/internal/application/service"
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

func (r *chatRepository) GetById(id int64) (entity.Chat, error) {
	var chat entity.Chat

	err := r.db.
		Preload("Messages").
		Preload("Users").
		Where("id = ?", id).
		First(&chat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return chat, service.ErrChatNotFound
		}
		return chat, err
	}
	return chat, nil
}

func (r *chatRepository) GetAll() ([]entity.Chat, error) {
	//TODO implement me
	panic("implement me")
}
