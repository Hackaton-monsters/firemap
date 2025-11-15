package contract

import (
	"firemap/internal/application/command"
	"firemap/internal/application/response"
)

type MarkerCreator interface {
	CreateMarker(token string, command *command.CreateMarker) (*response.Marker, error)
}
