package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}
	ix := len(s.data) - 1
	item := s.data[ix]
	s.data = s.data[:ix]
	return item, true
}
