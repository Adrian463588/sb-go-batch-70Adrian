// File: latihan/latihan.go

package latihan

import (
	"errors"
	"fmt"
	"math"
	"strings"

	// Menggunakan import path berbasis modul
	"FORMATIVE-9/matematika"
)

// --- Soal 1 ---
// Interface, Struct, dan Method sudah public (diawali huruf kapital).
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi))
}

// --- Soal 2 ---
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneDisplayer interface {
	ShowData() string
}

func NewPhone(name, brand string, year int, colors []string) Phone {
	return Phone{
		Name:   name,
		Brand:  brand,
		Year:   year,
		Colors: colors,
	}
}

func (p Phone) ShowData() string {
	colorsStr := strings.Join(p.Colors, ", ")
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: [%s]", p.Name, p.Brand, p.Year, colorsStr)
}

// --- Soal 3 ---
// Fungsi ini tidak perlu di-export karena hanya digunakan oleh JalankanSoal3
func luasPersegiSesuaiSoal(sisi int, withText bool) interface{} {
	if sisi <= 0 {
		if withText {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}
	luas := sisi * sisi
	if withText {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

func luasPersegiIdiomatic(sisi int) (string, error) {
	if sisi <= 0 {
		return "", errors.New("sisi harus lebih besar dari 0")
	}
	luas := sisi * sisi
	return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas), nil
}

// --- Fungsi Publik untuk Setiap Soal ---

// JalankanSoal1 diekspor agar bisa dipanggil dari package main.
func JalankanSoal1() {
	fmt.Println("--- Soal 1 ---")
	var bangunDatar HitungBangunDatar = SegitigaSamaSisi{Alas: 6, Tinggi: 4}
	fmt.Println("Segitiga Sama Sisi")
	fmt.Printf("Luas: %d, Keliling: %d\n", bangunDatar.Luas(), bangunDatar.Keliling())

	bangunDatar = PersegiPanjang{Panjang: 8, Lebar: 5}
	fmt.Println("\nPersegi Panjang")
	fmt.Printf("Luas: %d, Keliling: %d\n", bangunDatar.Luas(), bangunDatar.Keliling())

	var bangunRuang HitungBangunRuang = Tabung{JariJari: 5, Tinggi: 10}
	fmt.Println("\nTabung")
	fmt.Printf("Volume: %.2f, Luas Permukaan: %.2f\n", bangunRuang.Volume(), bangunRuang.LuasPermukaan())

	bangunRuang = Balok{Panjang: 6, Lebar: 4, Tinggi: 5}
	fmt.Println("\nBalok")
	fmt.Printf("Volume: %.2f, Luas Permukaan: %.2f\n", bangunRuang.Volume(), bangunRuang.LuasPermukaan())
}

func JalankanSoal2() {
	fmt.Println("--- Soal 2 ---")
	myPhone := NewPhone(
		"iPhone 15 Pro",
		"Apple",
		2023,
		[]string{"Natural Titanium", "Blue Titanium"},
	)
	var displayer PhoneDisplayer = myPhone
	fmt.Println(displayer.ShowData())
}

func JalankanSoal3() {
	fmt.Println("--- Soal 3 (Sesuai Soal) ---")
	fmt.Println(luasPersegiSesuaiSoal(4, true))
	fmt.Println(luasPersegiSesuaiSoal(8, false))
	fmt.Println(luasPersegiSesuaiSoal(0, true))
	fmt.Println(luasPersegiSesuaiSoal(0, false))

	fmt.Println("\n--- Soal 3 (Cara Idiomatic) ---")
	if result, err := luasPersegiIdiomatic(4); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	if _, err := luasPersegiIdiomatic(0); err != nil {
		fmt.Println("Error:", err)
	}
}

func JalankanSoal4() {
	fmt.Println("--- Soal 4 ---")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixStr, ok1 := prefix.(string)
	angka1, ok2 := kumpulanAngkaPertama.([]int)
	angka2, ok3 := kumpulanAngkaKedua.([]int)

	if !ok1 || !ok2 || !ok3 {
		fmt.Println("Gagal melakukan type assertion, tipe data tidak sesuai.")
		return
	}

	semuaAngka := append(angka1, angka2...)
	total := 0
	var angkaStr []string
	for _, angka := range semuaAngka {
		total += angka
		angkaStr = append(angkaStr, fmt.Sprint(angka))
	}
	fmt.Printf("%s%s = %d\n", prefixStr, strings.Join(angkaStr, " + "), total)
}

func JalankanSoal5() {
	fmt.Println("--- Soal 5 ---")
	hasilTambah := matematika.Tambah(10, 5)
	hasilKali := matematika.Kali(10, 5)
	fmt.Printf("Hasil Penjumlahan: %d\n", hasilTambah)
	fmt.Printf("Hasil Perkalian: %d\n", hasilKali)
}