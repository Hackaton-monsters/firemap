package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type chatHistoryGetter struct {
	userService   service.UserService
	markerService service.MarkerService
	chatService   service.ChatService
}

func NewChatHistoryGetter(
	userService service.UserService,
	markerService service.MarkerService,
	chatService service.ChatService,
) contract.ChatHistoryGetter {
	return &chatHistoryGetter{
		userService:   userService,
		markerService: markerService,
		chatService:   chatService,
	}
}

func (u *chatHistoryGetter) GetChatHistory(token string, chatID int64) (*response.Chat, error) {
	_, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	marker, err := u.markerService.GetByChatId(chatID)
	if err != nil {
		return nil, err
	}

	messagesResponse := make([]response.Message, 0)
	for _, message := range marker.Chat.Messages {
		messagesResponse = append(messagesResponse, response.Message{
			ID:   message.ID,
			Text: message.Text,
			User: response.User{
				ID:       message.User.ID,
				Nickname: message.User.Nickname,
				Email:    message.User.Email,
				Role:     message.User.Role,
			},
			CreatedAt: message.CreatedAt,
		})
	}

	historyResponse := &response.Chat{
		ID: chatID,
		Marker: response.Marker{
			ID:     marker.ID,
			ChatID: marker.ChatID,
			Lat:    marker.Lat,
			Lon:    marker.Lon,
			Type:   marker.Type,
			Title:  marker.Title,
		},
		Messages: messagesResponse,
	}

	return historyResponse, nil
}
