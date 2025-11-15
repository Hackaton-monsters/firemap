package handlers

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	useCase contract.UserAuthenticator
}

func NewAuth(
	useCase contract.UserAuthenticator,
) *Auth {
	return &Auth{
		useCase: useCase,
	}
}

func (h *Auth) Handle(c *gin.Context) {
	var request *command.AuthenticateUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.useCase.AuthenticateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
