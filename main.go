package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	middlewareAuthor := middleware.RequesMiddleware(author, token)
	r := gin.Default()
	r.Static("/static", "./static")

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
