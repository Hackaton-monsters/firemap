package repository

import (
	"errors"
	"firemap/internal/application/service"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type chatUserRepository struct {
	db *gorm.DB
}

func NewChatUserRepository(
	db *gorm.DB,
) contract.ChatUserRepository {
	return &chatUserRepository{
		db: db,
	}
}

func (r *chatUserRepository) Add(chatUser entity.ChatUser) (entity.ChatUser, error) {
	tx := r.db.Create(&chatUser)
	if tx.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(tx.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return chatUser, service.ErrUserAlreadyPresentInChat
		}
		return chatUser, tx.Error
	}
	return chatUser, nil
}
