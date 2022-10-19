package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		counter := 1
		counter++
		fmt.Println("Output dari closure ", counter)
	}
	// Hati-hati saat menggunakan closure ketika variabel yang dibuat sudah final tidak boleh di ganti lagi
	increment()
	fmt.Println(counter)
}
