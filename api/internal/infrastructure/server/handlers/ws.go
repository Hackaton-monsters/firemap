package handlers

import (
	"firemap/internal/infrastructure/chat"

	"github.com/gin-gonic/gin"
)

type WS struct {
	useCase *chat.Hub
}

func NewWS(
	useCase *chat.Hub,
) *WS {
	return &WS{
		useCase: useCase,
	}
}

func (h *WS) Handle(c *gin.Context) {
	chat.ServeWs(h.useCase, c.Writer, c.Request)
}
