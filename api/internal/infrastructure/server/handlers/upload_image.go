package handlers

import (
	"firemap/internal/application/contract"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadImage struct {
	useCase contract.ImageUploader
}

func NewUploadImage(
	useCase contract.ImageUploader,
) *UploadImage {
	return &UploadImage{
		useCase: useCase,
	}
}

func (h *UploadImage) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get file from request"})
		return
	}
	defer file.Close()

	imageID, err := h.useCase.UploadImage(token, file, header.Filename)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(200, gin.H{"id": imageID})
}
