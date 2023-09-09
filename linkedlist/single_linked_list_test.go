package linkedlist

import (
	"testing"
)

func TestSingleLinkedListNew(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 20}

	list := NewSingle(elems...)

	list.ForEach(func(i int, elem int, list *SingleList[int]) {
		if elems[i] != elem {
			t.Fatalf("expected element on index %d, to equal %d, got %d", i, elems[i], elem)
		}
	})

	t.Log("TestSingleLinkedListNew pass")
}

func TestSingleListZeroValue(t *testing.T) {
	list := SingleList[int]{}

	if list.IsEmpty() != true {
		t.Fatalf("expected zero value list IsEmpty() to equal true, but got %v", list.IsEmpty())
	}

	if list.Len() != 0 {
		t.Fatalf("expected zero value list Len() to equal 0, but got %v", list.Len())
	}

	list.ForEach(func(i int, elem int, list *SingleList[int]) {
		t.Fatalf("expected ForEach on zero value list to not execute callback")
	})

	elems := []int{1, 2}
	list.PushFront(1)
	list.PushFront(2)

	list.ForEach(func(i int, elem int, list *SingleList[int]) {
		if elems[i] != elem {
			t.Fatalf("expected PushFront on zero list to work, expected elem on index %d, to equal %d, got %d", i, elems[i], elem)
		}
	})
}

func TestSingeListPushFront(t *testing.T) {
	list := SingleList[string]{}

	if list.Len() != 0 {
		t.Fatalf("expected 0 elems list, to have length 0")
	}

	list.PushFront("A")
	if last, ok := list.Last(); !ok || last != "A" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "A", ok, last)
	}
	if first, ok := list.First(); !ok || first != "A" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "A", ok, first)
	}

	list.PushFront("B")
	if last, ok := list.Last(); !ok || last != "B" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "B", ok, last)
	}
	if first, ok := list.First(); !ok || first != "A" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "A", ok, first)
	}

	list.PushFront("C")
	if last, ok := list.Last(); !ok || last != "C" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "C", ok, last)
	}
	if first, ok := list.First(); !ok || first != "A" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "A", ok, first)
	}

	list.PushFront("D")
	if last, ok := list.Last(); !ok || last != "D" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "D", ok, last)
	}
	if first, ok := list.First(); !ok || first != "A" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "A", ok, first)
	}
}

func TestSingleListPushBack(t *testing.T) {
	list := SingleList[string]{}

	if list.Len() != 0 {
		t.Fatalf("expected 0 elems list, to have length 0")
	}

	list.PushBack("A")
	if last, ok := list.Last(); !ok || last != "A" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "A", ok, last)
	}
	if first, ok := list.First(); !ok || first != "A" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "A", ok, first)
	}

	list.PushBack("B")
	if last, ok := list.Last(); !ok || last != "A" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "A", ok, last)
	}
	if first, ok := list.First(); !ok || first != "B" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "B", ok, first)
	}

	list.PushBack("C")
	if last, ok := list.Last(); !ok || last != "A" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "A", ok, last)
	}
	if first, ok := list.First(); !ok || first != "C" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "C", ok, first)
	}

	list.PushBack("D")
	if last, ok := list.Last(); !ok || last != "A" {
		t.Fatalf("expected last value to exists and equal %s, got exists = %v, value = %s", "A", ok, last)
	}
	if first, ok := list.First(); !ok || first != "D" {
		t.Fatalf("expected first value to exists and equal %s, got exists = %v, value = %s", "D", ok, first)
	}
}

func TestSingleListIter(t *testing.T) {
	emptyList := SingleList[int]{}
	emptyIter := emptyList.Iter()
	for emptyIter.Next() {
		t.Fatalf("empty list iterator.Next() should not return true ever")
	}

	elems := []int{1, 2, 3, 10, 2, 3, 5, 8, 50, 100, 60, 20, 9}

	list := NewSingle(elems...)
	iterator := list.Iter()

	totalIterations := 0

	for iterator.Next() {
		i, elem := iterator.Value()

		if i != totalIterations {
			t.Fatalf("expected iterator.Value() to return idx = %v, got %v", totalIterations, i)
		}
		totalIterations += 1

		if elems[i] != elem {
			t.Fatalf("expected iterator.Value() to return elem %v, on idx %v, got %v", elems[i], i, elem)
		}
	}

	if totalIterations != len(elems) {
		t.Fatalf("expected iter to perform %v iterations, but got %v", len(elems), totalIterations)
	}
}

