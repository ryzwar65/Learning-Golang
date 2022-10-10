package main

import "fmt"

func main() {
	number1 := 40
	number2 := 50

	total := number1 + number2

	switch total {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	// Short Statment

	switch length := len("AN"); length == 2 {
	case true:
		fmt.Println("BENAR PANJANG : ", length)
	default:
		fmt.Println("SALAH")
	}

	// Switch Tanpa Kondisi

	switch {
	case total > 85:
		fmt.Println("A")
	case total >= 75 && total < 80:
		fmt.Println("B")
	case total >= 75 && total < 80:
		fmt.Println("B")
	default:
		fmt.Println("E")

	}
}
