package bsearch

import (
	"math/rand"
	"sort"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func createSortedArr(size int) []int {
	s := rand.Perm(size)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

type TCaseSimple struct {
	i  int
	e  int
	ok bool
}

func TestSimple(t *testing.T) {
	s := createSortedArr(1000)

	cases := []TCaseSimple{
		{i: 0, e: s[0], ok: true},
		{i: len(s) - 1, e: s[len(s)-1], ok: true},
		{i: 10, e: s[10], ok: true},
		{i: 210, e: s[210], ok: true},
		{i: 521, e: s[521], ok: true},
		{i: 891, e: s[891], ok: true},
		{i: 0, e: 1511, ok: false},
		{i: 0, e: 100000, ok: false},
		{i: 0, e: -1, ok: false},
		{i: 0, e: -12341, ok: false},
	}

	for _, c := range cases {
		i, ok := BinarySearch(c.e, s)
		if i != c.i {
			t.Fatalf("\t%s\tShould be able to fund element : expected index %v got %v", failed, c.i, i)
		}
		if ok != c.ok {
			t.Fatalf("\t%s\tShould be able to find elemnt : expected ok %v got %v, i = %v, e = %v", failed, c.ok, ok, c.i, c.e)
		}
		t.Logf("\t%s\tShould be able to find element", success)
	}
}

func BenchmarkBinarySearch1000(b *testing.B) {
	size := 1000
	s := createSortedArr(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(rand.Intn(size), s)
	}
}

func BenchmarkBinarySearch10_000(b *testing.B) {
	size := 10_000
	s := createSortedArr(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(rand.Intn(size), s)
	}
}

func BenchmarkBinarySearch100_000(b *testing.B) {
	size := 100_000
	s := createSortedArr(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(rand.Intn(size), s)
	}
}
