package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	TokenProvider "main.go/component"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func Login(db *gorm.DB, provider TokenProvider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.Login
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha265Hash()
		business := biz.NewLoginUser(store, provider, hash)
		token, errLogin := business.NewLogin(c.Request.Context(), &data, 60*60*24*30)
		if errLogin != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errLogin.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token.Token, "is_email": token.IsEmail})
	}
}
