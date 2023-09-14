package hsort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Elem int

func (r Elem) Less(b Elem) bool {
	return r < b
}

func TestHeapSort(t *testing.T) {
	slice := []Elem{10, 4, -1, -10, 20, 50, -10, -20, 50, 100, -200, -300, 400, 1, 5, 2, 5, 8, 20}

	HeapSort(slice)

	require.Equal(t, []Elem{-300, -200, -20, -10, -10, -1, 1, 2, 4, 5, 5, 8, 10, 20, 20, 50, 50, 100, 400}, slice)
}
