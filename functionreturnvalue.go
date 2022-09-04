package main

import (
	"fmt"
)

func totalHarga(hargaProduk uint32, qty uint16) uint32 {
	return hargaProduk * uint32(qty)
}

func diskonHarga(hargaProduk float64, qty float64) float64 {
	return hargaProduk * qty * (float64(30) / float64(100))
}

func main() {	
	fmt.Println(totalHarga(100000, 10))
	fmt.Printf("%.2f",float64(diskonHarga(100000,10)))	
}