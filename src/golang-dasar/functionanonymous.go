package main

import "fmt"

type FuncBlackListUsername func(string) bool

func main() {
	usernameblocked := func(name string) bool {
		return name == "anjingJago"
	}

	name := usernameblocked("anjingJago")
	if name {
		fmt.Println("Your username is Blocked")
	} else {
		fmt.Println("Not Blocked")
	}
}
