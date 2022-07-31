package main

import "fmt"

func totalHarga(hargaProduk uint32, qty uint16) uint32 {
	return hargaProduk * uint32(qty)
}

func main() {
	fmt.Println(totalHarga(100000, 10))
}