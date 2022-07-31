package main

import "fmt"

func main() {
	angka16 := 127
	var (
		firstName = "Riyanwar"
		lastName  = "Setiadi"
	)

	var r byte = firstName[0]
	var rString string = string(r)
	var s byte = lastName[0]
	var sString string = string(s)

	// Conversion
	fmt.Println(int8(angka16))
	fmt.Println(rString)
	fmt.Println(sString)

}
