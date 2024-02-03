package stack

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	Top  *Node[T]
	Size int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.Size++
	if s.Top == nil {
		s.Top = &Node[T]{Value: v}
		return
	}
	c := s.Top
	s.Top = &Node[T]{Value: v}
	s.Top.Next = c
}

func (s *Stack[T]) Pop() *T {
	if s.Top == nil {
		return nil
	}
	s.Size--
	v := s.Top.Value
	s.Top = s.Top.Next
	return &v
}

func (s *Stack[T]) Peek() *T {
	if s.Top == nil {
		return nil
	}
	return &s.Top.Value
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack[T]) Show() (in []T) {
	current := s.Top
	for current != nil {
		in = append(in, current.Value)
		current = current.Next
	}
	return
}
