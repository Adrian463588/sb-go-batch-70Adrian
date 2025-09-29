package controllers

import (
	"go-bioskop-api/config"
	"go-bioskop-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Definisikan struct untuk input agar bisa divalidasi
type CreateBioskopInput struct {
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float32 `json:"rating"`
}

// POST /bioskop
// Menambahkan bioskop baru
func CreateBioskop(c *gin.Context) {
	var input CreateBioskopInput

	// Validasi input JSON
	// Jika ada field 'nama' atau 'lokasi' yang kosong, akan return error
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat objek bioskop baru
	bioskop := models.Bioskop{
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	// Simpan ke database
	config.DB.Create(&bioskop)

	// Kembalikan data yang baru dibuat sebagai response
	c.JSON(http.StatusCreated, gin.H{"data": bioskop})
}