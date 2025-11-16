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
	authMeHandler *handlers.AuthMe,
	createMarkerMeHandler *handlers.CreateMarker,
	getMarkersMeHandler *handlers.GetMarkers,
	getChatHistoryHandler *handlers.GetChatHistory,
	sendMessageHandler *handlers.SendMessage,
	connectToChatHandler *handlers.ConnectToChat,
	translateMessageHandler *handlers.TranslateMessage,
	getAllChatsHandler *handlers.GetAllChats,
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
		{
			Name:    "auth_me",
			Method:  http.MethodGet,
			Path:    "/auth/me",
			Handler: authMeHandler.Handle,
		},
		{
			Name:    "create_marker",
			Method:  http.MethodPost,
			Path:    "/marker",
			Handler: createMarkerMeHandler.Handle,
		},
		{
			Name:    "get_markers",
			Method:  http.MethodGet,
			Path:    "/marker/all",
			Handler: getMarkersMeHandler.Handle,
		},
		{
			Name:    "get_chat",
			Method:  http.MethodGet,
			Path:    "/chat/:id/history",
			Handler: getChatHistoryHandler.Handle,
		},
		{
			Name:    "get_all_chats",
			Method:  http.MethodGet,
			Path:    "/chat/all",
			Handler: getAllChatsHandler.Handle,
		},
		{
			Name:    "connect_chat",
			Method:  http.MethodPost,
			Path:    "/chat/:id/connect",
			Handler: connectToChatHandler.Handle,
		},
		{
			Name:    "send_message",
			Method:  http.MethodPost,
			Path:    "/message",
			Handler: sendMessageHandler.Handle,
		},
		{
			Name:    "translate_message",
			Method:  http.MethodGet,
			Path:    "/message/:id/translate/",
			Handler: translateMessageHandler.Handle,
		},
	}
}
