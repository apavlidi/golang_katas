package stringCalculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringCalculator(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Add empty string", "", 0},
		{"Add 1,2", "1,2", 3},
		{"Add 4", "4", 4},
		{"Add 1\n2,3", "1\n2,3", 6},
	}

	for _, tc := range tests {
		actual, err := StringCalculator(tc.input)
		require.NoError(t, err)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestStringCalculatorInvalidChars(t *testing.T) {
	actual, err := StringCalculator("1,2,s,5")
	require.Error(t, err)
	assert.Equal(t, -1, actual)
}
