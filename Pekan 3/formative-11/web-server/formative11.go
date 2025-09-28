package main

import (
	"fmt"
	"math"
	"net/http"
)

// Fungsi handler untuk menghitung dan menampilkan data tabung
func cylinderHandler(w http.ResponseWriter, r *http.Request) {
	// Deklarasi konstanta dan variabel
	const jariJari float64 = 7
	const tinggi float64 = 10
	const pi float64 = math.Pi

	// Perhitungan
	luasAlas := pi * math.Pow(jariJari, 2)
	kelilingAlas := 2 * pi * jariJari
	volume := luasAlas * tinggi

	// Membuat response text
	responseText := fmt.Sprintf(
		"jariJari : %.0f, tinggi: %.0f, volume : %.2f, luas alas: %.2f, keliling alas: %.2f",
		jariJari,
		tinggi,
		volume,
		luasAlas,
		kelilingAlas,
	)

	// Mengirim response ke client
	fmt.Fprint(w, responseText)
}

func main() {
	// Mendaftarkan handler untuk root URL ("/")
	http.HandleFunc("/", cylinderHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	// Menjalankan web server di port 8080
	http.ListenAndServe(":8080", nil)
}