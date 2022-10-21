package helpers

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Riyanwar")
	if result != "Hello Riyanwar" {
		t.Error("Result Must be Hello Riyanwar")
	}
	fmt.Println("Done")
}

func TestHelloWorldRiyanwarSetiadi(t *testing.T) {
	result := HelloWorld("Riyanwar Setiadi")
	if result != "Hello Riyanwar Setiadi" {
		t.Fatal("Result Must be Hello Riyanwar Setiadi")
	}
}
