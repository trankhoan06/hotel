package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"net/http"
)

func Profile(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrUser)
		c.JSON(http.StatusOK, gin.H{"user": user})

	}
}
