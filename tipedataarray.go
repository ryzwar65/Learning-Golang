package main

import "fmt"

func main() {
	var listAngka = [3]int{
		100,
		200,
		300,
	}
	var listHuruf [3]string
	fmt.Println(listAngka)
	listHuruf[0] = "A"
	listHuruf[1] = "B"
	listHuruf[2] = "C"
	fmt.Println(listHuruf)
	fmt.Println(len(listAngka))
	fmt.Println(len(listHuruf))
}
