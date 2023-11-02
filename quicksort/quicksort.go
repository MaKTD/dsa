package quicksort

// TODO(rewrite properly)
// TODO(add pivot choosing by median of median principle)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSortRec(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortRec(arr, low, p-1)
		arr = quickSortRec(arr, p+1, high)
	}
	return arr
}

func QuickSort(arr []int) []int {
	return quickSortRec(arr, 0, len(arr)-1)
}
