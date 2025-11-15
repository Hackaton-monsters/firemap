package di

import (
	"firemap/internal/application/service"

	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	service.NewUserService,
	service.NewMarkerService,
	service.NewReportService,
	service.NewChatService,
)
