package main

import "fmt"

func main() {
	type NoKTP int32

	var noktp NoKTP = 1234567890

	fmt.Println(noktp)
}
