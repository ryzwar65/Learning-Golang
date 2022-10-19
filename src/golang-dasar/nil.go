package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	newMap := NewMap("")
	if newMap == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(newMap)
	}
}
