package quick_sort

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] <= pivot {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func recQuickSort(arr []int, l, r int) {
	if l < 0 || r < 0 {
		return
	}
	if l >= r {
		return
	}
	m := partition(arr, l, r)
	recQuickSort(arr, l, m-1)
	recQuickSort(arr, m+1, r)
}

func QuickSort(arr []int) { recQuickSort(arr, 0, len(arr)-1) }
