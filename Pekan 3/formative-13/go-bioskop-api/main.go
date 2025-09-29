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
	r.POST("/bioskop", controllers.CreateBioskop)

	// Menjalankan server di port 8080
	r.Run(":8080")
}