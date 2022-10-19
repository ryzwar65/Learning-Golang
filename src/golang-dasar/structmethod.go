package main

import "fmt"

type Penjual struct {
	Address, Fullname string
	phone             uint64
}

func (penjual *Penjual) identitasPenjual() {
	fmt.Println("Nama Penjual : ", penjual.Fullname)
}

func main() {
	penjual := Penjual{"Getrakmoyan", "Riyanwar Setiadi", 6283120615229}
	penjual.identitasPenjual()
}
