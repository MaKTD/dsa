package unionfind

import "math/rand"

type IntUnionFind struct {
	data []int
}

func (r *IntUnionFind) Find(x int) (int, bool) {
	if x < 0 || (len(r.data)-1) < x {
		return 0, false
	}

	root := x
	for r.data[root] != root {
		root = r.data[root]
	}

	// optimize tree
	for r.data[x] != x {
		next := r.data[x]
		r.data[x] = root
		x = next
	}

	return root, true
}

func (r *IntUnionFind) IsConnected(x, y int) bool {
	idX, okX := r.Find(x)
	idY, okY := r.Find(y)
	if !okX || !okY {
		return false
	}
	return idX == idY
}

func (r *IntUnionFind) Unite(x, y int) {
	rootX, okX := r.Find(x)
	rootY, okY := r.Find(y)

	if !okX || !okY {
		return
	}

	if rootX == rootY {
		return
	}

	if rand.Int()%2 == 0 {
		r.data[rootY] = rootX
	} else {
		r.data[rootX] = rootY
	}
}

func (r *IntUnionFind) Count() int {
	return len(r.data)
}

func NewInt(n int) IntUnionFind {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return IntUnionFind{
		data: slice,
	}
}
