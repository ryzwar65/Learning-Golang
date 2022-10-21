package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtest(t *testing.T) {
	t.Run("Riyanwar", func(t *testing.T) {
		result := HelloWorld("Riyanwar")
		assert.Equal(t, "Hello Riyanwar", result, "Result must be Riyanwar")
	})
	t.Run("Setiadi", func(t *testing.T) {
		result := HelloWorld("Setiadi")
		assert.Equal(t, "Hello Setiadi", result, "Result must be Setiadi")
	})
}
