package treap

import "testing"

type Elem int

func (r Elem) Compare(another Elem) int {
	return int(r) - int(another)
}

func (r Elem) Minus(x int) Elem {
	return Elem(int(r) - x)
}

func Test_Treap(t *testing.T) {
	treap := UniqTreap[Elem]{}
	treap.Insert(1)
	treap.Insert(2)
	treap.Insert(3)
	treap.Insert(4)
	treap.Insert(5)
	treap.Insert(6)
	treap.Insert(7)
	treap.Insert(8)
	treap.Insert(9)
	treap.Insert(10)
	treap.Insert(11)
	treap.Insert(12)
	treap.Insert(13)
	treap.Insert(14)
	treap.Insert(15)
	treap.Insert(16)
	treap.Insert(17)
	treap.Insert(18)

	PrintTree(treap)

	treap.Remove(5)
	treap.Remove(4)
	treap.Remove(18)
	treap.Remove(12)

	PrintTree(treap)

}
