package main

import "fmt"

func getIdentifySiswa(fullname string, age int, address string) (string, int, string) {
	return fullname, age, address
}

func main() {
	fullname := "Riyanwar Setiadi"
	age := 23
	address := "Getrakmoyan"
	fmt.Println(getIdentifySiswa(fullname, age, address))
}
