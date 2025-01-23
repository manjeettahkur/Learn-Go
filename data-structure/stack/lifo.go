package stack

import "errors"

type Stack[T comparable] struct {
	Items []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.Items = append(s.Items, val)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack[T]) Pop(val T) (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	return s.Items[len(s.Items)-1], nil

}
