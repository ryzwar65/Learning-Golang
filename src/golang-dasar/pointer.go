// Secara default golang passing data by value

package main

import "fmt"

type Alamat struct {
	City, Provinsi, Country string
}

func main() {
	mamat := Alamat{"Cirebon", "Jawa Barat", "Indonesia"}
	budi := mamat
	budi.City = "Jakarta"
	// Diatas merupakan Pass by Value dimana Alamat yang dimiliki mamat akan di clone ke data budi, sehingga budi merubah city itu tidak mengganti city mamat
	// Fungsi Pointer merupakan ketika kita ingin merubah data secara total makanya menggunakan pass reference dengan cara memanfaatkan pointer

	// Contoh Pass reference Pointer

	var amad *Alamat = &mamat
	amad.City = "Bandung"

	// Contoh diatas merupakan pass by reference sehingga alamat mamat pun ikut terganti

	// Pointer jg bisa digunakan sebagai operator
	// Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
	// Semua Variable yang mengacu ke data yang sama tidak akan berubah
	// Jika kita ingin mengubah variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
	*amad = Alamat{"Indramayu", "Jawa Timur", "Indonesia"}

	// Function New merupakan pembuatan pointer hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

	billy := new(Alamat)
	billy.City = "Brebes"

	fmt.Println(mamat)
	fmt.Println(budi)
	fmt.Println(amad)
	fmt.Println("Update alamat budi sekaligus mamat ", amad)
	fmt.Println(billy)
}
