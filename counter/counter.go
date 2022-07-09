package counter

import (
	"sort"
)

type Counter struct {
	data map[string]int
}

func NewCounter(items ...string) *Counter {
	c := Counter{data: make(map[string]int)}
	for _, item := range items {
		c.Add(item)
	}
	return &c
}

func (c *Counter) Add(item string) {
	c.data[item] += 1
}

type CounterElem struct {
	Item  string
	Count int
}

func (c *Counter) MostCommon(n int) []CounterElem {
	elemCounts := []CounterElem{}

	for k, v := range c.data {
		elemCounts = append(elemCounts, CounterElem{Item: k, Count: v})
	}

	sort.Slice(elemCounts, func(a, b int) bool {
		return elemCounts[a].Count > elemCounts[b].Count
	})

	return elemCounts[:n]
}

func (c *Counter) Total() int {
	s := 0
	for _, v := range c.data {
		s += v
	}
	return s
}

func (c *Counter) Subtract(other *Counter) {
	for k, v := range other.data {
		c.data[k] -= v
	}
}

func (c *Counter) UpdateFromMap(m map[string]int) {
	for k, v := range m {
		c.data[k] += v
	}
}

func (c *Counter) Update(other *Counter) {
	c.UpdateFromMap(other.data)
}
