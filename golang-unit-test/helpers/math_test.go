package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Contoh Assert
func TestSum(t *testing.T) {
	result := Sum(int8(10), int8(10))
	assert.Equal(t, int8(20), result, "Result Must be 20")
	fmt.Println("Done")
}

func TestSumRequire(t *testing.T) {
	result := Sum(int8(10), int8(10))
	require.Equal(t, int8(20), result, "Result Must be 20")
	fmt.Println("Done")
}
