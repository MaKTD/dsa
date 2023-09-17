package bst

type Iterator[T Comparable[T]] interface {
	Value() T
	Next() bool
}

type PreorderIterator[T Comparable[T]] struct {
	cur   T
	stack []*node[T]
}

func (r *PreorderIterator[T]) Value() T {
	return r.cur
}

func (r *PreorderIterator[T]) Next() bool {
	if len(r.stack) == 0 {
		return false
	}

	n := r.stack[len(r.stack)-1]
	r.cur = n.val
	r.stack = r.stack[:len(r.stack)-1]

	if n.right != nil {
		r.stack = append(r.stack, n.right)
	}
	if n.left != nil {
		r.stack = append(r.stack, n.left)
	}

	return true
}

type InorderIterator[T Comparable[T]] struct {
	cur      T
	travNode *node[T]
	stack    []*node[T]
}

func (r *InorderIterator[T]) Value() T {
	return r.cur
}

func (r *InorderIterator[T]) Next() bool {
	if len(r.stack) == 0 {
		return false
	}

	for r.travNode != nil && r.travNode.left != nil {
		r.stack = append(r.stack, r.travNode.left)
		r.travNode = r.travNode.left
	}

	n := r.stack[len(r.stack)-1]
	r.stack = r.stack[:len(r.stack)-1]

	r.cur = n.val

	if n.right != nil {
		r.stack = append(r.stack, n.right)
		r.travNode = n.right
	}

	return true
}

type PostorderIterator[T Comparable[T]] struct {
	cur          T
	prepareStack []*node[T]
	orderedStack []*node[T]
}

func (r *PostorderIterator[T]) Value() T {
	return r.cur
}

func (r *PostorderIterator[T]) Next() bool {
	if len(r.orderedStack) == 0 && len(r.prepareStack) != 0 {
		for len(r.prepareStack) != 0 {
			n := r.prepareStack[len(r.prepareStack)-1]
			r.prepareStack = r.prepareStack[:len(r.prepareStack)-1]

			r.orderedStack = append(r.orderedStack, n)

			if n.left != nil {
				r.prepareStack = append(r.prepareStack, n.left)
			}
			if n.right != nil {
				r.prepareStack = append(r.prepareStack, n.right)
			}
		}
	}

	if len(r.orderedStack) == 0 {
		return false
	}

	n := r.orderedStack[len(r.orderedStack)-1]
	r.orderedStack = r.orderedStack[:len(r.orderedStack)-1]
	r.cur = n.val

	return true
}
