package main

import "fmt"

func getMultipleValue()(firstName string, lastName string,age byte) {
	firstName = "Anwar"
	lastName = "Setiadi"
	age = 23

	return
}

func main() {
	first, last, _ := getMultipleValue()

	

	fmt.Println(first,last)
}