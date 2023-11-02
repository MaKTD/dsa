package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue_resize(t *testing.T) {
	queue := Queue[int]{}

	require.Equal(t, 0, queue.total)
	require.Equal(t, 0, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, 0, len(queue.s))
	require.Equal(t, 0, cap(queue.s))

	queue.resize()

	require.Equal(t, 0, queue.total)
	require.Equal(t, 0, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, 2, len(queue.s))
	require.Equal(t, 2, cap(queue.s))

	queue = Queue[int]{
		tail:  3,
		head:  0,
		s:     []int{1, 2, 3},
		total: 3,
	}

	queue.resize()
	require.Equal(t, 6, len(queue.s))
	require.Equal(t, 0, queue.head)
	require.Equal(t, 3, queue.tail)
	require.Equal(t, []int{1, 2, 3, 0, 0, 0}, queue.s)

	queue = Queue[int]{
		tail:  2,
		head:  3,
		s:     []int{1, 2, 3, 4, 5},
		total: 4,
	}

	queue.resize()
	require.Equal(t, 8, len(queue.s))
	require.Equal(t, 0, queue.head)
	require.Equal(t, 4, queue.tail)
	require.Equal(t, []int{4, 5, 1, 2, 0, 0, 0, 0}, queue.s)
}

func TestQueue_Enqueue(t *testing.T) {
	queue := Queue[int]{}

	queue.Enqueue(1)
	require.Equal(t, 1, queue.total)
	require.Equal(t, 1, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 0}, queue.s)

	queue.Enqueue(2)
	require.Equal(t, 2, queue.total)
	require.Equal(t, 0, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 3, queue.total)
	require.Equal(t, 3, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 0}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 4, queue.total)
	require.Equal(t, 0, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 3}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 5, queue.total)
	require.Equal(t, 5, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 3, 3, 0, 0, 0}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 6, queue.total)
	require.Equal(t, 6, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 3, 3, 3, 0, 0}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 7, queue.total)
	require.Equal(t, 7, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 3, 3, 3, 3, 0}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 8, queue.total)
	require.Equal(t, 0, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 3, 3, 3, 3, 3}, queue.s)

	queue.Enqueue(3)
	require.Equal(t, 9, queue.total)
	require.Equal(t, 9, queue.tail)
	require.Equal(t, 0, queue.head)
	require.Equal(t, []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0}, queue.s)
}

func TestQueue_Peek(t *testing.T) {
	queue := Queue[int]{}

	elem, ok := queue.Peek()
	require.Equal(t, false, ok)
	require.Equal(t, 0, elem)

	queue.Enqueue(1)
	elem, ok = queue.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, 1, elem)

	queue.Enqueue(2)
	elem, ok = queue.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, 1, elem)

	queue.Enqueue(3)
	elem, ok = queue.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, 1, elem)
}

func TestQueue_Dequeue(t *testing.T) {
	queue := Queue[int]{}

	elem, ok := queue.Dequeue()
	require.Equal(t, false, ok)
	require.Equal(t, 0, elem)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 1, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 2, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 3, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 4, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 5, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 6, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, true, ok)
	require.Equal(t, 7, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, false, ok)
	require.Equal(t, 0, elem)

	elem, ok = queue.Dequeue()
	require.Equal(t, false, ok)
	require.Equal(t, 0, elem)
}
