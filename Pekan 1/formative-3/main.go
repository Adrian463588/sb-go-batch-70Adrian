package main

import (
	"fmt"
	"strings"
)

func main() {
	// ==================================
	// soal 1
	// olah variabel sehingga output akhir: `"Hi Hi bandung" - 2021`
	// ==================================
	kalimat := "halo halo bandung"
	angka := 2021

	// ganti "halo" menjadi "Hi" pada dua kemunculan pertama
	hasilKalimat := strings.Replace(kalimat, "halo", "Hi", 2)
	// tampilkan string dengan tanda kutip dan angka seperti contoh
	fmt.Printf("%q - %d\n\n", hasilKalimat, angka)

	// ==================================
	// soal 2
	// tentukan indeks nilai untuk John dan Doe sesuai ketentuan
	// ==================================
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println("Indeks nilai John:", indeksNilai(nilaiJohn))
	fmt.Println("Indeks nilai Doe :", indeksNilai(nilaiDoe))
	fmt.Println()

	// ==================================
	// soal 3
	// ubah tanggal/bulan/tahun sesuai tanggal lahir Anda
	// lalu cetak dalam format: 17 Agustus 1945
	// ==================================
	var tanggal = 17
	var bulan = 8
	var tahun = 1945
	// NOTE: silakan ganti tiga variabel di atas sesuai tanggal lahir Anda

	fmt.Printf("%d %s %d\n\n", tanggal, namaBulan(bulan), tahun)

	// ==================================
	// soal 4
	// tentukan istilah generasi berdasarkan tahun kelahiran
	// ==================================
	fmt.Printf("Tahun %d termasuk: %s\n\n", tahun, istilahGenerasi(tahun))

	// ==================================
	// soal 5
	// looping 1..20 dengan syarat:
	//  A. ganjil  -> Santai
	//  B. genap   -> Berkualitas
	//  C. kelipatan 3 dan ganjil -> I Love Coding
	// Urutan pengecekan penting agar aturan C diutamakan.
	// ==================================
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 == 1 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 1 {
			fmt.Println(i, "- Santai")
		} else {
			fmt.Println(i, "- Berkualitas")
		}
	}
}

// indeksNilai mengembalikan huruf indeks berdasarkan skor
func indeksNilai(n int) string {
	switch {
	case n >= 80:
		return "A"
	case n >= 70 && n < 80:
		return "B"
	case n >= 60 && n < 70:
		return "C"
	case n >= 50 && n < 60:
		return "D"
	default:
		return "E"
	}
}

// namaBulan mengubah angka bulan (1..12) menjadi nama bulan bahasa Indonesia
func namaBulan(b int) string {
	switch b {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	default:
		return "Bulan tidak dikenal"
	}
}

// istilahGenerasi menentukan istilah generasi berdasarkan tahun kelahiran
func istilahGenerasi(tahun int) string {
	switch {
	case tahun >= 1944 && tahun <= 1964:
		return "Baby boomer"
	case tahun >= 1965 && tahun <= 1979:
		return "Generasi X"
	case tahun >= 1980 && tahun <= 1994:
		return "Generasi Y (Millennials)"
	case tahun >= 1995 && tahun <= 2015:
		return "Generasi Z"
	default:
		return "Di luar rentang yang didefinisikan"
	}
}
