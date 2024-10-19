package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func UploadImage(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		img := Image{
			Url:       dst,
			Width:     200,
			Height:    200,
			Cloud:     "local",
			Extension: "",
		}
		img.Fullfill("http://localhost:3000")
		c.JSON(http.StatusOK, gin.H{"image": img})
	}
}
