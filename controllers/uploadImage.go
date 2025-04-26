package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	filename := filepath.Base(file.Filename)
	uniqueFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), filename)

	// Путь для сохранения
	savePath := filepath.Join("./images", uniqueFilename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	imageURL := uniqueFilename

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}
