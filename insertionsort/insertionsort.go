package insertionsort

import "golang.org/x/exp/constraints"

func SortIntAsc(s []int) {
	if len(s) < 2 {
		return
	}

	for i := 1; i < len(s); i++ {
		elem := s[i]

		for j := i - 1; j >= 0 && s[j] > elem; j-- {
			s[j+1] = s[j]
			s[j] = elem
		}
	}
}

func SortIntDesc(s []int) {
	if len(s) < 2 {
		return
	}

	for i := 1; i < len(s); i++ {
		elem := s[i]

		for j := i - 1; j >= 0 && s[j] < elem; j-- {
			s[j+1] = s[j]
			s[j] = elem
		}
	}
}

func SortAsc[K constraints.Ordered](s []K) {
	if len(s) < 2 {
		return
	}

	for i := 1; i < len(s); i++ {
		elem := s[i]

		for j := i - 1; j >= 0 && s[j] > elem; j-- {
			s[j+1] = s[j]
			s[j] = elem
		}
	}
}

func SortDesc[K constraints.Ordered](s []K) {
	if len(s) < 2 {
		return
	}

	for i := 1; i < len(s); i++ {
		elem := s[i]

		for j := i - 1; j >= 0 && s[j] < elem; j-- {
			s[j+1] = s[j]
			s[j] = elem
		}
	}
}
