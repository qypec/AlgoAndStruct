package quick_sort

func partition3(arr []int, l int, r int) (int, int) {
	if r-l <= 1 { // len(arr) = 2
		if arr[r] < arr[l] {
			arr[r], arr[l] = arr[l], arr[r]
		}
		return l, r
	}

	mid := l
	pivot := arr[r]
	for mid <= r {
		if arr[mid] < pivot {
			arr[l], arr[mid] = arr[mid], arr[l]
			l++
			mid++
		} else if arr[mid] == pivot {
			mid++
		} else if arr[mid] > pivot {
			arr[mid], arr[r] = arr[r], arr[mid]
			r--
		}
	}
	return l - 1, mid
}

func recQuickSort3(arr []int, l, r int) {
	if l < 0 || r < 0 {
		return
	}
	if l >= r {
		return
	}
	m1, m2 := partition3(arr, l, r)
	recQuickSort3(arr, l, m1)
	recQuickSort3(arr, m2, r)
}

func QuickSort3(arr []int) { recQuickSort3(arr, 0, len(arr)-1) }
