package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
	"fmt"
)

type chatGetter struct {
	userService   service.UserService
	markerService service.MarkerService
	chatService   service.ChatService
}

func NewChatGetter(
	userService service.UserService,
	markerService service.MarkerService,
	chatService service.ChatService,
) contract.ChatGetter {
	return &chatGetter{
		userService:   userService,
		markerService: markerService,
		chatService:   chatService,
	}
}

func (u *chatGetter) GetAllChats(token string) (*response.Chats, error) {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	fmt.Println(user.Email)
	fmt.Println(user.Email)
	fmt.Println(len(user.Chats))
	fmt.Println(len(user.Chats))
	fmt.Println(len(user.Chats))
	fmt.Println(len(user.Chats))

	responseChats := make([]response.Chat, 0)

	for _, chat := range user.Chats {
		marker, err := u.markerService.GetByChatId(chat.ID)
		if err != nil {
			return nil, err
		}

		responseChat := response.Chat{
			ID: chat.ID,
			Marker: response.Marker{
				ID:     marker.ID,
				ChatID: marker.ChatID,
				Lat:    marker.Lat,
				Lon:    marker.Lon,
				Type:   marker.Type,
				Title:  marker.Title,
			},
		}

		responseChats = append(responseChats, responseChat)
	}

	return &response.Chats{
		Chats: responseChats,
	}, err
}
