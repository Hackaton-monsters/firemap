package di

import (
	"database/sql"
	"embed"
	"firemap/internal/infrastructure/chat"
	"firemap/internal/infrastructure/config"
	"firemap/internal/infrastructure/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
)

type ProcessManager struct {
	config *config.Config
	db     *sql.DB
	routes []server.Route
	hub    *chat.Hub
}

func NewProcessManager(
	config *config.Config,
	db *sql.DB,
	routes []server.Route,
	hub *chat.Hub,
) *ProcessManager {
	return &ProcessManager{
		config: config,
		db:     db,
		routes: routes,
		hub:    hub,
	}
}

//go:generate cp -r ../../../migrations ./tmp-migrations
//go:embed tmp-migrations/*.sql
var embedMigrations embed.FS

func (pm *ProcessManager) Migrate() {

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(pm.db, "tmp-migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations successfully")
}

func (pm *ProcessManager) RunHTTPServer() {
	go pm.hub.Run()
	router := gin.Default()
	for _, route := range pm.routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Path, route.Handler)
		case http.MethodPost:
			router.POST(route.Path, route.Handler)
		case http.MethodPut:
			router.PUT(route.Path, route.Handler)
		case http.MethodDelete:
			router.DELETE(route.Path, route.Handler)
		}
	}

	router.Run(fmt.Sprintf(":%d", pm.config.HttpServerPort))
}
