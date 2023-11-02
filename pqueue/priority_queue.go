package pqueue

import "github.com/MaKTD/go-dsa/bheap"

// QueueMin simple renaming methods from binary heap
type QueueMin[T bheap.BinaryHeapElement[T]] struct {
	h bheap.BinaryHeap[T]
}

func (r *QueueMin[T]) IsEmpty() bool {
	return r.h.IsEmpty()
}

func (r *QueueMin[T]) Size() int {
	return r.h.Size()
}

func (r *QueueMin[T]) Enqueue(elem T) {
	r.h.Add(elem)
}

func (r *QueueMin[T]) Peek() (T, bool) {
	return r.h.Peek()
}

func (r *QueueMin[T]) Dequeue() (T, bool) {
	return r.h.Poll()
}

func New[T bheap.BinaryHeapElement[T]](elems ...T) QueueMin[T] {
	h := bheap.Heapify(elems)
	q := QueueMin[T]{h}

	return q
}
