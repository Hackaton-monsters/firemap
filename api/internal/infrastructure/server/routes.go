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
	authHandler *handlers.Auth,
	signupHandler *handlers.Signup,
) []Route {
	return []Route{
		{
			Name:    "signup",
			Method:  http.MethodPost,
			Path:    "signup",
			Handler: signupHandler.Handle,
		},
		{
			Name:    "auth",
			Method:  http.MethodGet,
			Path:    "/auth/me",
			Handler: authHandler.Handle,
		},
	}
}
