package handlers

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateMarker struct {
	useCase contract.MarkerCreator
}

func NewCreateMarker(
	useCase contract.MarkerCreator,
) *CreateMarker {
	return &CreateMarker{
		useCase: useCase,
	}
}

func (h *CreateMarker) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	var request *command.CreateMarker
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.useCase.CreateMarker(token, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
