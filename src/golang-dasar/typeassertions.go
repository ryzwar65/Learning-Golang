// Merubah tipe data menjadi tipe data yang kita ingingkan

package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	result := random()
	switch result.(type) {
	case string:
		fmt.Println(result.(string))
	case int:
		fmt.Println(result.(int))
	case bool:
		fmt.Println(result.(bool))
	default:
		fmt.Println("Tipe Tidak Di ketahui")
	}
}
