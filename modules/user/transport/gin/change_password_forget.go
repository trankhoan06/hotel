package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func ChangePasswordForget(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var forgot model.ForgotPassword
		if err := c.ShouldBindJSON(&forgot); err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha265Hash()
		business := biz.NewRegisterUser(store, hash)
		if err := business.ChangePasswordForgot(c.Request.Context(), forgot.Token, forgot.Password, forgot.UserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
