package controllers

import (
	"errors"
	"go-bioskop-api/config"
	"go-bioskop-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Menggunakan satu struct input untuk Create dan Update agar tidak duplikasi.
type BioskopInput struct {
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float32 `json:"rating"`
}

// POST /bioskop -> Membuat bioskop baru
func CreateBioskop(c *gin.Context) {
	var input BioskopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bioskop := models.Bioskop{
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	// Simpan ke database
	if err := config.DB.Create(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": bioskop})
}

// GET /bioskop -> Menampilkan semua data bioskop
func GetAllBioskops(c *gin.Context) {
	var bioskops []models.Bioskop
	config.DB.Find(&bioskops)

	c.JSON(http.StatusOK, gin.H{"data": bioskops})
}

// GET /bioskop/:id -> Menampilkan data bioskop berdasarkan ID
func GetBioskopById(c *gin.Context) {
	var bioskop models.Bioskop

	// Cari data berdasarkan ID, jika tidak ada, return error.
	// Menggunakan errors.Is untuk penanganan error yang lebih spesifik.
	if err := config.DB.First(&bioskop, c.Param("id")).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bioskop})
}

// PUT /bioskop/:id -> Memperbarui data bioskop berdasarkan ID
func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := config.DB.First(&bioskop, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	var input BioskopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data di database
	config.DB.Model(&bioskop).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": bioskop})
}

// DELETE /bioskop/:id -> Menghapus data bioskop berdasarkan ID
func DeleteBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := config.DB.First(&bioskop, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	// Hapus data dari database
	config.DB.Delete(&bioskop)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}