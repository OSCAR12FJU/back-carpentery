package collage

import (
	"backend-carpintery/internal/domain"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCollage(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	title := c.Request.FormValue("title")

	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	defer file.Close()

	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}
	filePath := filepath.Join(uploadDir, handler.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	serverURL := "http://localhost:8080"
	imageURL := fmt.Sprintf("%s/%s", serverURL, filePath)

	collage := &domain.Collage{
		Title:    title,
		ImageURL: imageURL,
	}

	colId, err := h.CollageService.CreateCollage(*collage)
	if err != nil {
		c.JSON(400, gin.H{"error": "oops!"})
	}
	c.JSON(200, gin.H{"ID": colId})

}
