package di

import (
	"firemap/internal/infrastructure/server/handlers"

	"github.com/google/wire"
)

var handlerSet = wire.NewSet(
	handlers.NewAuth,
	handlers.NewASignup,
)
