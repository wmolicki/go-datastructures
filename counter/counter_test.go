package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepareCounter() *Counter {
	c := NewCounter()

	c.Add("a")
	c.Add("a")
	c.Add("a")

	c.Add("b")
	c.Add("b")

	c.Add("c")

	return c
}

func TestMostCommon(t *testing.T) {
	c := prepareCounter()

	tcs := []struct {
		n    int
		want []CounterElem
	}{
		{n: 1, want: []CounterElem{{Item: "a", Count: 3}}},
		{n: 2, want: []CounterElem{{Item: "a", Count: 3}, {Item: "b", Count: 2}}},
		{n: 3, want: []CounterElem{{Item: "a", Count: 3}, {Item: "b", Count: 2}, {Item: "c", Count: 1}}},
	}

	for _, tc := range tcs {
		got := c.MostCommon(tc.n)
		assert.Equal(t, got, tc.want)
	}
}

func TestTotal(t *testing.T) {
	c := prepareCounter()

	assert.Equal(t, c.Total(), 6)
}

func TestSubtract(t *testing.T) {
	c := prepareCounter()
	c2 := prepareCounter()

	c.Subtract(c2)

	assert.Equal(t, c.Total(), 0)

	c3 := NewCounter("xd", "abc", "abc", "zzz")
	c.Subtract(c3)

	assert.Equal(t, c.Total(), -4)
}

func TestUpdateFromMap(t *testing.T) {
	c := prepareCounter()

	m := make(map[string]int)
	m["5"] = 6
	m["a"] = 2

	c.UpdateFromMap(m)

	assert.Equal(t, c.Total(), 14)
	assert.Equal(t, c.MostCommon(2), []CounterElem{{Item: "5", Count: 6}, {Item: "a", Count: 5}})
}

func TestUpdate(t *testing.T) {
	c := prepareCounter()
	c2 := prepareCounter()
	c2.Add("XX")
	c2.Add("XX")
	c2.Add("XX")
	c2.Add("XX")
	c2.Add("XX")

	c.Update(c2)

	assert.Equal(t, c.Total(), 17)
	assert.Equal(t, c.MostCommon(2), []CounterElem{{Item: "a", Count: 6}, {Item: "XX", Count: 5}})
}
