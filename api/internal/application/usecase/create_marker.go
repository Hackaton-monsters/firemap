package usecase

import (
	"context"
	"errors"
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
	"firemap/internal/infrastructure/geo_ip"
	"fmt"
)

const mergeRadiusMeters = 500.0

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

func (u *markerCreator) CreateMarker(token string, cmd *command.CreateMarker) (*response.CreatedMarker, error) {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	// ===== Пытаемся объединять ТОЛЬКО если создаём "fire" =====
	if cmd.Type == "fire" {
		existingMarker, found, err := u.markerService.FindNearestWithinRadius(cmd.Lat, cmd.Lon, mergeRadiusMeters)
		if err != nil {
			return nil, err
		}

		// объединяем только если рядом есть маркер и он тоже типа "fire"
		if found && existingMarker.Type == "fire" {
			// создаём только новый report
			report, err := u.reportService.Create(*cmd, existingMarker.ID)
			if err != nil {
				return nil, err
			}

			// подключаем юзера к чату этого маркера (может уже быть в чате)
			_, err = u.chatUserService.Connect(user.ID, existingMarker.ChatID)
			if err != nil && !errors.Is(err, service.ErrUserAlreadyPresentInChat) {
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

			// считаем, что к существующим репортам добавился ещё один
			reportsCount := len(existingMarker.Reports) + 1

			markerResponse := &response.Marker{
				ID:     existingMarker.ID,
				ChatID: existingMarker.ChatID,
				Lat:    existingMarker.Lat,
				Lon:    existingMarker.Lon,
				Reports: []response.Report{{
					ID:      report.ID,
					Comment: report.Comment,
					Photos:  photoURLs,
				}},
				ReportsCount: reportsCount,
				Type:         existingMarker.Type,
				Title:        existingMarker.Title,
			}

			return &response.CreatedMarker{
				Marker:   *markerResponse,
				IsNew:    false, // ← ВАЖНО: НЕ новый маркер, мы присоединились к существующему
				IsMember: true,
			}, nil
		}
	}

	// ===== сюда попадаем если:
	// - type != "fire", или
	// - type == "fire", но рядом нет fire-маркера =====

	chat, err := u.chatService.Create()
	if err != nil {
		return nil, err
	}

	title, err := u.infoGetter.GetDisplayNameByCoordinate(context.TODO(), cmd.Lat, cmd.Lon)
	fmt.Println(title)
	fmt.Println(title)
	fmt.Println(title)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println(err)
	if err != nil || title == "" {
		title = fmt.Sprintf("%f %f", cmd.Lat, cmd.Lon)
	}

	marker, err := u.markerService.Create(*cmd, chat.ID, title)
	if err != nil {
		return nil, err
	}

	report, err := u.reportService.Create(*cmd, marker.ID)
	if err != nil {
		return nil, err
	}

	_, err = u.chatUserService.Connect(user.ID, chat.ID)
	if err != nil && !errors.Is(err, service.ErrUserAlreadyPresentInChat) {
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
