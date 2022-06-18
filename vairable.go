package main

import "fmt"

func main() {
	var angka int8
	angka = 10

	fmt.Println("Variabel Integer : ",angka)
	angka += 15
	fmt.Println("Variabel Integer : ",angka)

	var name = "Riyanwar"
	fmt.Println(name)

	firstName := "Riyanwar"
	lastName := "Setiadi"

	fmt.Println(firstName+" "+lastName)

	firstName = "Nurlaela"
	lastName = "Khasannah"

	fmt.Println(firstName+" "+lastName)

	var (
		country = "Indonesia"
		province = "Jawa Barat"
	)

	fmt.Println(country)
	fmt.Println(province)
	

}