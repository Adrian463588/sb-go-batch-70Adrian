package models

// Bioskop merepresentasikan model data untuk sebuah bioskop
type Bioskop struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}