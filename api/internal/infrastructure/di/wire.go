//go:build wireinject
// +build wireinject

package di

import (
	"firemap/internal/infrastructure/config"
	"firemap/internal/infrastructure/db"
	"firemap/internal/infrastructure/server"

	"github.com/google/wire"
)

func InitializeProcessManager() *ProcessManager {
	wire.Build(
		NewProcessManager,
		repositorySet,
		serviceSet,
		handlerSet,
		useCaseSet,
		db.NewDB,
		db.NewDBForMigrations,
		config.LoadFromEnvironment,
		server.NewRoutes,
	)
	return &ProcessManager{}
}
