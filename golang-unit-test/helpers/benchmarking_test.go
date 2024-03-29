package helpers

import (
	"testing"
)

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("ANWAR")
	}
	b.Run("riyanwar-benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Riyanwar")
		}
	})
	b.Run("setiadi-benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Setiadi")
		}
	})
}
