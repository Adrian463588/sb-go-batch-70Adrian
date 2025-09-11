package main

import "fmt"

// soal 5: fungsi-fungsi luas, keliling, volume
func luasPersegiPanjang(panjang, lebar int) int {
    return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
    return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
    return panjang * lebar * tinggi
}

func main() {
    // soal 1
    for i := 1; i <= 7; i++ {
        for j := 0; j < i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }

    // soal 2
    var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
    // kita ingin [saya sangat senang belajar golang]
    // ambil slice dari elemen index 2 sampai akhir
    hasil := kalimat[2:] // mulai dari "saya"
    fmt.Println(hasil)

    // soal 3
    var sayuran = []string{}
    sayuran = append(sayuran, "Bayam")
    sayuran = append(sayuran, "Buncis")
    sayuran = append(sayuran, "Kangkung")
    sayuran = append(sayuran, "Kubis")
    sayuran = append(sayuran, "Seledri")
    sayuran = append(sayuran, "Tauge")
    sayuran = append(sayuran, "Timun")
    for idx, v := range sayuran {
        fmt.Printf("%d. %s\n", idx+1, v)
    }

    // soal 4
    var satuan = map[string]int{
        "panjang": 7,
        "lebar":   4,
        "tinggi":  6,
    }
    for key, val := range satuan {
        fmt.Printf("%s: %d\n", key, val)
    }

    // soal 5
    panjang := 12
    lebar := 4
    tinggi := 8

    luas := luasPersegiPanjang(panjang, lebar)
    keliling := kelilingPersegiPanjang(panjang, lebar)
    volume := volumeBalok(panjang, lebar, tinggi)

    fmt.Println(luas)
    fmt.Println(keliling)
    fmt.Println(volume)
}
