package helpers

import "testing"

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		FirstName, LastName string
	}{
		{
			FirstName: "Riyanwar",
			LastName:  "Setiadi",
		},
		{
			FirstName: "Jamal",
			LastName:  "Anugrah",
		},
	}
	for _, bench := range benchmark {
		b.Run(bench.FirstName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.LastName)
			}
		})
	}
}
