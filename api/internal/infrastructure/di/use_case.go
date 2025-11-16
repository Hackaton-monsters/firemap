package di

import (
	"firemap/internal/application/usecase"

	"github.com/google/wire"
)

var useCaseSet = wire.NewSet(
	usecase.NewUserAuthenticator,
	usecase.NewUserRegistrator,
	usecase.NewUserGetter,
	usecase.NewMarkerCreator,
	usecase.NewMarkersGetter,
	usecase.NewChatHistoryGetter,
	usecase.NewImageUploader,
	usecase.NewChatConnector,
	usecase.NewChatGetter,
)
