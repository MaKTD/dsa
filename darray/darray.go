package darray

import (
	"errors"
	"fmt"
)

type DynamicArray[T any] struct {
	len    int
	source []T // slice is used because size if part of real arrays and can not be parametrized yet
}

func (r *DynamicArray[T]) Get(i int) T {
	if i > r.len {
		panic(errors.New(fmt.Sprintf("index out of darray range i = %v, darray lenght = %v", i, r.len)))
	}

	return r.source[i]
}

func (r *DynamicArray[T]) Set(i int, v T) {
	if i > r.len {
		panic(errors.New(fmt.Sprintf("index out of darray range i = %v, darray lenght = %v", i, r.len)))
	}

	r.source[i] = v
}

func (r *DynamicArray[T]) Append(v T) {

	if len(r.source)+1 > cap(r.source) {
		source := make([]T, r.len*2)
		r.source = source
	}

	r.source[r.len+1] = v
	r.len += 1
}

func (r *DynamicArray[T]) Clear() {
	source := make([]T, len(r.source))
	r.source = source
	r.len = 0
}

func (r *DynamicArray[T]) IsEmpty() bool {
	return r.len == 0
}

func (r *DynamicArray[T]) Len() int {
	return r.len
}

func New[T any](elems ...T) DynamicArray[T] {
	var length int
	var capacity int

	if len(elems) == 0 {
		length = 0
		capacity = 2
	} else if len(elems) == 1 {
		length = 1
		capacity = 1
	} else {
		length = len(elems)
		capacity = len(elems)
	}

	source := make([]T, length, capacity)
	if len(elems) > 0 {
		for i, v := range elems {
			source[i] = v
		}
	}

	return DynamicArray[T]{
		len:    length,
		source: source,
	}
}
