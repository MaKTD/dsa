package treap

import (
	"errors"
	"fmt"
	"github.com/MaKTD/go-dsa/bst"
	"math/rand"
)

type Comparable[T any] interface {
	// Compare
	// Equal = 0, Positive > 0, Negative < 0
	Compare(another T) int
	Minus(x int) T
}

type node[T Comparable[T]] struct {
	val   T
	pr    int
	left  *node[T]
	right *node[T]
}

type UniqTreap[T Comparable[T]] struct {
	size int
	root *node[T]
}

func (r *UniqTreap[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *UniqTreap[T]) Size() int {
	return r.size
}

func (r *UniqTreap[T]) Iter(order int) bst.Iterator[T] {
	var stack []*node[T]
	if r.root != nil {
		stack = append(stack, r.root)
	}

	switch order {
	case bst.TravPre:
		return &PreorderIterator[T]{stack: stack}
	case bst.TravIn:
		return &InorderIterator[T]{stack: stack, travNode: r.root}
	case bst.TravPost:
		return &PostorderIterator[T]{prepareStack: stack}
	default:
		panic(errors.New(fmt.Sprintf("can not create iterator with order %d", order)))
	}
}

func (r *UniqTreap[T]) Contains(elem T) bool {
	node, _ := r.find(elem)
	return node != nil
}

func (r *UniqTreap[T]) Insert(elem T) bool {
	if r.Contains(elem) {
		return false
	}
	pr := rand.Int()
	n := &node[T]{
		val: elem,
		pr:  pr,
	}
	L, R := r.split(r.root, elem)
	r.root = r.merge(r.merge(L, n), R)
	return true
}

func (r *UniqTreap[T]) Remove(elem T) bool {
	if !r.Contains(elem) {
		return false
	}
	L, R := r.split(r.root, elem)
	LL, _ := r.split(L, elem.Minus(1))
	r.root = r.merge(LL, R)
	return true
}

func (r *UniqTreap[T]) find(elem T) (*node[T], **node[T]) {
	var travLink **node[T]
	trav := r.root
	for trav != nil {
		compareRes := trav.val.Compare(elem)

		if compareRes == 0 {
			return trav, travLink
		} else if compareRes < 0 {
			travLink = &trav.right
			trav = trav.right
		} else {
			travLink = &trav.left
			trav = trav.left
		}
	}

	return nil, nil
}

func (r *UniqTreap[T]) merge(L, R *node[T]) *node[T] {
	if L == nil {
		return R
	}
	if R == nil {
		return L
	}
	if L.pr > R.pr {
		R.left = r.merge(R.left, L)
		return R
	} else {
		L.right = r.merge(L.right, R)
		return L
	}
}

func (r *UniqTreap[T]) split(t *node[T], x T) (*node[T], *node[T]) {
	if t == nil {
		return nil, nil
	}

	compareRes := t.val.Compare(x)
	if compareRes <= 0 {
		t1, t2 := r.split(t.right, x)
		t.right = t1
		return t, t2
	} else {
		t1, t2 := r.split(t.left, x)
		t.left = t2
		return t1, t
	}
}
