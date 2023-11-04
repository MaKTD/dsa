package avltree

import (
	"errors"
	"fmt"
	"github.com/MaKTD/go-dsa/bst"
)

type Comparable[T any] interface {
	// Compare
	// Equal = 0, Positive > 0, Negative < 0
	Compare(another T) int
}

type node[T Comparable[T]] struct {
	val    T
	left   *node[T]
	right  *node[T]
	bf     int
	height int
}

type UniqAvlTree[T Comparable[T]] struct {
	size int
	root *node[T]
}

func (r *UniqAvlTree[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *UniqAvlTree[T]) Size() int {
	return r.size
}

func (r *UniqAvlTree[T]) Height() int {
	if r.root == nil {
		return 0
	}
	return r.root.height
}

func (r *UniqAvlTree[T]) Iter(order int) bst.Iterator[T] {
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

func (r *UniqAvlTree[T]) Insert(elem T) bool {
	if !r.Contains(elem) {
		r.root = r.insert(r.root, elem)
		r.size += 1
		return true
	}
	return false
}

func (r *UniqAvlTree[T]) insert(n *node[T], value T) *node[T] {
	if n == nil {
		return &node[T]{val: value}
	}

	compareRes := value.Compare(n.val)

	if compareRes < 0 {
		n.left = r.insert(n.left, value)
	} else {
		n.right = r.insert(n.right, value)
	}

	r.updateHeight(n)
	return r.balance(n)
}

func (r *UniqAvlTree[T]) Contains(elem T) bool {
	node, _ := r.find(elem)
	return node != nil
}

func (r *UniqAvlTree[T]) Remove(elem T) bool {
	if r.Contains(elem) {
		r.root = r.remove(r.root, elem)
		r.size -= 1
		return true
	}

	return false
}

func (r *UniqAvlTree[T]) remove(n *node[T], value T) *node[T] {
	if n == nil {
		return nil
	}

	compareRes := value.Compare(n.val)
	if compareRes < 0 {
		n.left = r.remove(n.left, value)
	} else if compareRes > 0 {
		n.right = r.remove(n.right, value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			if n.left.height > n.right.height {
				val := r.findMax(n.left)
				n.val = val
				n.left = r.remove(n.left, val)
			} else {
				val := r.findMin(n.right)
				n.val = val
				n.right = r.remove(n.right, val)
			}
		}
	}

	r.updateHeight(n)
	return r.balance(n)
}

func (r *UniqAvlTree[T]) find(elem T) (*node[T], **node[T]) {
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

func (r *UniqAvlTree[T]) updateHeight(n *node[T]) {
	leftH, rightH := -1, -1
	if n.left != nil {
		leftH = n.left.height
	}
	if n.right != nil {
		rightH = n.right.height
	}

	if leftH > rightH {
		n.height = leftH + 1
	} else {
		n.height = rightH + 1
	}

	n.bf = rightH - leftH
}

func (r *UniqAvlTree[T]) balance(n *node[T]) *node[T] {
	if n.bf == -2 {
		if n.left.bf <= 0 {
			return r.leftLeftCase(n)
		} else {
			return r.leftRightCase(n)
		}
	} else if n.bf == +2 {
		if n.right.bf >= 0 {
			return r.rightRightCase(n)
		} else {
			return r.rightLeftCase(n)
		}
	}

	return n
}

func (r *UniqAvlTree[T]) leftLeftCase(n *node[T]) *node[T] {
	return r.rightRotation(n)
}

func (r *UniqAvlTree[T]) leftRightCase(n *node[T]) *node[T] {
	n.left = r.leftRotation(n.left)
	return r.leftLeftCase(n)
}

func (r *UniqAvlTree[T]) rightRightCase(n *node[T]) *node[T] {
	return r.leftRotation(n)

}
func (r *UniqAvlTree[T]) rightLeftCase(n *node[T]) *node[T] {
	n.right = r.rightRotation(n.right)
	return r.rightRightCase(n)
}

func (r *UniqAvlTree[T]) rightRotation(n *node[T]) *node[T] {
	le := n.left
	n.left = le.right
	le.right = n
	r.updateHeight(n)
	r.updateHeight(le)
	return le
}

func (r *UniqAvlTree[T]) leftRotation(n *node[T]) *node[T] {
	re := n.right
	n.right = re.left
	re.left = n
	r.updateHeight(n)
	r.updateHeight(re)
	return re
}

func (r *UniqAvlTree[T]) findMin(n *node[T]) T {
	for n.left != nil {
		n = n.left
	}
	return n.val
}

func (r *UniqAvlTree[T]) findMax(n *node[T]) T {
	for n.right != nil {
		n = n.right
	}
	return n.val
}
