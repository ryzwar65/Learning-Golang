package main

import "fmt"

type HasName interface {
	SetName() string
}

type typeAttribute struct {
	FirstName, LastName string
}

func (ta typeAttribute) SetName() string {
	return ta.FirstName + " " + ta.LastName
}

func GetName(hasName HasName) {
	fmt.Println(hasName.SetName())
}

func main() {
	person := typeAttribute{"Riyanwar", "Setiadi"}
	GetName(person)
}
