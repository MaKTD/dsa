package hsort

import "github.com/MaKTD/go-dsa/bheap"

func HeapSort[T bheap.BinaryHeapElement[T]](slice []T) {
	// is is actually possibly to do in place sort if binary heap is max based
	h := bheap.New(slice...)

	for i := 0; i < len(slice); i++ {
		min, _ := h.Poll()
		slice[i] = min
	}
}
