package main

import (
	"go-bioskop-api/config"
	"go-bioskop-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDatabase()

	// Inisialisasi router Gin
	r := gin.Default()

	// Membuat endpoint
	r.POST("/bioskop", controllers.CreateBioskop)      // Create
	r.GET("/bioskop", controllers.GetAllBioskops)      // Read (All)
	r.GET("/bioskop/:id", controllers.GetBioskopById)  // Read (by ID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)   // Update
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)// Delete

	// Menjalankan server di port 8080
	r.Run(":8080")
}