func TestSingleList_RemoveFirst(t *testing.T) {
	elems := []string{"A", "B", "C", "D", "E", "F"}

	list := NewSingle(elems...)

	for i := 0; i < len(elems); i++ {
		removed, err := list.RemoveFirst()

		if err != nil {
			t.Fatalf("expected RemoveFirst() not to error when list not empty, but got %v", err)
		}
		if removed != elems[i] {
			t.Fatalf("expected RemoveFirst() to return %v element, but got %v", elems[i], removed)
		}
		if list.Len() != len(elems)-i-1 {
			t.Fatalf("exepcted list.Len() to be %v, but got %v", len(elems)-i-1, list.Len())
		}

		if i == len(elems)-1 {
			if !list.IsEmpty() {
				t.Fatalf("expected list to be empty when last element is removed")
			}
		} else {
			nowFirst, _ := list.First()
			if elems[i+1] != nowFirst {
				t.Fatalf("expected First element of the list to be %v, but got %v", elems[i+1], nowFirst)
			}
		}
	}

	_, err := list.RemoveFirst()
	if err == nil {
		t.Fatalf("expected RemoveFirst() on empty list to return error, but got %v", err)
	}
}

func TestSingleList_RemoveLast(t *testing.T) {
	elems := []string{"A", "B", "C", "D", "E", "F"}
	list := NewSingle(elems...)

	for i := 0; i < len(elems); i++ {
		if i == len(elems)-1 {
			last, _ := list.Last()
			first, _ := list.First()

			if last != first {
				t.Fatalf("expected last.Last to equal list.First when list become lenght 1 after removing, but got first = %v, last = %v", first, last)
			} else {
				t.Log("when list len should be 1, head and tail should be equal pass")
			}
		}

		removed, err := list.RemoveLast()
		t.Log(removed, err)

		if err != nil {
			t.Fatalf("expected list.RemoveLast() err to be nil, but got %v", err)
		}
		if removed != elems[len(elems)-1-i] {
			t.Fatalf("exepcted list.RemoveLast() removed element to equal %v, but got %v", elems[len(elems)-1-i], removed)
		}

		if list.Len() != len(elems)-i-1 {
			t.Fatalf("expected list to be length of %v after remove, but got %v", len(elems)-i-1, list.Len())
		}

		if i == len(elems)-1 {
			if list.Len() != 0 {
				t.Fatalf("expected list.Len() to be zero 0 after removing last element")
			}
		} else {
			last, _ := list.Last()
			if last != elems[len(elems)-i-1-1] {
				t.Fatalf("expected list.Last() after remove to return %v, but got %v", elems[len(elems)-i-1-1], last)
			}
		}
	}
}

func TestSingleList_RemoveAt(t *testing.T) {
	elems := []string{"A", "B", "C", "D", "E", "F"}
	list := NewSingle(elems...)

	removed, err := list.RemoveAt(3)
	if err != nil {
		t.Fatalf("expected RemoveAt to return nil error, but got %v", err)
	}
	if removed != "D" {
		t.Fatalf("expected RemoveAt to return C, but got %v", removed)
	}
	if list.Len() != 5 {
		t.Fatalf("expected list.Len() to equal 5, after remove, but got %v", list.Len())
	}

	list.ForEach(func(i int, elem string, list *SingleList[string]) {
		if i >= 3 {
			if elem != elems[i+1] {
				t.Fatalf("expected elemeent at index %v to equal %v, but got %v", i, elems[i+1], elem)
			}
		} else {
			if elem != elems[i] {
				t.Fatalf("expected elemeent at index %v to equal %v, but got %v", i, elems[i], elem)
			}
		}
	})
}

func TestSingleList_IndexOfl(t *testing.T) {
	elems := []string{"A", "B", "C", "D", "E", "F"}
	list := NewSingle(elems...)

	idx, ok := list.IndexOf(func(_ int, elem string, _ *SingleList[string]) bool {
		return elem == "B"
	})
	if idx != 1 || ok != true {
		t.Fatalf("expected idx, ok to equal 1, true but got %v, %v", idx, ok)
	}

	idx, ok = list.IndexOf(func(_ int, elem string, _ *SingleList[string]) bool {
		return elem == "F"
	})
	if idx != 5 || ok != true {
		t.Fatalf("expected idx, ok to equal 1, true but got %v, %v", idx, ok)
	}

	idx, ok = list.IndexOf(func(_ int, elem string, _ *SingleList[string]) bool {
		return elem == "Super"
	})
	if idx != -1 || ok != false {
		t.Fatalf("expected idx, ok to equal 1, true but got %v, %v", idx, ok)
	}
}
