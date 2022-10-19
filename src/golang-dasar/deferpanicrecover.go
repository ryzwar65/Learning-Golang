// PENJELASAN DEFER
// Defer function merupakan function yang bisa di jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// Defer pun akan tetap dieksekusi walaupun terjadi error di function yang di eksekusi

// Example Defer

package main

import "fmt"

func logError() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
}

func log() {
	fmt.Println("SYSTEM BERJALAN....")
}

func luasKubus(s int) int {

	return s * s * s
}

func runApp(error bool) bool {
	defer logError()
	if error {
		panic("APLIKASI ERROR")
	}
	return error
}

func main() {
	// Running Defer
	fmt.Println(luasKubus(10))
	runApp(true)
	defer log()
}
