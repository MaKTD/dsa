package linkedlist

import (
	"errors"
	"fmt"
)

type SingleListIterator[T any] struct {
	idx int
	cur *SingleNode[T]
}

func (s *SingleListIterator[T]) Value() (int, T) {
	return s.idx, s.cur.Value
}

func (s *SingleListIterator[T]) Next() bool {
	if s.cur == nil || s.cur.next == nil {
		return false
	}
	if s.idx == -1 {
		s.idx += 1

		return true
	}

	s.cur = s.cur.next
	s.idx += 1
	return true
}

type SingleNode[T any] struct {
	next  *SingleNode[T]
	Value T
}

type SingleList[T any] struct {
	len  int
	head *SingleNode[T]
	tail *SingleNode[T]
}

func (s *SingleList[T]) Len() int {
	return s.len
}

func (s *SingleList[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *SingleList[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.len = 0
}

func (s *SingleList[T]) PushFront(value T) int {
	node := SingleNode[T]{
		next:  nil,
		Value: value,
	}

	if s.len == 0 {
		s.head = &node
		s.tail = &node
	} else if s.len == 1 {
		s.head.next = &node
		s.tail = &node
	} else {
		s.tail.next = &node
		s.tail = &node
	}
	s.len += 1

	return s.len - 1
}

func (s *SingleList[T]) PushBack(value T) int {
	node := SingleNode[T]{
		next:  nil,
		Value: value,
	}

	if s.len == 0 {
		s.head = &node
		s.tail = &node
	} else if s.len == 1 {
		node.next = s.head
		s.head = &node
	} else {
		node.next = s.head
		s.head = &node
	}
	s.len += 1

	return s.len - 1
}

func (s *SingleList[T]) RemoveFirst() (T, error) {
	return s.remove(0)
}

func (s *SingleList[T]) RemoveLast() (T, error) {
	return s.remove(s.len - 1)
}

func (s *SingleList[T]) RemoveAt(idx int) (T, error) {
	return s.remove(idx)
}

func (s *SingleList[T]) remove(idx int) (T, error) {
	var removed T
	if s.head == nil {
		return removed, errors.New("can not remove element on empty list")
	}
	if idx >= s.Len() {
		return removed, errors.New(fmt.Sprintf("can not remove element on index %v, index out of range", idx))
	}

	if idx == 0 {
		removed = s.head.Value

		if s.Len() == 1 {
			s.head = nil
			s.tail = nil
		} else {
			s.head = s.head.next
		}
		s.len -= 1

		return removed, nil
	}

	for tNode, i := s.head, 0; tNode != nil; tNode, i = tNode.next, i+1 {
		if idx-1 != i {
			continue
		}

		removed = tNode.next.Value
		if tNode.next == s.tail {
			s.tail = tNode
		} else {
			tNode.next = tNode.next.next
		}
		s.len -= 1
		return removed, nil
	}

	panic("SingleList.remove unexpected nodes traversal loop end")
}

func (s *SingleList[T]) IndexOf(f func(i int, elem T, list *SingleList[T]) bool) (int, bool) {
	for tNode, i := s.head, 0; tNode != nil; tNode, i = tNode.next, i+1 {
		if f(i, tNode.Value, s) {
			return i, true
		}
	}

	return -1, false
}

func (s *SingleList[T]) ForEach(f func(i int, elem T, list *SingleList[T])) {
	cur := s.head
	for i := 0; i < s.len; i++ {
		f(i, cur.Value, s)
		cur = cur.next
	}
}

func (s *SingleList[T]) Iter() SingleListIterator[T] {
	return SingleListIterator[T]{
		cur: s.head,
		idx: -1,
	}
}

func (s *SingleList[T]) Last() (T, bool) {
	return s.tail.Value, s.len != 0
}

func (s *SingleList[T]) First() (T, bool) {
	return s.head.Value, s.len != 0
}

func NewSingle[T any](elems ...T) SingleList[T] {
	list := SingleList[T]{}

	for _, v := range elems {
		list.PushFront(v)
	}

	return list
}
