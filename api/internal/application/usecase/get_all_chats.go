package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type chatGetter struct {
	userService   service.UserService
	markerService service.MarkerService
	chatService   service.ChatService
	imageService  service.ImageService
}

func NewChatGetter(
	userService service.UserService,
	markerService service.MarkerService,
	chatService service.ChatService,
	imageService service.ImageService,
) contract.ChatGetter {
	return &chatGetter{
		userService:   userService,
		markerService: markerService,
		chatService:   chatService,
		imageService:  imageService,
	}
}

func (u *chatGetter) GetAllChats(token string) (*response.Chats, error) {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	responseChats := make([]response.Chat, 0)

	for _, chat := range user.Chats {
		marker, err := u.markerService.GetByChatId(chat.ID)
		if err != nil {
			return nil, err
		}

		messages := make([]response.Message, 0)

		if len(marker.Chat.Messages) > 0 {
			lastMessageI := len(marker.Chat.Messages) - 1

			message := response.Message{
				ID:   marker.Chat.Messages[lastMessageI].ID,
				Text: marker.Chat.Messages[lastMessageI].Text,
				User: response.User{
					ID:       marker.Chat.Messages[lastMessageI].User.ID,
					Nickname: marker.Chat.Messages[lastMessageI].User.Nickname,
					Email:    marker.Chat.Messages[lastMessageI].User.Email,
					Role:     marker.Chat.Messages[lastMessageI].User.Role,
				},
				CreatedAt: marker.Chat.Messages[lastMessageI].CreatedAt,
			}
			messages = append(messages, message)
		}

		reportsResponse := make([]response.Report, 0)
		for _, report := range marker.Reports {
			photoURLs := make([]string, 0, len(report.Photos))
			for _, photoID := range report.Photos {
				image, err := u.imageService.GetByID(photoID)
				if err != nil {
					continue
				}
				photoURLs = append(photoURLs, image.URL)
			}

			reportsResponse = append(reportsResponse, response.Report{
				ID:      report.ID,
				Comment: report.Comment,
				Photos:  photoURLs,
			})
		}

		responseChat := response.Chat{
			ID: chat.ID,
			Marker: response.Marker{
				ID:           marker.ID,
				ChatID:       marker.ChatID,
				Lat:          marker.Lat,
				Lon:          marker.Lon,
				Type:         marker.Type,
				Title:        marker.Title,
				Reports:      reportsResponse,
				ReportsCount: len(reportsResponse),
			},
			Messages:  messages,
			CreatedAt: chat.CreatedAt,
		}

		responseChats = append(responseChats, responseChat)
	}

	chats := response.Chats{
		Chats: responseChats,
	}

	chats.SortByLastActivityDesc()

	return &chats, err
}
