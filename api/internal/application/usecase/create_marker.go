package usecase

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type markerCreator struct {
	userService     service.UserService
	markerService   service.MarkerService
	reportService   service.ReportService
	chatService     service.ChatService
	chatUserService service.ChatUserService
}

func NewMarkerCreator(
	userService service.UserService,
	markerService service.MarkerService,
	reportService service.ReportService,
	chatService service.ChatService,
	chatUserService service.ChatUserService,
) contract.MarkerCreator {
	return &markerCreator{
		userService:     userService,
		markerService:   markerService,
		reportService:   reportService,
		chatService:     chatService,
		chatUserService: chatUserService,
	}
}

func (u *markerCreator) CreateMarker(token string, command *command.CreateMarker) (*response.CreatedMarker, error) {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	chat, err := u.chatService.Create()
	if err != nil {
		return nil, err
	}

	marker, err := u.markerService.Create(*command, chat.ID)
	if err != nil {
		return nil, err
	}

	report, err := u.reportService.Create(*command, marker.ID)
	if err != nil {
		return nil, err
	}

	_, err = u.chatUserService.Connect(user.ID, chat.ID)
	if err != nil {
		return nil, err
	}

	markerResponse := response.Marker{
		ID:     marker.ID,
		ChatID: marker.ChatID,
		Lat:    marker.Lat,
		Lon:    marker.Lon,
		Reports: []response.Report{{
			ID:      report.ID,
			Comment: report.Comment,
			Photos:  report.Photos,
		}},
		ReportsCount: 1,
		Type:         marker.Type,
		Title:        marker.Title,
	}

	return &response.CreatedMarker{
		Marker:   markerResponse,
		IsNew:    true,
		IsMember: true,
	}, nil
}
