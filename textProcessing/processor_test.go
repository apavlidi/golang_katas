package textProcessing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessor_CountWords(t *testing.T) {

	processor := TextProcessor{}
	analytics := processor.analyse("Hello, this is an example for you to practice. You should grab this text and make it as your test case.")

	//assert.Equal(t, []string{"you", "this", "your", "and", "hello,", "it", "grab", "should", "for", "an"}, actual)
	assert.Equal(t, analytics.wordCount, 21)
}

func TestProcessor_CountTime(t *testing.T) {

}
