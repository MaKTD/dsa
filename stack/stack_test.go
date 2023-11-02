package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStack(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5, 6, 7, 8}
	stack := NewStack(elems...)

	require.Equal(t, false, stack.IsEmpty())
	require.Equal(t, len(elems), stack.Size())
}

func TestZeroValueStack(t *testing.T) {
	stack := Stack[int]{}

	require.Equal(t, 0, stack.Size())
	require.Equal(t, true, stack.IsEmpty())

	elem, err := stack.Peek()
	require.NotEqual(t, nil, err)
	require.Equal(t, 0, elem)

	elem, err = stack.Pop()
	require.NotEqual(t, nil, err)
	require.Equal(t, 0, elem)

	stack.Push(1)

	elem, err = stack.Peek()
	require.Equal(t, nil, err)
	require.Equal(t, 1, elem)

	elem, err = stack.Pop()
	require.Equal(t, nil, err)
	require.Equal(t, 1, elem)
}
func TestStack(t *testing.T) {
	elems := []string{"A", "B"}
	stack := NewStack(elems...)

	for i := len(elems) - 1; i >= 0; i-- {
		elem, err := stack.Peek()
		require.Equal(t, nil, err)
		require.Equal(t, elems[i], elem)

		elem, err = stack.Pop()
		require.Equal(t, nil, err)
		require.Equal(t, elems[i], elem)
	}

	require.Equal(t, 0, stack.Size())
	require.Equal(t, true, stack.IsEmpty())

	stack.Push("C")
	stack.Push("D")
	stack.Push("F")

	elem, err := stack.Peek()
	require.Equal(t, nil, err)
	require.Equal(t, "F", elem)

	elem, err = stack.Pop()
	require.Equal(t, nil, err)
	require.Equal(t, "F", elem)

	elem, err = stack.Pop()
	require.Equal(t, nil, err)
	require.Equal(t, "D", elem)

	require.Equal(t, 1, stack.Size())

	elem, err = stack.Pop()
	require.Equal(t, nil, err)
	require.Equal(t, "C", elem)

	require.Equal(t, 0, stack.Size())
}
