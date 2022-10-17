package main

import (
	"errors"
	"fmt"
)

func Pembagian(angka int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return angka / pembagi, nil
	}
}

func main() {
	hasil, error := Pembagian(10, 0)
	if error == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error : ", error)
	}
}
