package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type markersGetter struct {
	userService   service.UserService
	markerService service.MarkerService
	imageService  service.ImageService
}

func NewMarkersGetter(
	userService service.UserService,
	markerService service.MarkerService,
	imageService service.ImageService,
) contract.MarkerGetter {
	return &markersGetter{
		userService:   userService,
		markerService: markerService,
		imageService:  imageService,
	}
}

func (u *markersGetter) GetMarkers(token string) (*response.Markers, error) {
	user, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	markers, _ := u.markerService.GetAll()

	var markersResponse = make([]*response.MapMarker, 0)
	for _, marker := range markers {
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

		IsMember := false

		for _, userChat := range user.Chats {
			if userChat.ID == marker.ChatID {
				IsMember = true
				break
			}
		}

		markersResponse = append(markersResponse, &response.MapMarker{
			Marker: response.Marker{
				ID:           marker.ID,
				ChatID:       marker.ChatID,
				Lat:          marker.Lat,
				Lon:          marker.Lon,
				Reports:      reportsResponse,
				ReportsCount: len(marker.Reports),
				Type:         marker.Type,
				Title:        marker.Title,
			},
			IsMember: IsMember,
		})
	}

	return &response.Markers{Markers: markersResponse}, nil
}
