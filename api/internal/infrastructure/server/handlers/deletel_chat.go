package handlers

import (
	"firemap/internal/application/contract"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeleteChat struct {
	useCase contract.ChatDeleter
}

func NewDeleteChat(
	useCase contract.ChatDeleter,
) *DeleteChat {
	return &DeleteChat{
		useCase: useCase,
	}
}

func (h *DeleteChat) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	chatIDIDStr := c.Param("id")
	chatID, err := strconv.ParseInt(chatIDIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	err = h.useCase.DeleteChat(token, chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "chat successfully deleted"})
}
