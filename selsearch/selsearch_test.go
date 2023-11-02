package selsearch

import (
	"math/rand"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

type TCaseSimple struct {
	e  int
	ok bool
}

func TestSelectionSortSimpleAsc(t *testing.T) {
	size := 100
	s := rand.Perm(size)

	SelectionSort(s, func(i, j int) bool { return s[i] >= s[j] })

	t.Log(s)

	for i, v := range s {
		if i != v {
			t.Fatalf("\t%s\tShould be able to sort elements : expected i = v, got i = %v, v = %v", failed, i, v)
		}
	}
	t.Logf("\t%s\tShould be able to sort elements", success)
}

func TestSelectionSortSimpleDesc(t *testing.T) {
	size := 100
	s := rand.Perm(size)

	SelectionSort(s, func(i, j int) bool { return s[i] <= s[j] })

	t.Log(s)

	for i, v := range s {
		if i != size-1-v {
			t.Fatalf("\t%s\tShould be able to sort elements : expected i = v, got i = %v, v = %v", failed, i, v)
		}
	}
	t.Logf("\t%s\tShould be able to sort elements", success)
}
