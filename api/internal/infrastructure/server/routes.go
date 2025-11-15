package server

import (
	"firemap/internal/infrastructure/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func NewRoutes(
	loginHandler *handlers.Login,
	registerHandler *handlers.Register,
) []Route {
	return []Route{
		{
			Name:    "register",
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: registerHandler.Handle,
		},
		{
			Name:    "login",
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: loginHandler.Handle,
		},
	}
}
