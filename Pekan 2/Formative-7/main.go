package main

import "fmt"

// --- SOAL 1 ---
// Best Practice: Nama struct dan field diawali huruf kapital (PascalCase) agar exported.
type Mahasiswa struct {
	Nama string
	NIM  string
	Usia int
}

// --- SOAL 2 ---
// Best Practice: Dibuat interface `Bentuk` untuk abstraksi.
type Bentuk interface {
	Luas() float64
}

// Nama struct dan field diubah menjadi PascalCase.
type Segitiga struct {
	Alas, Tinggi int
}

// Method Luas() untuk Segitiga (mengimplementasikan interface Bentuk).
func (s Segitiga) Luas() float64 {
	return 0.5 * float64(s.Alas) * float64(s.Tinggi)
}

type Persegi struct {
	Sisi int
}

// Method Luas() untuk Persegi.
func (p Persegi) Luas() float64 {
	return float64(p.Sisi * p.Sisi)
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

// Method Luas() untuk PersegiPanjang.
func (pp PersegiPanjang) Luas() float64 {
	return float64(pp.Panjang * pp.Lebar)
}

// --- SOAL 3 ---
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

// Best Practice: Membuat fungsi konstruktor untuk inisialisasi struct.
func NewPhone(name, brand string, year int) *Phone {
	return &Phone{
		Name:   name,
		Brand:  brand,
		Year:   year,
		Colors: []string{}, // Slice diinisialisasi agar tidak nil
	}
}

// Method TambahWarna (pointer receiver sudah merupakan best practice di sini).
func (p *Phone) TambahWarna(warnaBaru string) {
	p.Colors = append(p.Colors, warnaBaru)
}

// --- SOAL 4 ---
type Movie struct {
	Title    string
	Genre    string
	Duration int // dalam menit
	Year     int
}

// Best Practice: Mengembalikan slice baru, bukan memodifikasi via pointer.
// Ini membuat fungsi lebih fungsional dan mengikuti pola `append`.
func tambahDataFilm(dataFilm []Movie, title string, duration int, genre string, year int) []Movie {
	filmBaru := Movie{
		Title:    title,
		Duration: duration,
		Genre:    genre,
		Year:     year,
	}
	return append(dataFilm, filmBaru)
}

// --- SOAL 5 ---
type Hewan interface {
	Suara() string
}

type Kucing struct{}

func (k Kucing) Suara() string {
	return "Meow!"
}

type Anjing struct{}

func (a Anjing) Suara() string {
	return "Guk Guk!"
}

// Fungsi ini menunjukkan kekuatan interface sebagai parameter.
func cetakSuara(h Hewan) {
	fmt.Printf("Hewan ini bersuara: %s\n", h.Suara())
}

func main() {
	// soal 1
	fmt.Println("--- Soal 1 ---")
	mhs := Mahasiswa{
		Nama: "Budi",
		NIM:  "12345678",
		Usia: 20,
	}
	fmt.Printf("Nama: %s, NIM: %s, Usia: %d tahun\n", mhs.Nama, mhs.NIM, mhs.Usia)
	fmt.Println()

	// soal 2
	fmt.Println("--- Soal 2 ---")
	// Semua bentuk dapat disimpan dalam satu slice `Bentuk`.
	daftarBentuk := []Bentuk{
		Segitiga{Alas: 10, Tinggi: 5},
		Persegi{Sisi: 4},
		PersegiPanjang{Panjang: 8, Lebar: 5},
	}
	for _, bentuk := range daftarBentuk {
		fmt.Printf("Luas: %.2f\n", bentuk.Luas())
	}
	fmt.Println()

	// soal 3
	fmt.Println("--- Soal 3 ---")
	myPhone := NewPhone("iPhone 15", "Apple", 2023)
	fmt.Printf("Warna awal: %v\n", myPhone.Colors)

	myPhone.TambahWarna("Blue")
	myPhone.TambahWarna("Titanium")
	fmt.Printf("Warna setelah ditambah: %v\n", myPhone.Colors)
	fmt.Println()

	// soal 4
	fmt.Println("--- Soal 4 ---")
	var dataFilm = []Movie{}

	// Menangkap hasil return dari fungsi untuk memperbarui slice.
	dataFilm = tambahDataFilm(dataFilm, "LOTR", 178, "action", 1999)
	dataFilm = tambahDataFilm(dataFilm, "avenger", 181, "action", 2019)
	dataFilm = tambahDataFilm(dataFilm, "spiderman", 121, "action", 2004)
	dataFilm = tambahDataFilm(dataFilm, "juon", 92, "horror", 2004)

	for i, film := range dataFilm {
		fmt.Printf("%d.\n", i+1)
		fmt.Printf("   Title    : %s\n", film.Title)
		fmt.Printf("   Duration : %d minutes\n", film.Duration)
		fmt.Printf("   Genre    : %s\n", film.Genre)
		fmt.Printf("   Year     : %d\n", film.Year)
		fmt.Println()
	}

	// soal 5
	fmt.Println("--- Soal 5 ---")
	kucing := Kucing{}
	anjing := Anjing{}
	
	cetakSuara(kucing)
	cetakSuara(anjing)
}