package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct untuk data nilai mahasiswa yang akan disimpan
type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Nilai       uint   `json:"nilai"`
	IndeksNilai string `json:"indeks_nilai"`
}

// Struct untuk menerima input JSON dari client (hanya 3 field)
type NilaiInput struct {
	Nama       string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	Nilai      uint   `json:"nilai"`
}

// Variabel global sebagai database in-memory
var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// Middleware untuk Basic Authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic Auth hanya diwajibkan untuk method POST
		if r.Method == http.MethodPost {
			user, pass, ok := r.BasicAuth()
			// Hardcoded username: "admin", password: "password123"
			if !ok || user != "admin" || pass != "password123" {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
				return // Hentikan proses jika otentikasi gagal
			}
		}
		// Lanjutkan ke handler utama jika otentikasi berhasil atau method bukan POST
		next.ServeHTTP(w, r)
	})
}

// Handler utama untuk endpoint /nilai
func nilaiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// soal 2: API GET Nilai Mahasiswa
		// Menampilkan semua data nilai mahasiswa yang sudah ada
		if len(nilaiNilaiMahasiswa) == 0 {
			fmt.Fprint(w, "[]") // Tampilkan array JSON kosong jika tidak ada data
			return
		}

		err := json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case http.MethodPost:
		// soal 1: API POST Nilai Mahasiswa
		var input NilaiInput

		// Decode JSON body dari request ke struct NilaiInput
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validasi: Nilai tidak boleh lebih dari 100
		if input.Nilai > 100 {
			http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
			return
		}

		// Kalkulasi IndeksNilai berdasarkan kondisi
		var indeks string
		switch {
		case input.Nilai >= 80:
			indeks = "A"
		case input.Nilai >= 70:
			indeks = "B"
		case input.Nilai >= 60:
			indeks = "C"
		case input.Nilai >= 50:
			indeks = "D"
		default:
			indeks = "E"
		}

		// Generate ID baru (berdasarkan jumlah data yang sudah ada + 1)
		newID := uint(len(nilaiNilaiMahasiswa) + 1)

		// Buat struct NilaiMahasiswa yang lengkap
		newData := NilaiMahasiswa{
			ID:          newID,
			Nama:        input.Nama,
			MataKuliah:  input.MataKuliah,
			Nilai:       input.Nilai,
			IndeksNilai: indeks,
		}

		// Tambahkan data baru ke slice global
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newData)

		// Kirim response status 201 Created dan data yang baru saja dibuat
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newData)

	default:
		// Jika method bukan GET atau POST, kirim error
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Buat handler dari fungsi nilaiHandler
	mainHandler := http.HandlerFunc(nilaiHandler)

	// Terapkan middleware Auth ke handler utama
	http.Handle("/nilai", AuthMiddleware(mainHandler))

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}