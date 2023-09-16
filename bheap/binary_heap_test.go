package bheap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Elem int

func (r Elem) Less(b Elem) bool {
	return r < b
}

func TestBinaryHeapSiftUp(t *testing.T) {
	h := BinaryHeap[Elem]{
		data: []Elem{3, 5, 7, 2},
	}
	h.siftUp(3)

	require.Equal(t, []Elem{2, 3, 7, 5}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{1, 3, 5, 9, 12, 8, 9, 11, 12, 20, 21, 35, 50, 42, 48, 2, 16},
	}
	h.siftUp(15)

	require.Equal(t, []Elem{1, 2, 5, 3, 12, 8, 9, 9, 12, 20, 21, 35, 50, 42, 48, 11, 16}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{2, 5, 8, 9, 1},
	}
	h.siftUp(4)
	require.Equal(t, []Elem{1, 2, 8, 9, 5}, h.data)
}

func TestBinaryHeapSiftDown(t *testing.T) {
	h := BinaryHeap[Elem]{
		data: []Elem{2, 50, 7, 9, 12, 8},
	}
	h.siftDown(1)
	require.Equal(t, []Elem{2, 9, 7, 50, 12, 8}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{2, 5, 50, 9, 12, 8},
	}
	h.siftDown(2)
	require.Equal(t, []Elem{2, 5, 8, 9, 12, 50}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{2, 5, 50, 9, 12, 8},
	}
	h.siftDown(0)
	require.Equal(t, []Elem{2, 5, 50, 9, 12, 8}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{2, 5, 50, 9, 12, 8},
	}
	h.siftDown(3)
	require.Equal(t, []Elem{2, 5, 50, 9, 12, 8}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{2, 5, 50, 9, 12, 8},
	}
	h.siftDown(5)
	require.Equal(t, []Elem{2, 5, 50, 9, 12, 8}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{5},
	}
	h.siftDown(0)
	require.Equal(t, []Elem{5}, h.data)

	h = BinaryHeap[Elem]{
		data: []Elem{50, 8, 16, 20, 21, 18, 20, 41, 33},
	}
	h.siftDown(0)
	require.Equal(t, []Elem{8, 20, 16, 33, 21, 18, 20, 41, 50}, h.data)
}

func TestHeap_Add(t *testing.T) {
	h := BinaryHeap[Elem]{}

	h.Add(1)
	require.Equal(t, []Elem{1}, h.data)

	h.Add(2)
	require.Equal(t, []Elem{1, 2}, h.data)

	h.Add(5)
	require.Equal(t, []Elem{1, 2, 5}, h.data)

	h.Add(10)
	require.Equal(t, []Elem{1, 2, 5, 10}, h.data)

	h.Add(3)
	require.Equal(t, []Elem{1, 2, 5, 10, 3}, h.data)

	h.Add(20)
	require.Equal(t, []Elem{1, 2, 5, 10, 3, 20}, h.data)

	h.Add(50)
	require.Equal(t, []Elem{1, 2, 5, 10, 3, 20, 50}, h.data)

	h.Add(4)
	require.Equal(t, []Elem{1, 2, 5, 4, 3, 20, 50, 10}, h.data)

	h.Add(1)
	require.Equal(t, []Elem{1, 1, 5, 2, 3, 20, 50, 10, 4}, h.data)

	h.Add(0)
	require.Equal(t, []Elem{0, 1, 5, 2, 1, 20, 50, 10, 4, 3}, h.data)

	h.Add(13)
	require.Equal(t, []Elem{0, 1, 5, 2, 1, 20, 50, 10, 4, 3, 13}, h.data)

	h.Add(4)
	require.Equal(t, []Elem{0, 1, 4, 2, 1, 5, 50, 10, 4, 3, 13, 20}, h.data)
}

func TestHeap_Peek(t *testing.T) {
	h := BinaryHeap[Elem]{}

	min, ok := h.Peek()
	require.Equal(t, false, ok)
	require.Equal(t, Elem(0), min)

	h.Add(1000)
	min, ok = h.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(1000), min)

	h.Add(1100)
	min, ok = h.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(1000), min)

	h.Add(100)
	min, ok = h.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(100), min)

	h.Add(80)
	h.Add(20)
	min, ok = h.Peek()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(20), min)
}

func TestBinaryHeap_Poll(t *testing.T) {
	h := BinaryHeap[Elem]{}

	min, ok := h.Poll()
	require.Equal(t, false, ok)
	require.Equal(t, Elem(0), min)

	h.Add(0)
	min, ok = h.Poll()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(0), min)
	require.Equal(t, 0, len(h.data))

	h.Add(1)
	h.Add(2)
	h.Add(10)
	h.Add(0)
	min, ok = h.Poll()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(0), min)

	min, ok = h.Poll()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(1), min)

	min, ok = h.Poll()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(2), min)

	min, ok = h.Poll()
	require.Equal(t, true, ok)
	require.Equal(t, Elem(10), min)

	min, ok = h.Poll()
	require.Equal(t, false, ok)
	require.Equal(t, Elem(0), min)
}

func TestNew(t *testing.T) {
	elems := []Elem{10, 2, 5, 2, 0, 10, 20, -1, 8, 19, 1, 1}
	h := New(elems...)

	require.Equal(t, []Elem{-1, 0, 1, 2, 1, 5, 20, 10, 8, 19, 2, 10}, h.data)
}

func TestHeapify(t *testing.T) {
	elems := []Elem{10, 2, 5, 2, 0, 10, 20, -1, 8, 19, 1, 1}
	h := Heapify(elems)

	require.Equal(t, []Elem{-1, 0, 1, 2, 1, 5, 20, 2, 8, 19, 10, 10}, h.data)
}
