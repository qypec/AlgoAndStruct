/*
	Дано целое число n, и массив A[1..n], содержащий неотрицательные целые числа.
	Найдите наибольшую невозрастающую подпоследовательность в A. В первой строке выведите её длину k,
	во второй — её индексы 1 <= i1 < i2 < ... <= ik (таким образом, A[i1] >= A[i2] >= ... >= A[in]).
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isAscending(a []int) bool {
	if a[0] <= a[len(a)-1] {
		return true
	}
	return false
}

/* finds the first element that is strictly less than the given(toFind) */
/* array a must be sorted */
/* log(n) */
func lowerBound(a []int, toFind int, l, r int) int {
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

/* nlog(n) */
func lds(arr []int, n int) []int {
	d := make([]int, n+1) // d[i] is the last element of a subsequence of length i
	d[0] = math.MaxInt32
	for i := 1; i < n+1; i++ {
		d[i] = math.MinInt32
	}
	lengths := make([]int, n) // lengths[i] is the maximum length of a subsequence with the last element a[i]
	prev := make([]int, n+1)  // prev[i] is the index d[i] in the array a

	subSequenceSize := 0
	for i := 0; i < n; i++ { // filling lengths, d, prev
		firstSuitableIndex := lowerBound(d, arr[i], 0, len(d)-1)
		if d[firstSuitableIndex-1] >= arr[i] && d[firstSuitableIndex] <= arr[i] {
			if d[firstSuitableIndex] == math.MinInt32 {
				lengths[i] = firstSuitableIndex
				subSequenceSize = firstSuitableIndex
			} else {
				lengths[i] = lengths[prev[firstSuitableIndex]]
			}
			d[firstSuitableIndex] = arr[i]
			prev[firstSuitableIndex] = i
		}
	}

	subSequence := make([]int, subSequenceSize)
	subIndex := subSequenceSize - 1
	prevElem := subSequenceSize
	for i := n - 1; i >= 0 && prevElem != 0; i-- { // subsequence recovery
		if lengths[i] == prevElem {
			subSequence[subIndex] = i + 1
			prevElem--
			subIndex--
		}
	}
	return subSequence
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	/* scanning n */
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	/* scanning array */
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, x)
	}

	subSequence := lds(arr, n)
	fmt.Println(len(subSequence))
	for _, elem := range subSequence {
		fmt.Print(elem, " ")
	}
}
