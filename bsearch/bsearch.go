package bsearch

func BinarySearch(e int, s []int) (int, bool) {
	r := 0
	l := len(s) - 1

	for r <= l {
		m := (r + l) / 2
		v := s[m]

		if v == e {
			return m, true
		}
		if e > v {
			r = m + 1
		} else {
			l = m - 1
		}
	}

	return 0, false
}
