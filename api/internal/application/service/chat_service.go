package service

import (
	"errors"
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

var ErrChatNotFound = errors.New("chat not found")

type ChatService interface {
	Create() (entity.Chat, error)
	History(chatID int64) (entity.Chat, error)
}

type chatService struct {
	chatRepository    contract.ChatRepository
	messageRepository contract.MessageRepository
}

func NewChatService(
	chatRepository contract.ChatRepository,
	messageRepository contract.MessageRepository,
) ChatService {
	return &chatService{
		chatRepository:    chatRepository,
		messageRepository: messageRepository,
	}
}

func (s *chatService) Create() (entity.Chat, error) {
	return s.chatRepository.Add(entity.Chat{})
}

func (s *chatService) History(chatID int64) (entity.Chat, error) {
	return s.chatRepository.GetById(chatID)
}
