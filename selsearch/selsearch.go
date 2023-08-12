package selsearch

func SelectionSort(s []int, comp func(i, j int) bool) {
	for i := range s {
		max := i
		for j := i + 1; j < len(s); j++ {
			if comp(max, j) {
				max = j
			}
		}
		s[i], s[max] = s[max], s[i]
	}
}
