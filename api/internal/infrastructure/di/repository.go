package di

import (
	"firemap/internal/infrastructure/repository"

	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewChatRepository,
	repository.NewReportRepository,
	repository.NewMarkerRepository,
	repository.NewMessagesRepository,
	repository.NewChatUserRepository,
)
