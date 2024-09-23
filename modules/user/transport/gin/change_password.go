package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/user/biz"
	"main.go/modules/user/storage"
	"net/http"
)

func ChangePassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		password := c.Query("password")
		request := c.MustGet(common.CurrUser).(common.Requester)
		hash := common.NewSha265Hash()
		store := storage.NewSqlModel(db)
		business := biz.NewRegisterUser(store, hash)
		if err := business.NewChangePassword(c.Request.Context(), password, request.GetEmail()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
