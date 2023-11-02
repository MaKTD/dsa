package unionfind

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	f := NewInt(0)

	require.Equal(t, f.data, make([]int, 0))
	require.Equal(t, f.Count(), 0)

	f = NewInt(1)

	require.Equal(t, f.data, []int{0})
	require.Equal(t, f.Count(), 1)

	f = NewInt(10)

	require.Equal(t, f.data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	require.Equal(t, f.Count(), 10)
}

func TestIntUnionFind_Find(t *testing.T) {
	f := NewInt(0)

	setId, ok := f.Find(0)
	require.Equal(t, setId, 0)
	require.Equal(t, ok, false)

	setId, ok = f.Find(1)
	require.Equal(t, setId, 0)
	require.Equal(t, ok, false)

	setId, ok = f.Find(10)
	require.Equal(t, setId, 0)
	require.Equal(t, ok, false)

	f = NewInt(10)

	setId, ok = f.Find(-1)
	require.Equal(t, setId, 0)
	require.Equal(t, ok, false)

	for i := 0; i < 10; i++ {
		setId, ok = f.Find(i)

		require.Equal(t, i, setId)
		require.Equal(t, true, ok)
	}

	f = IntUnionFind{
		data: []int{0, 0, 1, 2, 3, 4, 5, 7, 0, 7},
	}

	setId, ok = f.Find(5)
	require.Equal(t, 0, setId)
	require.Equal(t, true, ok)

	setId, ok = f.Find(8)
	require.Equal(t, 0, setId)
	require.Equal(t, true, ok)

	setId, ok = f.Find(9)
	require.Equal(t, 7, setId)
	require.Equal(t, true, ok)

	f = IntUnionFind{
		data: []int{0, 0, 1, 2, 3, 4, 5, 7, 0, 7},
	}

	setId, ok = f.Find(5)
	require.Equal(t, f.data, []int{0, 0, 0, 0, 0, 0, 5, 7, 0, 7})
	require.Equal(t, setId, 0)
	require.Equal(t, true, ok)

	f = IntUnionFind{
		data: []int{4, 4, 4, 13, 4, 4, 1, 13, 10, 3, 4, 12, 5, 8},
	}
	setId, ok = f.Find(3)
	require.Equal(t, 4, setId)
	require.Equal(t, true, ok)
	require.Equal(t, f.data, []int{4, 4, 4, 4, 4, 4, 1, 13, 4, 3, 4, 12, 5, 4})
}

func TestIntUnionFind_Unite(t *testing.T) {
	f := NewInt(10)

	f.Unite(-1, 1)

	require.Equal(t, f.data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	f.Unite(0, -1)
	require.Equal(t, f.data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	f.Unite(0, 0)
	require.Equal(t, f.data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	f.Unite(5, 5)
	require.Equal(t, f.data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	f.Unite(0, 1)
	require.Equal(t, f.IsConnected(0, 1), true)

	f.Unite(1, 2)
	require.Equal(t, f.IsConnected(1, 2), true)
	require.Equal(t, f.IsConnected(0, 2), true)

	f.Unite(3, 5)
	require.Equal(t, f.IsConnected(3, 5), true)

	f.Unite(5, 8)
	require.Equal(t, f.IsConnected(5, 8), true)
	require.Equal(t, f.IsConnected(3, 8), true)

	f.Unite(8, 6)
	require.Equal(t, f.IsConnected(6, 8), true)
	require.Equal(t, f.IsConnected(6, 5), true)
	require.Equal(t, f.IsConnected(6, 3), true)

	f.Unite(6, 5)
	require.Equal(t, f.IsConnected(6, 5), true)
	require.Equal(t, f.IsConnected(3, 8), true)
	require.Equal(t, f.IsConnected(5, 8), true)
	require.Equal(t, f.IsConnected(3, 6), true)
}
