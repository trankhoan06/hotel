package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"main.go/common"
	TokenProvider "main.go/component"
	"main.go/modules/user/model"
	"net/http"
	"strings"
)

type AuthenStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

func ExtractToken(token string) (string, error) {
	t := strings.Split(token, " ")
	if t[0] != "Bearer" || len(t) < 2 || strings.TrimSpace(t[1]) == "" {
		return "", errors.New("token has been fault")
	}
	return t[1], nil
}
func RequestAuthenrize(auth AuthenStorage, provider TokenProvider.Provider) func(*gin.Context) {

	return func(c *gin.Context) {
		token, err := ExtractToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		payload, err := provider.Validate(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		user, err := auth.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.GetUser()})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		c.Set(common.CurrUser, user)
		c.Next()
	}

}
