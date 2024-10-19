package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/user/biz"
	"main.go/modules/user/storage"
	"net/http"
	"strconv"
)

func VerifyCode(db *gorm.DB) func(*gin.Context) {
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
		business := biz.NewUserBiz(store)
		token1, err := business.NewVerifyCodeForgot(c.Request.Context(), code, token, userId, 30*60)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": token1})
	}
}
