package service

import (
	"interview/internal/models"
	"interview/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBook(c *gin.Context, db *gorm.DB) {
	var bookPayload models.Book

	if err := c.ShouldBindJSON(&bookPayload); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	newBook := models.Book{
		ID:     uuid.New().String(),
		Title:  bookPayload.Title,
		Author: bookPayload.Author,
	}

	db.Create(newBook)
	utils.SendResponse(c, http.StatusCreated, "Book Created", newBook)
}

func GetBooks(c *gin.Context, db *gorm.DB) {
	var books []models.Book

	results := db.Find(&books)

	if results.Error != nil {
		panic(results.Error)
	}

	utils.SendResponse(c, http.StatusOK, "Book Data", books)
}
