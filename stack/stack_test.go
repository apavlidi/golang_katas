package fizzBuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := Stack{}
	assert.Equal(t, stack.IsEmpty(), true)
}

func TestIsNotEmpty(t *testing.T) {
	stack := Stack{elements: make([]int, 1)}
	assert.Equal(t, stack.IsEmpty(), false)
}

func TestSize(t *testing.T) {
	stack := Stack{}
	assert.Equal(t, stack.Size(), 0)
}

func TestIsPushed(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	assert.Equal(t, stack.Size(), 1)
}

func TestIsPopped(t *testing.T) {
	stack := Stack{}
	stack.push(5)
	item := stack.pop()
	assert.Equal(t, 5, item)
	assert.Equal(t, stack.Size(), 0)
}

func TestPushAndPopMultipleItems(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(2)
	stack.push(2)
	stack.push(3)
	item := stack.pop()
	item2 := stack.pop()
	assert.Equal(t, 3, item)
	assert.Equal(t, 2, item2)
	assert.Equal(t, stack.Size(), 3)
}

func TestPeek(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.push(5)
	stack.push(23)
	item := stack.peek()
	assert.Equal(t, stack.Size(), 3)
	assert.Equal(t, item, 23)
}
