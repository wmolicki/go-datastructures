package set

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	s := Set[T]{data: make(map[T]struct{})}
	for _, item := range items {
		s.Add(item)
	}
	return &s
}

func (s *Set[T]) Items() []T {
	r := make([]T, 0, len(s.data))

	for k := range s.data {
		r = append(r, k)
	}

	return r
}

func (s *Set[T]) Add(item T) {
	s.data[item] = struct{}{}
}

func (s *Set[T]) In(item T) bool {
	_, ok := s.data[item]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) Update(sets ...*Set[T]) {
	for _, other := range sets {
		for k := range other.data {
			s.Add(k)
		}
	}
}

func (s *Set[T]) Remove(item T) {
	delete(s.data, item)
}

// Difference returns a new set with elements in the set that are not in the others.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for k := range s.data {
		_, ok := other.data[k]
		if ok {
			continue
		}
		result.Add(k)
	}

	return result
}

// Intersection returns a new set with elements in both sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for k := range s.data {
		_, ok := other.data[k]
		if !ok {
			continue
		}
		result.Add(k)
	}

	return result
}

// IsSubsetOf returns whether a set is a subset of other set
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	// all my keys must be in other
	for k := range s.data {
		_, ok := other.data[k]
		if !ok {
			return false
		}
	}

	return true
}

// IsSupersetOf returns whether a set is a superset of other set
func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {
	// if other set is a subset of me, it means I am the superset of other
	return other.IsSubsetOf(s)
}

func (s *Set[T]) Copy() *Set[T] {
	c := Set[T]{data: make(map[T]struct{}, len(s.data))}
	for k, v := range s.data {
		c.data[k] = v
	}
	return &c
}

// Union returns a new set with elements from both sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := s.Copy()

	for k := range other.data {
		result.Add(k)
	}

	return result
}
