package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		fmt.Println("ITERASI KE ",i)
		if i == 5 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}	
		fmt.Println("INI PERULANGAN GANJIL NUMBER ",i)
	}
}