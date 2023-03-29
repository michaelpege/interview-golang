package main

import (
	"fmt"
	"net/http"

	"interview/internal/middleware"
	"interview/internal/models"
	"interview/internal/service"
	"interview/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ping(c *gin.Context) {
	utils.SendResponse(c, http.StatusOK, "pong", nil)
}

func main() {
	router := gin.Default()
	dsn := "root:@tcp(127.0.0.1:3306)/interview?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})
	fmt.Println("Database migrated")

	router.GET("/", ping)

	router.Use(middleware.KeyProtection())

	router.POST("/books", func(c *gin.Context) {
		service.CreateBook(c, db)
	})

	router.GET("/books", func(c *gin.Context) {
		service.GetBooks(c, db)
	})

	router.Run("localhost:5000")
	fmt.Println("Hello world")
}
