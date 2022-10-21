package main

import "fmt"

func getWelcome(name string) string {
	return "Welcome!! " + name
}

func main() {
	welcome := getWelcome
	fmt.Println(welcome("Anwar"))
}
