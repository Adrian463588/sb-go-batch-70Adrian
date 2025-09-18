package main

import (
	"errors"
	"fmt"
	"math"
	"strings"

	// Best Practice: Menggunakan import path berbasis modul.
	// Jalankan 'go mod init nama-proyek' di terminal Anda.
	"FORMATIVE-8/matematika"
)

// --- Soal 1 ---
// Best Practice: Nama interface dan method diawali huruf kapital agar exported.
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

// Nama struct dan field juga diawali huruf kapital.
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

// Best Practice: Membuat fungsi konstruktor untuk inisialisasi yang konsisten.
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
// Versi asli yang diminta soal. Penggunaan `interface{}` kurang ideal karena menghilangkan type safety.
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

// Best Practice: Menggunakan multiple return values (string, error) untuk hasil yang jelas dan aman.
func luasPersegiIdiomatic(sisi int) (string, error) {
	if sisi <= 0 {
		return "", errors.New("sisi harus lebih besar dari 0")
	}
	luas := sisi * sisi
	return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas), nil
}


// Best Practice: Memecah `main` menjadi fungsi-fungsi yang lebih kecil.
func jalankanSoal1() {
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

func jalankanSoal2() {
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

func jalankanSoal3() {
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

func jalankanSoal4() {
	fmt.Println("--- Soal 4 ---")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Best Practice: Menggunakan "comma, ok" idiom untuk type assertion yang aman.
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

func jalankanSoal5() {
	fmt.Println("--- Soal 5 ---")
	hasilTambah := matematika.Tambah(10, 5)
	hasilKali := matematika.Kali(10, 5)

	fmt.Printf("Hasil Penjumlahan: %d\n", hasilTambah)
	fmt.Printf("Hasil Perkalian: %d\n", hasilKali)
}

func main() {
	jalankanSoal1()
	fmt.Println()
	jalankanSoal2()
	fmt.Println()
	jalankanSoal3()
	fmt.Println()
	jalankanSoal4()
	fmt.Println()
	jalankanSoal5()
}