package recursion_basic

import "testing"

func TestSumSimple(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := Sum(s)
	if sum != 55 {
		t.Fatalf("expected %d, got %d", 55, sum)
	}
	t.Logf("TestSumSimple success")
}

func TestMaxSimple(t *testing.T) {
	s := []int{1, 5, 1, 100, 1341, 1231, 123, 512, 10, 100, 13341, 1234, 123, 343, 313, 10000, 1341, 1364}
	max := Max(s)
	if max != 13341 {
		t.Fatalf("expected %d, got %d", 13341, max)
	}
	t.Logf("TestMaxSimple sucess")
}
