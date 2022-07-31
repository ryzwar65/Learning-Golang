package main

import "fmt"

func main() {
	// ... digunakan saat indexnya tidak bisa di tentukan (Bakal Banyak Array disitu)
	var bulan = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[1:5]
	var panjangSlice1 = len(slice1)
	var kapasitasSlice1 = cap(slice1)
	fmt.Println(slice1)
	fmt.Println("Panjang Slice 1 : ", panjangSlice1)
	fmt.Println("Kapasitas Slice 1 : ", kapasitasSlice1)

	// membuat Slice Baru

	var slice2 = make([]string, panjangSlice1, kapasitasSlice1)
	slice2[0] = "INI SLICE BARU"
	slice2[1] = "Riyanwar Setiadi"
	fmt.Println("Slice 2 ", slice2)

	// Append Data Baru

	sliceAppend := append(bulan[:3], "Riyanwar")

	// Perubahan Array Slice
	// slice1[0] = "January"
	// fmt.Println(slice1)

	fmt.Println(sliceAppend)
	fmt.Println("BULAN", bulan)

	// Perbedaan Slice dan Array

	iniArray := [...]int8{1, 2, 3, 4, 5}
	iniSlice := []int8{1, 2, 3, 4, 5}

	fmt.Println("Ini Array (Harus Deklarasi Index berapa atau ... artinya tak terbatas)", iniArray)
	fmt.Println("Ini Array (Tanpa Deklarasi Index)", iniSlice)
}
