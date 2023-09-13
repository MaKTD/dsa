package queue

const minSizeToCollapse = 16

type Queue[T any] struct {
	total int
	tail  int
	head  int
	s     []T
}

func (r *Queue[T]) Size() int {
	return r.total
}

func (r *Queue[T]) Enqueue(elem T) {
	if r.total == len(r.s) {
		r.resize()
	}

	r.s[r.tail] = elem
	r.tail = (r.tail + 1) & (len(r.s) - 1)
	r.total += 1
}

func (r *Queue[T]) Peek() (T, bool) {
	var elem T
	if r.total == 0 {
		return elem, false
	} else {
		return r.s[r.head], true
	}
}

func (r *Queue[T]) Dequeue() (T, bool) {
	var zeroVal T

	if r.total == 0 {
		return zeroVal, false
	}

	elem := r.s[r.head]
	r.s[r.head] = zeroVal

	r.head = (r.head + 1) & (len(r.s) - 1)
	r.total -= 1

	// Resize down if buffer 1/4 full.
	if len(r.s) > minSizeToCollapse && r.total<<2 == len(r.s) {
		r.resize()
	}

	return elem, true
}

func (r *Queue[T]) resize() {
	var newS []T
	if r.total == 0 {
		newS = make([]T, 2)
	} else {
		newS = make([]T, r.total*2)
	}

	if r.tail > r.head {
		copy(newS, r.s[r.head:r.tail])
	} else {
		copied := copy(newS, r.s[r.head:])
		copy(newS[copied:], r.s[:r.tail])
	}
	r.tail = r.total
	r.head = 0
	r.s = newS
}
