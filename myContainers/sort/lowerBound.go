package sort

func isAscending(a []int) bool {
	if a[0] <= a[len(a) - 1] {
		return true
	}
	return false
}

/* finds the first element that is strictly less than the given(toFind) */
/* array a must be sorted */
/* log(n) */
func LowerBound(a []int, toFind int, l, r int) int {
	if isAscending(a) {
		if a[0] < toFind {
			return 0
		} else {
			return -1
		}
	}

	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		if a[m] < toFind {
			fix = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return fix
}
