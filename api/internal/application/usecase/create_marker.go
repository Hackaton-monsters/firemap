package usecase

import (
	"context"
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
	"firemap/internal/infrastructure/geo_ip"
	"fmt"
)

type markerCreator struct {
	userService     service.UserService
	markerService   service.MarkerService
	reportService   service.ReportService
	chatService     service.ChatService
	chatUserService service.ChatUserService
	imageService    service.ImageService
	infoGetter      geo_ip.InfoGetter
}

func NewMarkerCreator(
	userService service.UserService,
	markerService service.MarkerService,
	reportService service.ReportService,
	chatService service.ChatService,
	chatUserService service.ChatUserService,
	imageService service.ImageService,
	infoGetter geo_ip.InfoGetter,
) contract.MarkerCreator {
	return &markerCreator{
		userService:     userService,
		markerService:   markerService,
		reportService:   reportService,
		chatService:     chatService,
		chatUserService: chatUserService,
		imageService:    imageService,
		infoGetter:      infoGetter,
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

	title, err := u.infoGetter.GetDisplayNameByCoordinate(context.TODO(), command.Lat, command.Lon)
	if err != nil {
		fmt.Println(err)
		title = fmt.Sprintf("%f %f", command.Lat, command.Lon)
	}

	title = fmt.Sprintf("%s - %s", command.Type, title)

	marker, err := u.markerService.Create(*command, chat.ID, title)
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

	photoURLs := make([]string, 0, len(report.Photos))
	for _, photoID := range report.Photos {
		image, err := u.imageService.GetByID(photoID)
		if err != nil {
			continue
		}
		photoURLs = append(photoURLs, image.URL)
	}

	markerResponse := &response.Marker{
		ID:     marker.ID,
		ChatID: marker.ChatID,
		Lat:    marker.Lat,
		Lon:    marker.Lon,
		Reports: []response.Report{{
			ID:      report.ID,
			Comment: report.Comment,
			Photos:  photoURLs,
		}},
		ReportsCount: 1,
		Type:         marker.Type,
		Title:        marker.Title,
	}

	return &response.CreatedMarker{
		Marker:   *markerResponse,
		IsNew:    true,
		IsMember: true,
	}, nil
}
