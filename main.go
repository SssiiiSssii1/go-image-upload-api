package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()

	router.POST("/images", uploadImg)
	router.GET("/images/:filename", getImg)
	router.Run(":3000")
}

func getImg(c *gin.Context) {
	filename := c.Param("filename")

	c.File("uploads/" + filename)
}

func uploadImg(c *gin.Context) {
	f, err := c.FormFile("img")

	if err != nil {
		c.JSON(400, gin.H{
			"error": "image is required",
		})
		return
	}

	contentType := f.Header.Get("Content-Type")

	if contentType != "image/jpeg" &&
		contentType != "image/png" &&
		contentType != "image/webp" {
		c.JSON(400, gin.H{
			"error": "only jpeg, png and webp images are allowed",
		})
		return
	}

	const maxSize = 5 * 1024 * 1024

	if f.Size > maxSize {
		c.JSON(400, gin.H{
			"error": "image is too large",
		})
		return
	}

	ext := filepath.Ext(f.Filename)
	name := uuid.New().String() + ext

	err = c.SaveUploadedFile(f, "uploads/"+name)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to save image",
		})
		return
	}

	c.JSON(201, gin.H{
		"filename": name,
		"size":     f.Size,
	})
}
