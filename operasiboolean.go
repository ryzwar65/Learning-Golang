package main

import "fmt"

func main() {
	var a, b bool
	a = true
	b = false
	fmt.Println(a && b)
	fmt.Println(booleanFuntion(76))
}

func booleanFuntion(nilai int8) bool {
	var lulusUjian bool = nilai > 75
	return lulusUjian
}
