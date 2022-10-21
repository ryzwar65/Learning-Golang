package main

import "fmt"

func main() {
	number1 := 40
	number2 := 50

	total := number1 + number2

	if total >= 70 && total <= 80 {
		fmt.Println("B")
	} else if total >= 60 && total < 70 {
		fmt.Println("C")
	} else if total < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("A")
	}

	// Short Statement

	if nilai := 'A'; nilai == 'A' {
		fmt.Println("Kamu Pintar banget!")
	}
}
