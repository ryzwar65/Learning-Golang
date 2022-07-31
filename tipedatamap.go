package main

import "fmt"

func main() {
	type Mahasiswa map[string]string
	mahasiswa := Mahasiswa{
		"name":    "Riyanwar Setiadi",
		"address": "Desa Getrakmoyan",
		"jurusan": "Teknik Informatika",
	}

	mahasiswa["phone"] = "083120615229"

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["name"])

	fmt.Println(mahasiswa["address"])
	delete(mahasiswa, "phone")

	var phone map[string]uint64 = make(map[string]uint64)
	phone["phone"] = 83120615229

	mahasiswa["name"] = "Nurlaela Khasannah"

	fmt.Println(mahasiswa)
	fmt.Println(len(mahasiswa))
	fmt.Println(phone)

}
