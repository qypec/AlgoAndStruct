package sort

func isDescending(a []float64) bool {
	if a[0] >= a[len(a) - 1] {
		return true
	}
	return false
}

/* finds the first element that is strictly larger than the given(toFind) */
/* array a must be sorted */
/* log(n) */
func UpperBound(a []float64, toFind float64, l, r int) int {
	if isDescending(a) {
		if a[0] > toFind {
			return 0
		} else {
			return -1
		}
	}

	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		if a[m] > toFind {
			fix = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return fix
}
