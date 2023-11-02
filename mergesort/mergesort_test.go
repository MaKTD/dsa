package mergesort

import (
	"fmt"
	"testing"
)

func TestMergeIntoSortedSlice(t *testing.T) {
	a := []int{1, 2, 5, 8, 10, 12, 15, 18}
	b := []int{1, 2, 6, 7, 10, 14, 19, 25}

	res := mergeIntoSortedSliceAsc(a, b)
	t.Log(res)
}

func TestMergeSort(t *testing.T) {
	s := []int{1, 10, 3, 2, 5, 29, 12, 58, 123, -10, 1, 52}

	res := MergeSort(s, func(a, b int) bool {
		return a > b
	})

	fmt.Println(res)
}
