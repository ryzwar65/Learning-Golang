package main

import "fmt"

func sumAll(number ...int) int{
	total := 0
	
	for _, value := range number {
		total += value
	}

	return total
}

func main() {
	myNumber := [] int {10,2,10,20,12,2,0}

	total := sumAll(myNumber...)

	fmt.Println(total)
}