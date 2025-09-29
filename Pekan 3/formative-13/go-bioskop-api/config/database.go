package config

import (
	"fmt"
	"go-bioskop-api/models" // Import package models

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ganti dengan detail koneksi databasemu
	dsn := "host=localhost user=postgres password=postgres dbname=bioskop_db port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal terkoneksi ke database!")
	}

	fmt.Println("Koneksi ke database berhasil.")

	// AutoMigrate akan membuat tabel 'bioskops' berdasarkan struct models.Bioskop
	DB.AutoMigrate(&models.Bioskop{})
}