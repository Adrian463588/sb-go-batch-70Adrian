package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// helper: kapitalisasi huruf pertama
func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}

// helper: kapitalisasi huruf terakhir
func capitalizeLast(s string) string {
	if s == "" {
		return s
	}
	rs := []rune(s)
	rs[len(rs)-1] = unicode.ToUpper(rs[len(rs)-1])
	return string(rs)
}

func main() {
	// soal 1
	// tampilkan "Bootcamp Digital Skill Sanbercode Golang" dari 5 variabel terpisah
	k1 := "Bootcamp"
	k2 := "Digital"
	k3 := "Skill"
	k4 := "Sanbercode"
	k5 := "Golang"
	fmt.Println(k1, k2, k3, k4, k5)

	// soal 2
	// ganti kata "Dunia" jadi "Golang" memakai package strings
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)

	// soal 3
	// hasil akhir: "saya Senang belajaR GOLANG"
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Println(
		kataPertama,
		capitalizeFirst(kataKedua),     // "Senang"
		capitalizeLast(kataKetiga),     // "belajaR"
		strings.ToUpper(kataKeempat),   // "GOLANG"
	)

	// soal 4
	// ubah ke integer lalu jumlahkan
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	a1, _ := strconv.Atoi(angkaPertama)
	a2, _ := strconv.Atoi(angkaKedua)
	a3, _ := strconv.Atoi(angkaKetiga)
	a4, _ := strconv.Atoi(angkaKeempat)
	fmt.Println(a1 + a2 + a3 + a4)

	// soal 5 (materi hari berikutnya)
	// hitung: luas & keliling persegi panjang, dan luas segitiga
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	p, _ := strconv.Atoi(panjangPersegiPanjang)
	l, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = p * l
	var kelilingPersegiPanjang int = 2 * (p + l)
	var luasSegitiga int = (alas * tinggi) / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)
}
