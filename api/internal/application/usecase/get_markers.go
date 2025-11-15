package usecase

import (
	"firemap/internal/application/contract"
	"firemap/internal/application/response"
	"firemap/internal/application/service"
)

type markersGetter struct {
	userService   service.UserService
	markerService service.MarkerService
}

func NewMarkersGetter(
	userService service.UserService,
	markerService service.MarkerService,
) contract.MarkerGetter {
	return &markersGetter{
		userService:   userService,
		markerService: markerService,
	}
}

func (u *markersGetter) GetMarkers(token string) (*response.Markers, error) {
	_, err := u.userService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	markers, _ := u.markerService.GetAll()

	var markersResponse []*response.Marker
	for _, marker := range markers {
		reportsResponse := make([]response.Report, 0)
		for _, report := range marker.Reports {
			reportsResponse = append(reportsResponse, response.Report{
				ID:      report.ID,
				Comment: report.Comment,
				Photos:  report.Photos,
			})
		}

		markersResponse = append(markersResponse, &response.Marker{
			ID:           marker.ID,
			ChatID:       marker.ChatID,
			Lat:          marker.Lat,
			Lon:          marker.Lon,
			Reports:      reportsResponse,
			ReportsCount: len(marker.Reports),
			Type:         marker.Type,
			Title:        marker.Title,
		})
	}

	return &response.Markers{Markers: markersResponse}, nil
}
