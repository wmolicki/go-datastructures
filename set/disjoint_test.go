package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisjointSetUnion(t *testing.T) {
	ds := NewDisjointSet(10)

	ds.Union(1, 2)
	ds.Union(2, 5)
	ds.Union(5, 6)
	ds.Union(6, 7)
	ds.Union(3, 8)
	ds.Union(8, 9)

	assert.Equal(t, true, ds.Connected(1, 5))
	assert.Equal(t, true, ds.Connected(5, 7))
	assert.Equal(t, false, ds.Connected(4, 9))

	ds.Union(9, 4)

	assert.Equal(t, true, ds.Connected(4, 9))
}
