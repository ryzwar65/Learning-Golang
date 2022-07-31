package main

import "fmt"

func main() {
	iterasi := 10

	for i := 0; i < iterasi; i++ {
		fmt.Println(i)
	}

	for iterasi < 20 {
		fmt.Println(iterasi)
		iterasi++
	}

	// For Range untuk data collection

	dataSlice := []string{"Riyanwar","Setiadi"}

	// Contoh 1
	for j := 0; j < len(dataSlice); j++ {
		fmt.Println(dataSlice[j])
	}
	// Contoh 2
	for _,name := range dataSlice{
		fmt.Println(name)
	}
}
