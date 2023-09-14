package hsort

import "github.com/MaKTD/go-dsa/bheap"

func HeapSort[T bheap.BinaryHeapElement[T]](slice []T) {
	h := bheap.New(slice...)

	for i := 0; i < len(slice); i++ {
		min, _ := h.Poll()
		slice[i] = min
	}
}
