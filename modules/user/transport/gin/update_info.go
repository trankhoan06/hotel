package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func UpdateInfo(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateUser
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		data.Email = c.MustGet(common.CurrUser).(common.Requester).GetEmail()
		if err := store.UpdateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
