package lsearch

func LinearSearch(e int, s []int) (index int, ok bool) {
	for i, v := range s {
		if v == e {
			index = i
			ok = true
			return
		}
	}
	return
}
