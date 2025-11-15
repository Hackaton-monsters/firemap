package handlers

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	useCase contract.UserAuthenticator
}

func NewLogin(
	useCase contract.UserAuthenticator,
) *Login {
	return &Login{
		useCase: useCase,
	}
}

func (h *Login) Handle(c *gin.Context) {
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
