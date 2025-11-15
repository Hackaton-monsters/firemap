package server

import "github.com/gin-gonic/gin"

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func NewRoutes() []Route {
	return []Route{}
}
