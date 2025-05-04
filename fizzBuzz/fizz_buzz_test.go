package fizzBuzz

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			assert.Equal(t, test.expected, FizzBuzz(test.input))
		})
	}
}
