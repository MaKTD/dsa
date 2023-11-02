package recursion_basic

func Sum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}
	return s[0] + Sum(s[1:])
}

func Max(s []int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}

	guess := s[0]
	compare := Max(s[1:])
	if guess >= compare {
		return guess
	} else {
		return compare
	}
}
