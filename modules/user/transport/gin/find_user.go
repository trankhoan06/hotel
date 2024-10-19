package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/user/storage"
	"net/http"
	"strconv"
)

func FindUser(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": userId})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})

	}
}
