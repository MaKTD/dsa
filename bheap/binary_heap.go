package bheap

type BinaryHeapElement[T any] interface {
	Less(another T) bool
}

type BinaryHeap[T BinaryHeapElement[T]] struct {
	data []T
}

func (r *BinaryHeap[T]) Size() int {
	return len(r.data)
}

func (r *BinaryHeap[T]) IsEmpty() bool {
	return len(r.data) == 0
}

func (r *BinaryHeap[T]) Add(elem T) {
	r.data = append(r.data, elem)
	r.siftUp(len(r.data) - 1)
}

func (r *BinaryHeap[T]) Peek() (T, bool) {
	if len(r.data) == 0 {
		var elem T
		return elem, false
	}

	return r.data[0], true
}

func (r *BinaryHeap[T]) Poll() (T, bool) {
	if len(r.data) == 0 {
		var elem T
		return elem, false
	}

	min := r.data[0]
	r.data[0], r.data[len(r.data)-1] = r.data[len(r.data)-1], r.data[0]
	r.data = r.data[:len(r.data)-1]
	r.siftDown(0)
	return min, true
}

func (r *BinaryHeap[T]) siftUp(i int) {
	for i != 0 {
		parentIdx := (i - 1) / 2
		if r.data[i].Less(r.data[parentIdx]) {
			r.data[parentIdx], r.data[i] = r.data[i], r.data[parentIdx]
			i = parentIdx
		} else {
			break
		}
	}
}

func (r *BinaryHeap[T]) siftDown(i int) {
	for 2*i+1 < len(r.data) {
		childIdx := 2*i + 1
		if childIdx+1 < len(r.data) && r.data[childIdx+1].Less(r.data[childIdx]) {
			childIdx += 1
		}
		if r.data[childIdx].Less(r.data[i]) {
			r.data[i], r.data[childIdx] = r.data[childIdx], r.data[i]
			i = childIdx
		} else {
			break
		}
	}
}

func New[T BinaryHeapElement[T]](elems ...T) BinaryHeap[T] {
	h := BinaryHeap[T]{}

	for _, v := range elems {
		h.Add(v)
	}

	return h
}
