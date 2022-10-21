package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTabletest(t *testing.T) {
	data := []struct {
		name_test string
		request   string
		expected  string
	}{
		{
			name_test: "HelloWorld(Riyanwar)",
			request:   "Riyanwar",
			expected:  "Hello Riyanwar",
		},
		{
			name_test: "HelloWorld(Setiadi)",
			request:   "Setiadi",
			expected:  "Hello Setiadi",
		},
	}
	for _, d := range data {
		t.Run(d.name_test, func(t *testing.T) {
			result := HelloWorld(d.request)
			assert.Equal(t, d.expected, result)
		})
	}
}
