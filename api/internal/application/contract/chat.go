package contract

import (
	"firemap/internal/application/response"
)

type ChatHistoryGetter interface {
	GetChatHistory(token string, chatID int64) (*response.Chat, error)
}
