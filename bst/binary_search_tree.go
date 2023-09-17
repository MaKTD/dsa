package bst

import (
	"errors"
	"fmt"
)

const (
	TravPre  = 0
	TravIn   = 1
	TravPost = 2
)

type Comparable[T any] interface {
	// Compare Equal = 0, Positive = greater, Negative = lesser
	Compare(another T) int
}

type node[T Comparable[T]] struct {
	val   T
	left  *node[T]
	right *node[T]
}

type UniqTree[T Comparable[T]] struct {
	size int
	root *node[T]
}

func (r *UniqTree[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *UniqTree[T]) Size() int {
	return r.size
}

func (r *UniqTree[T]) Iter(order int) Iterator[T] {
	var stack []*node[T]
	if r.root != nil {
		stack = append(stack, r.root)
	}

	switch order {
	case TravPre:
		return &PreorderIterator[T]{stack: stack}
	case TravIn:
		return &InorderIterator[T]{stack: stack, travNode: r.root}
	case TravPost:
		return &PostorderIterator[T]{prepareStack: stack}
	default:
		panic(errors.New(fmt.Sprintf("can not create iterator with order %d", order)))
	}
}

func (r *UniqTree[T]) Insert(elem T) bool {
	n := &node[T]{
		val: elem,
	}
	r.size += 1

	if r.root == nil {
		r.root = n
		return true
	}

	parent := r.root
	for {
		compareRes := elem.Compare(parent.val)

		// if equals do nothing
		if compareRes == 0 {
			return false
		}

		// if element less than parent
		if compareRes < 0 {
			if parent.left != nil {
				parent = parent.left
				continue
			} else {
				parent.left = n
				return true
			}
		}

		// if element greater than parent
		if parent.right != nil {
			parent = parent.right
			continue
		} else {
			parent.right = n
			return true
		}
	}
}

func (r *UniqTree[T]) Contains(elem T) bool {
	node, _ := r.find(elem)
	return node != nil
}

func (r *UniqTree[T]) Remove(elem T) bool {
	node, link := r.find(elem)
	if node == nil {
		return false
	}
	if link == nil {
		link = &r.root
	}

	if node.left == nil && node.right == nil {
		*link = nil
	}
	if node.left != nil && node.right == nil {
		*link = node.left
	}
	if node.right != nil && node.left == nil {
		*link = node.right
	}
	if node.right != nil && node.left != nil {
		min, minLink := r.findMinInSubTree(&node.right, node.right)

		if min.left == nil && min.right == nil {
			*minLink = nil
		} else if min.left != nil && min.right == nil {
			*minLink = min.left
		} else {
			*minLink = min.right
		}
		min.left = node.left
		min.right = node.right
		*link = min
	}

	r.size -= 1

	return true
}

func (r *UniqTree[T]) find(elem T) (*node[T], **node[T]) {
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

func (r *UniqTree[T]) findMinInSubTree(link **node[T], n *node[T]) (*node[T], **node[T]) {
	for n.left != nil {
		link = &n.left
		n = n.left
	}
	return n, link
}
