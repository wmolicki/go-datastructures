package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()

	assert.True(t, true, s.IsEmpty())

	s.Push(1)
	s.Push(3)
	s.Push(5)
	s.Push(0)

	var numbers []int
	for _, i := range *s {
		numbers = append(numbers, i)
	}
	assert.Equal(t, []int{1, 3, 5, 0}, numbers)

	item, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 0, item)

	item, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 5, item)

	item, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, item)

	item, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, item)

	item, ok = s.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, item)
}
