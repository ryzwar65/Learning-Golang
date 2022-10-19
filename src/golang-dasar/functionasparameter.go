package main

import (
	"fmt"
	"strings"
)

type FuncString func(string) string

func getName(name string, setName FuncString) {

	setname := setName(name)

	fmt.Println("Hello!! ", setname)

}

func setName(name string) string {
	return name
}

func upperFunc(name string) string {
	return strings.ToTitle(name)
}

func main() {
	getName("Riyanwar", setName)
	getName("Riyanwar", upperFunc)
}
