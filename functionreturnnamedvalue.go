package main

import "fmt"

func getIdentifySiswa() (fullname string, age int8, address string) {
	fullname = "Riyanwar Setiadi"
	age = 23
	address = "Getrakmoyan"

	return fullname, age, address
}

func main() {
	fullname, age, address := getIdentifySiswa()

	fmt.Println(fullname, age, address)
}
