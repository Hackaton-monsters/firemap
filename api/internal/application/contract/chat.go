package contract

import (
	"firemap/internal/application/response"
)

type ChatHistoryGetter interface {
	GetChatHistory(token string, chatID int64) (*response.Chat, error)
}

type ChatConnector interface {
	ConnectToChat(token string, chatID int64) error
}

type ChatGetter interface {
	GetAllChats(token string) (*response.Chats, error)
}

type ChatDeleter interface {
	DeleteChat(token string, chatID int64) error
}
