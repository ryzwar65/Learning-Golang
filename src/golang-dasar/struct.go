package main

import "fmt"

type Mahasiswa struct {
	Fullname, Address string
}

func main() {
	var mahasiswa Mahasiswa

	mahasiswa.Fullname = "Budi Setiawan"
	mahasiswa.Address = "Cilandak Barat"

	fmt.Println(mahasiswa)

	// Struct Literal

	fajar := Mahasiswa{
		Fullname: "Fajar Fahrobi",
		Address:  "Ciledug",
	}

	alex := Mahasiswa{"Alex Sulistian", "Mangga Dua"}

	fmt.Println(fajar)

	fmt.Println(alex)
}
