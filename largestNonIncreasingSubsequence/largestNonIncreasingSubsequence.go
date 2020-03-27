/*
	Дано целое число n, и массив A[1..n], содержащий неотрицательные целые числа.
	Найдите наибольшую невозрастающую подпоследовательность в A. В первой строке выведите её длину k,
	во второй — её индексы 1 <= i1 < i2 < ... <= ik (таким образом, A[i1] >= A[i2] >= ... >= A[in]).
 */

package main

import (
	"github.com/AlgoAndStruct/myContainers/sort"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/* nlog(n) */
func LnIS(arr []float64, n int) []int {
	d := make([]float64, n + 1) // d[i] is the last element of a subsequence of length i
	d[0] = math.Inf(1)
	for i := 1; i < n + 1; i++ {
		d[i] = math.Inf(-1)
	}
	lengths := make([]int, n) // lengths[i] is the maximum length of a subsequence with the last element a[i]
	prev := make([]int, n + 1) // prev[i] is the index d[i] in the array a

	subSequenceSize := 0
	for i := 0; i < n; i++ { // filling lengths, d, prev
		firstSuitableIndex := sort.LowerBound(d, arr[i], 0, len(d) - 1)
		if d[firstSuitableIndex - 1] >= arr[i] && d[firstSuitableIndex] <= arr[i] {
			if d[firstSuitableIndex] == math.Inf(-1) {
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
	arr := make([]float64, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, float64(x))
	}

	subSequence := LnIS(arr, n)
	fmt.Println(len(subSequence))
	for _, elem := range subSequence {
		fmt.Print(elem, " ")
	}
}
