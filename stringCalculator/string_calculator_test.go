package stringCalculator

import (
	"errors"
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
		{"Add single character", "4", 4},
		{"Add multiple characters 1,2", "1,2", 3},
		{"Multiple Separators", "1\n2,3", 6},
		{"Custom separators", "//;\n1;2;4", 7},
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

func TestStringCalculatorNegativeNumber(t *testing.T) {
	actual, err := StringCalculator("1,2,3,-5")
	require.Error(t, err)
	var customErr *NotAllowedNegativeNumbers
	ok := errors.As(err, &customErr)
	assert.True(t, ok, "error should be of type NotAllowedNegativeNumbers")
	assert.Equal(t, -1, actual)
}
