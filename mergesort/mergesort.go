package mergesort

import (
	"golang.org/x/exp/constraints"
)

func mergeIntoSortedSliceAsc[T constraints.Ordered](l, r []T) []T {
	result := make([]T, len(l)+len(r))

	j := 0
	k := 0
	i := 0
	for j < len(l) && k < len(r) {
		if l[j] < r[k] {
			result[i] = l[j]
			j += 1
		} else {
			result[i] = r[k]
			k += 1
		}
		i += 1
	}
	for j < len(l) {
		result[i] = l[j]
		i += 1
		j += 1
	}
	for k < len(r) {
		result[i] = r[k]
		i += 1
		k += 1
	}

	return result
}

func merge[T any](l, r []T, compare func(a, b T) bool) []T {
	result := make([]T, len(l)+len(r))

	j := 0
	k := 0
	i := 0
	for j < len(l) && k < len(r) {
		if compare(l[j], r[k]) {
			result[i] = l[j]
			j += 1
		} else {
			result[i] = r[k]
			k += 1
		}
		i += 1
	}
	for j < len(l) {
		result[i] = l[j]
		i += 1
		j += 1
	}
	for k < len(r) {
		result[i] = r[k]
		i += 1
		k += 1
	}

	return result
}

func MergeSort[T any](s []T, compare func(a, b T) bool) []T {
	if len(s) < 2 {
		return s
	}
	left := MergeSort(s[:len(s)/2], compare)
	right := MergeSort(s[len(s)/2:], compare)
	return merge(left, right, compare)
}
