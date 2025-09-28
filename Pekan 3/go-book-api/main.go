
package main

import (
    "golang-book-api/database"
    "golang-book-api/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()

    router := gin.Default()

    // Grup API
    api := router.Group("/api")
    {
        // Endpoint Buku
        api.GET("/books", handler.GetAllBooks)
        api.GET("/books/:id", handler.GetBookDetail)
        api.POST("/books", handler.CreateBook)
        api.DELETE("/books/:id", handler.DeleteBook)

        // Endpoint Kategori
        api.GET("/categories", handler.GetAllCategories)
        api.GET("/categories/:id", handler.GetCategoryDetail)
        api.POST("/categories", handler.CreateCategory)
        api.DELETE("/categories/:id", handler.DeleteCategory)
        api.GET("/categories/:id/books", handler.GetBooksByCategory)
    }

    router.Run(":8080")
}