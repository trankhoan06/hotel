package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	TokenProvider "main.go/component"
	"main.go/modules/user/biz"
	"main.go/modules/user/storage"
	"net/http"
	"strconv"
)

func VerifyEmail(db *gorm.DB, provider TokenProvider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		code, errCode := strconv.Atoi(c.Query("code"))
		if errCode != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errCode.Error()})
		}
		token := c.Query("token")
		store := storage.NewSqlModel(db)
		hasher := common.NewSha265Hash()
		business := biz.NewLoginUser(store, provider, hasher)
		token1, err := business.VerifyCode(c.Request.Context(), token, userId, 60*60*24*30, code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token1})
	}
}
