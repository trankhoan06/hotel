package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main.go/component/jwt"
	"main.go/component/middleware"
	"main.go/modules/user/storage"
	ginUser "main.go/modules/user/transport/gin"
	"os"
)

func main() {

	dsn := os.Getenv("DOMAIN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	token := jwt.NewJwtProvider("jwt", "Khoandz123@")
	author := storage.NewSqlModel(db)
	middlewareAuthor := middleware.RequestAuthenrize(author, token)
	r := gin.Default()
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	{
		v1.POST("/register", ginUser.Register(db))
		v1.GET("/profile", ginUser.FindUser(db))
		v1.POST("/login", ginUser.Login(db, token))
		v1.PATCH("/verify_email", ginUser.VerifyEmail(db, token))
		v1.POST("/verify_code", ginUser.VerifyCode(db))
		v1.POST("/forgot_password", ginUser.ForgotPassword(db, token))
		v1.PATCH("/change_password", middlewareAuthor, ginUser.ChangePassword(db))
		v1.PATCH("/change_password_forgot", ginUser.ChangePasswordForget(db))
	}
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
