package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main.go/component/jwt"
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
	//author := storage.NewSqlModel(db)
	//middlewareAuthor := middleware.RequestAuthenrize(author, token)
	r := gin.Default()
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	{
		v1.POST("/register", ginUser.Register(db))
		v1.POST("/login", ginUser.Login(db, token))
		v1.PATCH("/verify_email", ginUser.VerifyEmail(db, token))
	}
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
