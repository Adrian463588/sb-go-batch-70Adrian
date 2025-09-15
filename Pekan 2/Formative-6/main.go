package main

import "fmt"

// --- SOAL 1 ---
// Best Practice: Mengembalikan string baru, bukan mengubah via pointer.
// Ini membuat fungsi menjadi lebih murni (pure function) dan alur data lebih jelas.
func introduce(name, gender, profession, age string) string {
	var title string
	if gender == "laki-laki" {
		title = "Pak"
	} else {
		title = "Bu"
	}
	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, profession, age)
}

// --- SOAL 2 ---
// Best Practice: Mengembalikan slice baru, mengikuti pola fungsi `append` bawaan Go.
func tambahBuah(listBuah []string) []string {
	buahToAdd := []string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}
	return append(listBuah, buahToAdd...)
}

// --- SOAL 3 ---
// Best Practice: Gunakan struct untuk data yang terstruktur agar lebih type-safe.
type Film struct {
	Title    string
	Duration string
	Genre    string
	Year     int // Menggunakan int untuk tahun lebih tepat secara tipe data
}

// Fungsi ini sekarang menerima slice of Film dan mengembalikan slice yang baru.
func tambahDataFilm(title, duration, genre string, year int, dataFilm []Film) []Film {
	filmBaru := Film{
		Title:    title,
		Duration: duration,
		Genre:    genre,
		Year:     year,
	}
	return append(dataFilm, filmBaru)
}

// --- SOAL 4 ---
// Best Practice: Gunakan slice `[]int` sebagai parameter agar lebih fleksibel daripada
// pointer ke array `*[5]int`. Modifikasi elemen slice akan mengubah data aslinya.
func gandakanNilai(numbers []int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}

// --- SOAL 5 ---
type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

// Best Practice: Implementasi interface fmt.Stringer dengan method String().
// Ini memungkinkan kita mengontrol bagaimana struct ini ditampilkan saat di-print.
func (b Buah) String() string {
	var statusBiji string
	if b.AdaBijinya {
		statusBiji = "berbiji"
	} else {
		statusBiji = "tanpa biji"
	}
	return fmt.Sprintf("Buah %s berwarna %s, harganya Rp%d (%s)", b.Nama, b.Warna, b.Harga, statusBiji)
}

func main() {
	// soal 1
	fmt.Println("--- Soal 1 ---")
	sentence := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	sentence = introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	fmt.Println()

	// soal 2
	fmt.Println("--- Soal 2 ---")
	var daftarBuah = []string{}
	daftarBuah = tambahBuah(daftarBuah) // Menangkap hasil return dari fungsi
	for i, item := range daftarBuah {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println()

	// soal 3
	fmt.Println("--- Soal 3 ---")
	var dataFilm = []Film{} // Menggunakan slice dari struct Film
	dataFilm = tambahDataFilm("LOTR", "2 jam", "action", 1999, dataFilm)
	dataFilm = tambahDataFilm("avenger", "2 jam", "action", 2019, dataFilm)
	dataFilm = tambahDataFilm("spiderman", "2 jam", "action", 2004, dataFilm)
	dataFilm = tambahDataFilm("juon", "2 jam", "horror", 2004, dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d.\n", i+1)
		fmt.Printf("   title    : %s\n", film.Title)    // Mengakses field struct
		fmt.Printf("   duration : %s\n", film.Duration) // Lebih aman dari typo
		fmt.Printf("   genre    : %s\n", film.Genre)
		fmt.Printf("   year     : %d\n", film.Year)
		fmt.Println()
	}

	// soal 4
	fmt.Println("--- Soal 4 ---")
	var angka = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array sebelum digandakan:", angka)

	gandakanNilai(angka[:]) // Mengirim array sebagai slice dengan notasi `[:]`
	fmt.Println("Array setelah digandakan: ", angka)
	fmt.Println()

	// soal 5
	// soal 5
	fmt.Println("--- Soal 5 ---")
	var koleksiBuah = []Buah{
		{Nama: "Nanas", Warna: "Kuning", AdaBijinya: false, Harga: 9000},
		{Nama: "Jeruk", Warna: "Oranye", AdaBijinya: true, Harga: 8000},
		{Nama: "Semangka", Warna: "Hijau & Merah", AdaBijinya: false, Harga: 10000},
		{Nama: "Pisang", Warna: "Kuning", AdaBijinya: true, Harga: 5000},
	}

	for _, b := range koleksiBuah {
		fmt.Println(b) // fmt.Println akan otomatis memanggil method String() yang sudah kita buat
	}
}