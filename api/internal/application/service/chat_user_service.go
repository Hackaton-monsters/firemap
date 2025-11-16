package service

import (
	"errors"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

var ErrUserAlreadyPresentInChat = errors.New("user already present in chat")

type ChatUserService interface {
	Connect(userID int64, chatID int64) (entity.ChatUser, error)
}

type chatUserService struct {
	chatUserRepository contract.ChatUserRepository
}

func NewChatUserService(
	chatUserRepository contract.ChatUserRepository,
) ChatUserService {
	return &chatUserService{
		chatUserRepository: chatUserRepository,
	}
}

func (s *chatUserService) Connect(userID int64, chatID int64) (entity.ChatUser, error) {
	return s.chatUserRepository.Add(entity.ChatUser{
		UserID: userID,
		ChatID: chatID,
	})
}
