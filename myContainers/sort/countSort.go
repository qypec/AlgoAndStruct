package sort

/* CountSort sorts arrays of integer [1...m] */
func CountSort(a []int, m int) []int {
	n := len(a) - 1
	countArray := make([]int, m + 1)
	result := make([]int, n + 1)

	/* filling countArray */
	for _, value := range a {
		countArray[value]++
	}

	/* summing countArray */
	prevIndex := 1
	for i := 2; i <= m; i++ {
		if countArray[i] != 0 {
			countArray[i] += countArray[prevIndex]
			prevIndex++
		}
	}

	/* filling result array */
	for i := n; i >= 0; i-- {
		result[countArray[a[i]] - 1] = a[i]
		countArray[a[i]]--
	}
	return result
}
