package helpers

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE")
	m.Run()
	fmt.Println("AFTER")
}
