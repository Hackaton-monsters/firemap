package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type chatConnector struct {
	userService     service.UserService
	chatService     service.ChatService
	markerService   service.MarkerService
	chatUserService service.ChatUserService
}

func NewChatConnector(
	userService service.UserService,
	chatService service.ChatService,
	markerService service.MarkerService,
	chatUserService service.ChatUserService,
) contract.ChatConnector {
	return &chatConnector{
		userService:     userService,
		chatService:     chatService,
		markerService:   markerService,
		chatUserService: chatUserService,
	}
}

func (u *chatConnector) ConnectToChat(token string, chatID int64) error {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return err
	}

	chat, err := u.chatService.History(chatID)
	if err != nil {
		return err
	}

	_, err = u.chatUserService.Connect(user.ID, chat.ID)
	if err != nil {
		return err
	}

	marker, err := u.markerService.GetByChatId(chat.ID)
	if err != nil {
		return err
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

	return nil
}
