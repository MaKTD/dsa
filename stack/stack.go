package stack

import (
	"errors"
	"github.com/MaKTD/go-dsa/linkedlist"
)

type Stack[T any] struct {
	list linkedlist.SingleList[T]
}

func (s *Stack[T]) Pop() (T, error) {
	var elem T

	if s.list.IsEmpty() {
		return elem, errors.New("can not pop empty stack")
	}

	elem, _ = s.list.RemoveFirst()
	return elem, nil
}

func (s *Stack[T]) Push(elem T) {
	s.list.PushBack(elem)
}

func (s *Stack[T]) Peek() (T, error) {
	var elem T

	if s.list.IsEmpty() {
		return elem, errors.New("can not peek empty stack")
	}

	elem, _ = s.list.First()
	return elem, nil
}

func (s *Stack[T]) Size() int {
	return s.list.Len()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func NewStack[T any](elems ...T) Stack[T] {
	stack := Stack[T]{
		list: linkedlist.NewSingle[T](),
	}

	for _, v := range elems {
		stack.Push(v)
	}

	return stack
}
