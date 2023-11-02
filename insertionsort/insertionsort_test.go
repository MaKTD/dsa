package insertionsort

import (
	"testing"
)

func TestInsertionSortIntAsc(t *testing.T) {
	s := []int{5, 2, 3, 6, 9, 7, 7}
	SortIntAsc(s)

	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			t.Logf("expected %v to be sorted asc", s)
		}
	}
	t.Log("TestInsertionSortIntAsc pass")
}

func TestInsertionSortIntDesc(t *testing.T) {
	s := []int{5, 2, 3, 6, 9, 7, 7}
	SortIntDesc(s)

	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			t.Fatalf("expected %v to be sorted desc", s)
		}
	}
	t.Log("TestInsertionSortIntDesc pass")
}

func TestSortAsc(t *testing.T) {
	ss := []string{"Z", "B", "C", "A", "D", "T"}
	si := []int{5, 2, 3, 6, 9, 7, 7}

	SortAsc(ss)
	SortAsc(si)

	for i := 0; i < len(ss)-1; i++ {
		if ss[i] > ss[i+1] {
			t.Fatalf("expected %v to be sorted asc", ss)
		}
	}

	for i := 0; i < len(si)-1; i++ {
		if si[i] > si[i+1] {
			t.Fatalf("expected %v to be sorted asc", si)
		}
	}

	t.Log("TestSortIntAsc pass")
}

func TestSortDesc(t *testing.T) {
	ss := []string{"Z", "B", "C", "A", "D", "T"}
	si := []int{5, 2, 3, 6, 9, 7, 7}

	SortDesc(ss)
	SortDesc(si)

	for i := 0; i < len(ss)-1; i++ {
		if ss[i] < ss[i+1] {
			t.Fatalf("expected %v to be sorted desc", ss)
		}
	}

	for i := 0; i < len(si)-1; i++ {
		if si[i] < si[i+1] {
			t.Fatalf("expected %v to be sorted desc", si)
		}
	}

	t.Log("TestSortDesc pass")
}
