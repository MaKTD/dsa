package darray

import (
	"testing"
)

func TestNew(t *testing.T) {
	darr := New(1, 2, 3, 4, 5)

	for i := 0; i < 5; i++ {
		elem := darr.Get(i)
		if elem != i+1 {
			t.Fatalf("expected darr to have elem %v, got %v", i, elem)
		}
	}
	t.Logf("darr New pass")
}

func TestLen(t *testing.T) {
	darr := New[int]()
	darr2 := New(1)
	darr3 := New(1, 2, 3, 4)

	if darr.Len() != 0 {
		t.Fatalf("expected darr to have len %v, got %v", 0, darr.Len())
	}
	if darr2.Len() != 1 {
		t.Fatalf("expected darr to have len %v, got %v", 1, darr2.Len())
	}
	if darr3.Len() != 4 {
		t.Fatalf("expected darr to have len %v, got %v", 4, darr3.Len())
	}

	t.Logf("darr Len pass")
}

func TestDynamicArray_IsEmpty(t *testing.T) {
	darr := New[int]()
	darr2 := New(1)
	darr3 := New("A", "C", "D")

	if darr.IsEmpty() != true {
		t.Fatalf("expected darr to be empty")
	}
	if darr2.IsEmpty() != false {
		t.Fatalf("expected darr NOT to be empty")
	}
	if darr3.IsEmpty() != false {
		t.Fatalf("expected darr NOT to be empty")
	}
	t.Logf("darr IsEmpty pass")
}
