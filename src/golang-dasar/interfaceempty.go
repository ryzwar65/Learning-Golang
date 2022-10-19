package main

import "fmt"

func booleanFunc() interface{} {
	return true
}

func stringFunc() interface{} {
	return "Ini String"
}

func integerFunc() interface{} {
	return 100
}

func main() {
	booleanVar := booleanFunc()
	stringVar := stringFunc()
	integerVar := integerFunc()
	fmt.Println(booleanVar)
	fmt.Println(stringVar)
	fmt.Println(integerVar)
}
