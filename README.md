# Kategori Suhu Indonesia (Indonesian Weather Categorizer)

Program sederhana berbasis **Golang** untuk mengkategorikan kondisi suhu (Celcius) berdasarkan standar kenyamanan cuaca di Indonesia.

## 📝 Deskripsi

Aplikasi _Command Line Interface_ (CLI) ini meminta pengguna untuk memasukkan angka suhu dalam satuan Celcius. Program kemudian akan mengevaluasi angka tersebut dan memberikan deskripsi kondisi cuaca yang sesuai dengan iklim tropis di Indonesia (mulai dari dingin pegunungan hingga panas terik).

## 🚀 Fitur

* **Validasi Input:** Memastikan input yang dimasukkan adalah angka.
* **Kategorisasi Lokal:** Standar suhu disesuaikan dengan persepsi orang Indonesia:
    * ❄️ **Dingin:** < 20°C (Area Pegunungan)
    * 🌿 **Sejuk:** 20°C - 25°C
    * ☀️ **Hangat:** 26°C - 32°C
    * 🔥 **Panas:** > 32°C

## 🛠️ Prasyarat

* [Go (Golang)](https://go.dev/dl/) versi 1.18 atau terbaru sudah terinstall di komputer Anda.

## 💻 Cara Menjalankan

1.  **Clone repository ini** (atau download filenya):
    ```bash
    git clone https://github.com/Sapiiik/Pengembangan-Sistem-Klasifikasi-Suhu-Wilayah-Indonesias
    cd kategori-suhu-indonesia
    ```

2.  **Jalankan program:**
    ```bash
    go run main.go
    ```

3.  **Build (Opsional):**
    Jika ingin membuat file executable (aplikasi):
    ```bash
    go build -o cek-suhu
    ./cek-suhu
    ```

## 📸 Contoh Penggunaan

```text
Masukkan suhu dalam Celcius: 18
Kategori: Dingin ❄️
Cocok untuk daerah pegunungan seperti Dieng atau Puncak.
