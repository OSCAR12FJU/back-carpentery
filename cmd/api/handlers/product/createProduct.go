package product

import (
	"backend-carpintery/internal/domain"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {
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

	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	serverURL := "http://localhost:8080"
	imageURL := fmt.Sprintf("%s/%s", serverURL, filePath)

	productCard := &domain.Product{
		Title:    title,
		Price:    price,
		ImageURL: imageURL,
	}
	newProduct, err := h.ProductService.CreateProduct(*productCard)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error al cargar la imagen"})
		return
	}
	c.JSON(200, gin.H{"Product": newProduct})

}
