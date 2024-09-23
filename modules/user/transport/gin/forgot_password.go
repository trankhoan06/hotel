package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	TokenProvider "main.go/component"
	"main.go/modules/user/biz"
	"main.go/modules/user/storage"
	"net/http"
)

func ForgotPassword(db *gorm.DB, provider TokenProvider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Query("email")
		hash := common.NewSha265Hash()
		store := storage.NewSqlModel(db)
		business := biz.NewLoginUser(store, provider, hash)
		Eto, err := business.ForgotPassword(c.Request.Context(), email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": Eto})
	}
}
