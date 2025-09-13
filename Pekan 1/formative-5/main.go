package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// --- Soal 1 ---

// Gunakan konstanta untuk menghindari 'magic strings' dan potensi typo.
const (
	GenderPria   = "laki-laki"
	GenderWanita = "perempuan"
)

// Best Practice: Hindari 'named return value' untuk fungsi sederhana.
// Menggunakan return eksplisit lebih jelas dan mudah dibaca.
func introduce(name, gender, profession, age string) string {
	var title string
	// Switch lebih rapi untuk perbandingan dengan beberapa nilai tetap.
	switch gender {
	case GenderPria:
		title = "Pak"
	case GenderWanita:
		title = "Bu"
	default:
		// Menangani kasus jika gender tidak valid
		title = ""
	}
	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, profession, age)
}

// --- Soal 2 ---

// Kode sudah cukup baik, tidak ada perubahan signifikan yang diperlukan.
func buahFavorit(name string, fruits ...string) string {
	buahStr := strings.Join(fruits, `", "`)
	return fmt.Sprintf(`halo nama saya %s dan buah favorit saya adalah "%s"`, name, buahStr)
}

// --- Soal 3 ---

// Best Practice: Gunakan 'struct' untuk data yang terstruktur.
// Ini memberikan type safety dan lebih jelas daripada map[string]string.
type Film struct {
	Judul  string
	Durasi string
	Genre  string
	Tahun  string
}

// Best Practice: Hindari variabel global.
// Gunakan closure untuk mengenkapsulasi state (data).
// Fungsi ini bertindak sebagai 'factory' yang membuat dan mengembalikan closure.
func filmManager() (func(string, string, string, string), func() []Film) {
	// 'dataFilm' sekarang menjadi variabel lokal yang hanya bisa diakses
	// oleh closure yang dikembalikan, bukan variabel global.
	var dataFilm = []Film{}

	// Closure untuk menambah data film
	tambahDataFilm := func(judul, durasi, genre, tahun string) {
		filmBaru := Film{
			Judul:  judul,
			Durasi: durasi,
			Genre:  genre,
			Tahun:  tahun,
		}
		dataFilm = append(dataFilm, filmBaru)
	}

	// Closure untuk mengambil data film
	getDataFilm := func() []Film {
		return dataFilm
	}

	return tambahDataFilm, getDataFilm
}

// --- Soal 4 ---

// Best Practice: Tambahkan komentar untuk menjelaskan kenapa menggunakan big.Int.
// function factorial untuk menghitung nilai faktorial dengan angka besar
// menggunakan big.Int untuk mencegah integer overflow.
func factorial(n int64) *big.Int {
	if n < 0 {
		// Faktorial tidak terdefinisi untuk angka negatif.
		return big.NewInt(0)
	}

	result := big.NewInt(1)
	if n == 0 {
		return result
	}

	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

// --- Soal 5 ---

// Best Practice: Cara paling idiomatik di Go adalah mengembalikan nilai, bukan menggunakan pointer sebagai parameter output.
// Fungsi ini lebih jelas menunjukkan input dan outputnya.
func hitungLingkaran(jariJari float64) (float64, float64) {
	luas := math.Pi * math.Pow(jariJari, 2)
	keliling := 2 * math.Pi * jariJari
	return luas, keliling
}

func main() {
	// --- Panggilan Soal 1 ---
	fmt.Println("--- Soal 1 ---")
	john := introduce("John", GenderPria, "penulis", "30")
	fmt.Println(john)
	sarah := introduce("Sarah", GenderWanita, "model", "28")
	fmt.Println(sarah)
	fmt.Println()

	// --- Panggilan Soal 2 ---
	fmt.Println("--- Soal 2 ---")
	buah := []string{"semangka", "jeruk", "melon", "pepaya"}
	pesanBuah := buahFavorit("John", buah...)
	fmt.Println(pesanBuah)
	fmt.Println()

	// --- Panggilan Soal 3 ---
	fmt.Println("--- Soal 3 ---")
	// Panggil 'factory' untuk mendapatkan fungsi penambah dan pengambil data.
	tambahDataFilm, getDataFilm := filmManager()

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	// Gunakan fungsi pengambil data untuk mencetak hasilnya.
	allFilms := getDataFilm()
	for _, item := range allFilms {
		fmt.Printf("%+v\n", item) // %+v mencetak nama field dan valuenya
	}
	fmt.Println()

	// --- Panggilan Soal 4 ---
	fmt.Println("--- Soal 4 ---")
	fmt.Printf("Factorial dari 5 adalah: %d\n", factorial(5))
	fmt.Printf("Factorial dari 7 adalah: %d\n", factorial(7))
	fmt.Printf("Factorial dari 20 adalah: %d\n", factorial(20))
	fmt.Println()

	// --- Panggilan Soal 5 ---
	fmt.Println("--- Soal 5 ---")
	// Gunakan fungsi yang mengembalikan nilai secara langsung.
	luas, keliling := hitungLingkaran(7)
	fmt.Printf("Jari-jari 7: Luas=%.2f, Keliling=%.2f\n", luas, keliling)

	luas, keliling = hitungLingkaran(10)
	fmt.Printf("Jari-jari 10: Luas=%.2f, Keliling=%.2f\n", luas, keliling)
}
