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
