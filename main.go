package main

import "fmt"

//Suhu celcius bedasarkan kondisi cuaca di Indonesia
//Variabel untuk menyimpan input suhu
func main() {
	var suhu float64

	//Meminta user memasukkan suhu
	fmt.Print("Masukkan suhu dalam Celcius: ")
	validasi, err := fmt.Scanln(&suhu)
	if err != nil || validasi != 1 {
		fmt.Println("Input tidak valid.")
		return
	}

	//Menentukan kategori suhu berdasarkan standar Indonesia
	if suhu < 20 {

		//Suhu di bawah 20°C (daerah pegunungan seperti Dieng, Puncak)
		fmt.Println("Kategori: Dingin ❄️")
		fmt.Println("Cocok untuk daerah pegunungan seperti Dieng atau Puncak.")
	} else if suhu >= 20 && suhu <= 25 {

		//Suhu antara 20-25°C
		fmt.Println("Kategori: Sejuk/Normal 🌿")
		fmt.Println("Suhu nyaman untuk beraktivitas.")
	} else if suhu >= 26 && suhu <= 32 {

		//	Suhu antara 26-32°C
		fmt.Println("Kategori: Hangat/Cukup Panas ☀️")
		fmt.Println("Suhu umum di sebagian besar wilayah Indonesia.")
	} else if suhu < 32 {

		//Suhu di atas 32°C
		fmt.Println("Kategori: Panas 🔥")
		fmt.Println("Disarankan untuk minum air yang cukup dan hindari terik matahari langsung.")
	}
}
