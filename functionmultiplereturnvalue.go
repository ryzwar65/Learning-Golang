package main

import "fmt"

func getMultipleValue()(string, string, byte) {
	return "Riyanwar", "Setiadi", 24
}

func main() {
	firstName, lastName, _ := getMultipleValue()

	fmt.Println(firstName,lastName)
}