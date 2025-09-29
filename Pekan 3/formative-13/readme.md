
# Bukti Koneksi: API Gin (Go) dengan Database PostgreSQL

Dokumen ini menunjukkan bukti bahwa aplikasi API yang dibangun menggunakan framework **Gin** di Go telah berhasil terhubung dan dapat melakukan operasi **tulis (write)** ke database **PostgreSQL**.

---

## 📌 Skenario Pengujian

1. Client mengirimkan request `POST` ke endpoint `/bioskop` pada API server.  
2. API Server (Gin) menerima data, memprosesnya, dan menyimpannya ke dalam tabel `bioskops` di database PostgreSQL.  
3. API Server (Gin) mengembalikan respon **201 Created** beserta data yang baru saja dibuat, yang kini sudah memiliki **ID dari database**.  
4. Database (PostgreSQL) diverifikasi untuk memastikan data tersebut benar-benar tersimpan secara permanen.  
5. Bukti tambahan berupa rekaman proses penambahan data dapat dilihat di sini:  
   - [Jam.dev Recording – Uji coba awal](https://jam.dev/c/43d931d0-fde0-400b-a63f-f5dbe5f685f2?startFrom=0.00)  
   - [Jam.dev Recording – Uji coba tugas 14](https://jam.dev/c/61a6195e-f054-4a0e-ac53-8c984f0a4114)  

---

## 🔹 Langkah 1: Pengiriman Request dari Client

Sebuah request `POST` dikirimkan ke `http://localhost:8080/bioskop` menggunakan **Postman**.  
Request body berisi data bioskop baru dalam format JSON.  

**Detail Request (Tugas 14):**

- Method: `POST`  
- URL: `http://localhost:8080/bioskop`  
- Body (raw JSON):  

```json
{
  "nama": "XXI Ambarukmo Plaza",
  "lokasi": "Yogyakarta",
  "rating": 4.8
}
````

---

## 🔹 Langkah 2: Respon Sukses dari API Server

Setelah request diproses, API server memberikan respon dengan status **201 Created**.
Respon ini menandakan bahwa data telah berhasil dibuat.

**Detail Response (Tugas 14):**

* Status Code: `201 Created`
* Body (JSON):

```json
{
  "data": {
    "id": 3,
    "nama": "XXI Ambarukmo Plaza",
    "lokasi": "Yogyakarta",
    "rating": 4.8
  }
}
```

👉 Pemberian `id: 3` oleh sistem adalah bukti bahwa aplikasi telah berinteraksi dengan database untuk menyimpan data dan mendapatkan **primary key** yang baru.

---

## 🔹 Langkah 3: Verifikasi Data di Database PostgreSQL

Langkah berikutnya adalah memeriksa langsung ke dalam database PostgreSQL untuk memastikan data benar-benar tersimpan.

**Langkah verifikasi dengan pgAdmin:**

1. Koneksi ke database `bioskop_db`.
2. Membuka tabel `public.bioskops`.
3. Melihat semua baris data (rows) yang ada di dalam tabel.

**Hasil Verifikasi (Tugas 14):**
Data untuk **"XXI Ambarukmo Plaza" dengan id 3** berhasil ditemukan dalam tabel.

Contoh hasil query:

| id | nama                | lokasi     | rating |
| -- | ------------------- | ---------- | ------ |
| 2  | CGV Jwalk Mall      | Yogyakarta | 4.6    |
| 3  | XXI Ambarukmo Plaza | Yogyakarta | 4.8    |

---

## ✅ Kesimpulan

Berdasarkan alur pengujian di atas:

* Request `POST` berhasil diterima dan diproses oleh server Gin.
* Server mengembalikan respon **201 Created** dengan data yang sudah dilengkapi ID dari database.
* Data tersebut terbukti ada dan tersimpan secara permanen di dalam tabel `bioskops` pada database PostgreSQL.
* Bukti tambahan (rekaman langsung proses penambahan data):

  * [Jam.dev Recording – Uji coba awal](https://jam.dev/c/43d931d0-fde0-400b-a63f-f5dbe5f685f2?startFrom=0.00)
  * [Jam.dev Recording – Uji coba tugas 14](https://jam.dev/c/61a6195e-f054-4a0e-ac53-8c984f0a4114)

👉 Dengan demikian, koneksi dan integrasi antara API server Gin dan database PostgreSQL **telah berhasil dikonfigurasi dan berfungsi dengan benar**.

