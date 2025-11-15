package service

import (
	"firemap/internal/domain/contract"
	"firemap/internal/domain/entity"
)

type ChatService interface {
	Create() (entity.Chat, error)
}

type chatService struct {
	repository contract.ChatRepository
}

func NewChatService(repository contract.ChatRepository) ChatService {
	return &chatService{
		repository: repository,
	}
}

func (s *chatService) Create() (entity.Chat, error) {
	return s.repository.Add(entity.Chat{})
}
