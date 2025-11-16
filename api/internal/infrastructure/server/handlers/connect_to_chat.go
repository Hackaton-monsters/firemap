package handlers

import (
	"firemap/internal/application/contract"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ConnectToChat struct {
	useCase contract.ChatConnector
}

func NewConnectToChat(
	useCase contract.ChatConnector,
) *ConnectToChat {
	return &ConnectToChat{
		useCase: useCase,
	}
}

func (h *ConnectToChat) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	chatIDStr := c.Param("id")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat id"})
		return
	}

	err = h.useCase.ConnectToChat(token, chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully connected to chat"})
}
