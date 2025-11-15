package di

import (
	"firemap/internal/application/usecase"

	"github.com/google/wire"
)

var useCaseSet = wire.NewSet(
	usecase.NewUserAuthenticator,
	usecase.NewUserRegistrator,
)
