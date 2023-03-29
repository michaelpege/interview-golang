package service

// func TestCreateBook(t *testing.T) {
// 	router := gin.Default()
// 	dsn := "root:@tcp(127.0.0.1:3306)/interview?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	db.AutoMigrate(&models.Book{})

// 	router.POST("/books", func(c *gin.Context) {
// 		CreateBook(c, db)
// 	})

// }
