package helpers

import (
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot Run on Linux")
	}
}
