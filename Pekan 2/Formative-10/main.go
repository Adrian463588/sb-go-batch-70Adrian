package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// soal 1
// Function ini akan dieksekusi di akhir program karena dipanggil menggunakan 'defer'.
func displayInfo(kalimat string, tahun int) {
	fmt.Println("--- Output Soal 1 (dari defer) ---")
	fmt.Printf("Kalimat: %s, Tahun: %d\n", kalimat, tahun)
}

// soal 2
// Function ini menghitung keliling segitiga sama sisi dengan berbagai kondisi,
// termasuk penanganan error dan panic/recover.
func kelilingSegitigaSamaSisi(sisi int, cetakKalimat bool) (string, error) {
	// Defer function ini akan menangkap panic jika terjadi.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic berhasil di-recover:", r)
		}
	}()

	if sisi <= 0 {
		errMsg := "Maaf anda belum menginput sisi dari segitiga sama sisi"
		if !cetakKalimat {
			// Memicu panic sesuai permintaan soal.
			panic(errMsg)
		}
		// Mengembalikan error biasa.
		return "", errors.New(errMsg)
	}

	keliling := sisi * 3
	if cetakKalimat {
		// Mengembalikan hasil dalam format kalimat lengkap.
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}

	// Mengembalikan hasil hanya dalam bentuk string angka.
	return strconv.Itoa(keliling), nil
}

// soal 3
// Function ini menerima pointer dan menambahkan nilai ke variabel yang ditunjuk.
func tambahAngka(nilai int, total *int) {
	*total += nilai
}

// Function ini menerima pointer dan mencetak nilai dari variabel yang ditunjuk.
func cetakAngka(total *int) {
	fmt.Println("\n--- Output Soal 3 (dari defer) ---")
	fmt.Printf("Total akhir angka adalah: %d\n", *total)
}

// soal 4
// Function ini menambahkan data ke sebuah slice via pointer, mengurutkannya,
// lalu menampilkannya satu per satu dengan jeda waktu.
func populateAndPrintPhones(phoneList *[]string) {
	newPhones := []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"}
	*phoneList = append(*phoneList, newPhones...)

	// Mengurutkan data dalam slice secara alphabetical.
	sort.Strings(*phoneList)

	fmt.Println("\n--- Output Soal 4 ---")
	for i, brand := range *phoneList {
		fmt.Printf("%d. %s\n", i+1, brand)
		time.Sleep(1 * time.Second)
	}
}

// soal 5
// Function ini menampilkan data dari slice menggunakan goroutine
// dan WaitGroup untuk memastikan semua proses selesai.
func printPhonesConcurrently() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup

	// Mengurutkan data sebelum diproses.
	sort.Strings(phones)

	fmt.Println("\n--- Output Soal 5 ---")

	for i, brand := range phones {
		// Menambah counter WaitGroup untuk setiap goroutine yang akan dijalankan.
		wg.Add(1)

		// Menjalankan goroutine untuk setiap item.
		// `i` dan `brand` dioper sebagai argumen agar tidak terjadi race condition.
		go func(index int, phoneBrand string) {
			// `defer wg.Done()` akan mengurangi counter saat goroutine selesai.
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("%d. %s\n", index+1, phoneBrand)
		}(i, brand)
	}

	// Menunggu hingga semua goroutine selesai (counter WaitGroup menjadi 0).
	wg.Wait()
}

func main() {
	// soal 1
	defer displayInfo("Golang Backend Development", 2021)
	fmt.Println("Main function dimulai...")

	// soal 2
	fmt.Println("\n--- Output Soal 2 ---")
	// Test case 1
	hasil1, err1 := kelilingSegitigaSamaSisi(4, true)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(hasil1)
	}
	// Test case 2
	hasil2, err2 := kelilingSegitigaSamaSisi(8, false)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(hasil2)
	}
	// Test case 3 (error)
	_, err3 := kelilingSegitigaSamaSisi(0, true)
	if err3 != nil {
		fmt.Println(err3)
	}
	// Test case 4 (panic)
	kelilingSegitigaSamaSisi(0, false)

	// soal 3
	angka := 1
	defer cetakAngka(&angka) // Panggilan ini akan dieksekusi sebelum main selesai.
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// soal 4
	var phones = []string{}
	populateAndPrintPhones(&phones)

	// soal 5
	printPhonesConcurrently()

	fmt.Println("\nMain function selesai.")
}