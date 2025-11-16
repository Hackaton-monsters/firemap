package handlers

import (
	"context"
	"firemap/internal/application/command"
	"firemap/internal/infrastructure/chat"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SendMessage struct {
	useCase *chat.Hub
}

func NewSendMessage(
	useCase *chat.Hub,
) *SendMessage {
	return &SendMessage{
		useCase: useCase,
	}
}

func (h *SendMessage) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	var request *command.SendMessage
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.useCase.CreateAndBroadcastMessage(context.TODO(), token, *request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
