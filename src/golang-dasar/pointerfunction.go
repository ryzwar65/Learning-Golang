package main

import "fmt"

type AlamatOrang struct {
	City, Provinsi, Country string
}

func RubahAlamatOrang(alamatOrang *AlamatOrang) {
	alamatOrang.City = "Kupang"
}

func main() {
	var sueb = AlamatOrang{
		City:     "Lombok",
		Provinsi: "NTB",
		Country:  "Indonesia",
	}

	RubahAlamatOrang(&sueb)

	fmt.Println(sueb)
}
