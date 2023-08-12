package lsearch

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

func createTestSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}

func TestLinearSearchSimple(t *testing.T) {
	s := createTestSlice(1000)

	cases := []TCaseSimple{
		{0, true},
		{1, true},
		{30, true},
		{60, true},
		{512, true},
		{682, true},
		{999, true},
		{1100, false},
		{1200, false},
		{10000, false},
		{5012310, false},
	}

	for _, c := range cases {
		i, ok := LinearSearch(c.e, s)
		if ok != c.ok {
			t.Fatalf("\t%s\tShould be able to find elemnt : expected ok %v got %v", failed, c.ok, ok)
		}
		if ok && c.e != s[i] {
			t.Fatalf("\t%s\tShould be able to fund element : expected index %v got %v", failed, c.e, s[i])
		}

		t.Logf("\t%s\tShould be able to find element", success)
	}
}

func BenchmarkLinearSearch1000(b *testing.B) {
	size := 1000
	s := createTestSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(rand.Intn(size), s)
	}
}

func BenchmarkLinearSearch10_000(b *testing.B) {
	size := 10_000
	s := createTestSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(rand.Intn(size), s)
	}
}

func BenchmarkLinearSearch100_000(b *testing.B) {
	size := 100_000
	s := createTestSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(rand.Intn(size), s)
	}
}
