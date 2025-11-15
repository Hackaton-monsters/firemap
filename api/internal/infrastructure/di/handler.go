package di

import (
	"firemap/internal/infrastructure/server/handlers"

	"github.com/google/wire"
)

var handlerSet = wire.NewSet(
	handlers.NewLogin,
	handlers.NewRegister,
	handlers.NewAuthMe,
	handlers.NewCreateMarker,
	handlers.NewGetMarkers,
	handlers.NewGetChatHistory,
	handlers.NewTranslateMessage,
)
