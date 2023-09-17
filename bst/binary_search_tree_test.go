package bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Elem int

func (r Elem) Compare(another Elem) int {
	return int(r) - int(another)
}

func TestUniqTree_Insert(t *testing.T) {
	tree := UniqTree[Elem]{}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(11)

	iter := tree.Iter(TravPre)
	var seq []Elem
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{10, 8, 6, 9, 20, 12, 11, 30}, seq)

	tree = UniqTree[Elem]{}
	tree.Insert(1)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(6)

	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{1, 2, 3, 4, 5, 6, 7, 8}, seq)
}

func TestUniqTree_Contains(t *testing.T) {
	tree := UniqTree[Elem]{}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(11)

	require.Equal(t, true, tree.Contains(11))
	require.Equal(t, true, tree.Contains(10))
	require.Equal(t, true, tree.Contains(6))
	require.Equal(t, false, tree.Contains(-6))
	require.Equal(t, false, tree.Contains(1))
	require.Equal(t, false, tree.Contains(0))
	require.Equal(t, false, tree.Contains(31))
}

func TestUniqTree_Remove(t *testing.T) {
	tree := UniqTree[Elem]{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(11)

	require.Equal(t, false, tree.Remove(-10))
	require.Equal(t, false, tree.Remove(-20))
	require.Equal(t, false, tree.Remove(5))
	require.Equal(t, 9, tree.Size())

	require.Equal(t, true, tree.Remove(9))
	require.Equal(t, 8, tree.Size())

	iter := tree.Iter(TravIn)
	var seq []Elem
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{6, 8, 10, 11, 12, 20, 30}, seq)

	require.Equal(t, true, tree.Remove(30))
	require.Equal(t, 7, tree.Size())

	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{6, 8, 10, 11, 12, 20}, seq)

	require.Equal(t, true, tree.Remove(11))
	require.Equal(t, 6, tree.Size())

	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{6, 8, 10, 12, 20}, seq)

	require.Equal(t, true, tree.Remove(6))
	require.Equal(t, 5, tree.Size())

	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{8, 10, 12, 20}, seq)

	tree = UniqTree[Elem]{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(4)

	require.Equal(t, true, tree.Remove(6))
	require.Equal(t, 9, tree.Size())
	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{4, 8, 9, 10, 11, 12, 20, 30}, seq)

	require.Equal(t, true, tree.Remove(12))
	require.Equal(t, 8, tree.Size())
	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{4, 8, 9, 10, 11, 20, 30}, seq)

	tree = UniqTree[Elem]{}
	tree.Insert(10)

	require.Equal(t, true, tree.Remove(10))
	require.Equal(t, 0, tree.Size())

	tree = UniqTree[Elem]{}
	tree.Insert(10)
	tree.Insert(20)

	require.Equal(t, true, tree.Remove(10))
	require.Equal(t, 1, tree.Size())

	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{20}, seq)

	tree = UniqTree[Elem]{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(4)

	require.Equal(t, true, tree.Remove(10))
	require.Equal(t, 9, tree.Size())
	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{4, 6, 8, 9, 11, 12, 20, 30}, seq)

	require.Equal(t, true, tree.Remove(20))
	require.Equal(t, 8, tree.Size())
	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{4, 6, 8, 9, 11, 12, 30}, seq)

	require.Equal(t, true, tree.Remove(8))
	require.Equal(t, 7, tree.Size())
	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{4, 6, 9, 11, 12, 30}, seq)
}

func TestUniqTree_Iter(t *testing.T) {
	tree := UniqTree[Elem]{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(11)

	iter := tree.Iter(TravPre)
	var seq []Elem
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{10, 8, 6, 9, 20, 12, 11, 30}, seq)

	iter = tree.Iter(TravIn)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{6, 8, 9, 10, 11, 12, 20, 30}, seq)

	iter = tree.Iter(TravPost)
	seq = nil
	for iter.Next() {
		seq = append(seq, iter.Value())
	}
	require.Equal(t, []Elem{6, 9, 8, 11, 12, 30, 20, 10}, seq)
}
