package merge_sort

import (
	"math"
)

func merge(first [][]int, second [][]int) [][]int {
	mergeSlice := make([][]int, 1)
	for len(first[0]) != 0 && len(second[0]) != 0 {
		if first[0][0] > second[0][0] {
			mergeSlice[0] = append(mergeSlice[0], second[0][0])
			second[0] = second[0][1:]
		} else {
			mergeSlice[0] = append(mergeSlice[0], first[0][0])
			first[0] = first[0][1:]
		}
	}
	mergeSlice[0] = append(mergeSlice[0], first[0]...)
	mergeSlice[0] = append(mergeSlice[0], second[0]...)
	return mergeSlice
}

/* Complexity: n*log(n) */
func MergeSort(a [][]int) [][]int {
	if len(a) == 1 {
		return a
	}
	m := int(math.Ceil(float64(len(a) / 2)))
	first := a[:m]
	second := a[m:]
	return merge(MergeSort(first), MergeSort(second))
}
