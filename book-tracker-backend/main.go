package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/gin-contrib/cors"

)

var DB *gorm.DB

type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Status string `json:"status"`
}


func main() {
    // Database connection
    dsn := "root:Monik@123@tcp(127.0.0.1:3306)/booktracker?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&Book{})

    // Initialize Gin router
    router := gin.Default()

    // Routes will be added here

	router.POST("/books", createBook)
    router.GET("/books", getBooks)
    router.PUT("/books/:id", updateBook)
    router.DELETE("/books/:id", deleteBook)

    // Run the server
    router.Run(":8080")

    router.Use(cors.Default())

}

// Create a new book
func createBook(c *gin.Context) {
    var book Book
    if err := c.ShouldBindJSON(&book); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := DB.Create(&book).Error; err != nil {
        log.Printf("Error creating book: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}


// Get all books
func getBooks(c *gin.Context) {
    var books []Book
    DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

// Update a book's status
func updateBook(c *gin.Context) {
    id := c.Param("id")
    var book Book
    if err := DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

// Delete a book
func deleteBook(c *gin.Context) {
    id := c.Param("id")
    if err := DB.Delete(&Book{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
