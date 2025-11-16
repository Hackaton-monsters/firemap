package handlers

import (
	"firemap/internal/application/contract"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetAllChats struct {
	useCase contract.ChatGetter
}

func NewGetAllChats(
	useCase contract.ChatGetter,
) *GetAllChats {
	return &GetAllChats{
		useCase: useCase,
	}
}

func (h *GetAllChats) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	response, err := h.useCase.GetAllChats(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
