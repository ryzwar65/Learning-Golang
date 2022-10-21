package main

import "fmt"

type Siswa struct {
	Fullname string
}

func (siswa *Siswa) Datadiri() {
	siswa.Fullname = "Nak " + siswa.Fullname
}

func main() {
	budi := Siswa{"Budi Santoso"}
	budi.Datadiri()
	asep := Siswa{"Asep cecep"}
	asep.Datadiri()
	fmt.Println(budi.Fullname)
	fmt.Println(asep.Fullname)
}
