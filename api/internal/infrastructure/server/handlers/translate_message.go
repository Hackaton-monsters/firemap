package handlers

import (
	"firemap/internal/domain/contract"
	"firemap/internal/infrastructure/translator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TranslateMessage struct {
	messageRepo contract.MessageRepository
	translator  translator.Translator
}

func NewTranslateMessage(
	messageRepo contract.MessageRepository,
	translator translator.Translator,
) *TranslateMessage {
	return &TranslateMessage{
		messageRepo: messageRepo,
		translator:  translator,
	}
}

func (h *TranslateMessage) Handle(c *gin.Context) {
	messageIDStr := c.Param("id")
	messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	language := c.Query("language")
	if language == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language parameter is required"})
		return
	}

	message, err := h.messageRepo.GetById(messageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	translatedMessage, err := h.translator.Translate(c, message.Text, language)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Translation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             message.ID,
		"translatedText": translatedMessage,
	})
}
