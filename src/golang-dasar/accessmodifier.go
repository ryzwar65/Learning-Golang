// Biasanya di Bahasa pemograman lain ada kata kunci untuk access dimanapun
// Di golang menentukan access modifer cukup dengan nama function atau variable
// Jika nama diawali dengan hurup besar maka artinya bisa diakses dari package yang lain, jika kecil tidak bisa di akses di package lain

package main

import (
	"fmt"
	"golang-dasar/helpers"
)

// Contoh di bawah type struct di package yang lain

func main() {
	alamat := make(map[string]string)
	alamat["City"] = "Denpasar"
	alamat["Province"] = "Bali"
	afi := new(helpers.AlamatSaya)
	fmt.Println(afi.Alamatku(alamat))
}
