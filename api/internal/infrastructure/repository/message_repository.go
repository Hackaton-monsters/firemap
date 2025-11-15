package repository

import (
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"

	"gorm.io/gorm"
)

type messagesRepository struct {
	db *gorm.DB
}

func NewMessagesRepository(
	db *gorm.DB,
) contract.MessageRepository {
	return &messagesRepository{
		db: db,
	}
}

func (r *messagesRepository) Add(message entity.Message) (entity.Message, error) {
	tx := r.db.Create(&message)
	if tx.Error != nil {
		return message, tx.Error
	}

	if message.UserID != 0 {
		if err := r.db.First(&message.User, message.UserID).Error; err != nil {
			return message, err
		}
	}
	return message, nil
}

func (r *messagesRepository) GetAll(chatID int64) ([]entity.Message, error) {
	var messages []entity.Message

	if err := r.db.
		Preload("Chat").
		Preload("User").
		Where("chat_id = ?", chatID).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *messagesRepository) GetById(messageID int64) (*entity.Message, error) {
	var message entity.Message
	if err := r.db.First(&message, messageID).Error; err != nil {
		return nil, err
	}

	return &message, nil
}
