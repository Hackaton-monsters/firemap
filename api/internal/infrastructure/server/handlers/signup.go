package handlers

import (
	"firemap/internal/application/command"
	"firemap/internal/application/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Signup struct {
	useCase contract.UserRegistrator
}

func NewASignup(
	useCase contract.UserRegistrator,
) *Signup {
	return &Signup{
		useCase: useCase,
	}
}

func (h *Signup) Handle(c *gin.Context) {
	var request *command.RegisterUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.useCase.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
