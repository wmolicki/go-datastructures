package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[int]()

	s.Add(5)
	s.Add(3)
	s.Add(3)
	s.Add(10)

	assert.Equal(t, s.In(5), true)
	assert.Equal(t, s.In(3), true)
	assert.Equal(t, s.In(10), true)
	assert.Equal(t, s.In(-3), false)
	assert.Equal(t, s.Len(), 3)
}

func TestDifference(t *testing.T) {
	s := NewSet(1, 2, 3)
	o := NewSet(4, 5, 3)

	assert.Equal(t, s.Difference(o), NewSet(1, 2))
	assert.Equal(t, o.Difference(s), NewSet(4, 5))

	o.Add(1)
	o.Remove(3)

	assert.Equal(t, NewSet(2, 3), s.Difference(o))
}

func TestIntersection(t *testing.T) {
	s := NewSet(1, 2, 3)
	o := NewSet(4, 5, 3)

	assert.Equal(t, NewSet(3), s.Intersection(o))
	assert.Equal(t, NewSet(3), o.Intersection(s))

	s.Add(4)
	o.Add(1)

	assert.Equal(t, NewSet(1, 3, 4), s.Intersection(o))
}

func TestIsSubsetOf(t *testing.T) {
	tcs := []struct {
		tested *Set[int]
		other  *Set[int]
		want   bool
	}{
		{tested: NewSet[int](), other: NewSet[int](), want: true},
		{tested: NewSet(1, 2, 3), other: NewSet[int](), want: false},
		{tested: NewSet[int](), other: NewSet(1, 2), want: true},
		{tested: NewSet(5), other: NewSet(5), want: true},
		{tested: NewSet(5, 6), other: NewSet(5), want: false},
		{tested: NewSet(5, 6), other: NewSet(5, 6, 7, 8, 7), want: true},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.want, tc.tested.IsSubsetOf(tc.other))
	}
}

func TestIsSupersetOf(t *testing.T) {
	tcs := []struct {
		tested *Set[int]
		other  *Set[int]
		want   bool
	}{
		{other: NewSet[int](), tested: NewSet[int](), want: true},
		{other: NewSet[int](), tested: NewSet(1, 2, 3), want: true},
		{other: NewSet[int](), tested: NewSet(1, 2), want: true},
		{other: NewSet(5), tested: NewSet(5), want: true},
		{other: NewSet(5, 6), tested: NewSet(5), want: false},
		{other: NewSet(5, 6), tested: NewSet(5, 6, 7, 8, 7), want: true},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.want, tc.tested.IsSupersetOf(tc.other), fmt.Sprintf("%+v", tc))
	}
}

func TestCopy(t *testing.T) {
	s := NewSet(1, 2, 3)
	c := s.Copy()

	assert.NotSame(t, s, c, "copy returned reference to the same object")
	assert.Equal(t, s.Difference(c), NewSet[int]())
	assert.Equal(t, c.In(1), true)
	assert.Equal(t, c.In(2), true)
	assert.Equal(t, c.In(3), true)
	assert.Equal(t, 3, c.Len())
	s.Add(5)
	assert.False(t, c.In(5))
}

func TestUnion(t *testing.T) {
	tcs := []struct {
		tested *Set[int]
		other  *Set[int]
		want   *Set[int]
	}{
		{other: NewSet[int](), tested: NewSet[int](), want: NewSet[int]()},
		{other: NewSet[int](), tested: NewSet(1, 2, 3), want: NewSet(1, 2, 3)},
		{other: NewSet(5), tested: NewSet(5), want: NewSet(5)},
		{other: NewSet(5, 6), tested: NewSet(5), want: NewSet(5, 6)},
		{other: NewSet(7, 9, 10), tested: NewSet(5, 6, 7, 8, 7), want: NewSet(5, 6, 7, 8, 9, 10)},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.want, tc.tested.Union(tc.other), fmt.Sprintf("%+v", tc))
	}
}

func TestItems(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5, 6, 7)
	items := s.Items()
	assert.Equal(t, 7, len(items))
	assert.Equal(t, 7, cap(items))
}
