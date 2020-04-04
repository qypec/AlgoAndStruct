package quick_sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	numberOfCases := 0

	/* Test 01 random test */
	{
		numberOfTests := 100
		numberOfItems := 10000

		for i := 0; i < numberOfTests; i++ {
			testsValue := rand.Perm(numberOfItems)

			testsExpected := make([]int, len(testsValue))
			copy(testsExpected, testsValue)

			sort.Slice(testsExpected, func(i, j int) bool {
				return testsExpected[i] < testsExpected[j]
			})
			QuickSort(testsValue)

			require.Equal(t, testsExpected, testsValue)
			numberOfCases++
		}
	}

	testsValue := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},                             // equal elements
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},    // sorted slice
		{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, // reversed slice
		{-1, 0, 1, 0},
		{0, 0, 1, -1},
		{1, 2, 0, 0, 2, 2, 1},
		{1, 2, 0, 2, 1, 1, 1},
	}

	for i := 0; i < len(testsValue); i++ {
		testCase := testsValue[i]
		testCaseExpected := make([]int, len(testCase))
		copy(testCaseExpected, testCase)

		sort.Slice(testCaseExpected, func(i, j int) bool {
			return testCaseExpected[i] < testCaseExpected[j]
		})
		QuickSort(testCase)

		require.Equal(t, testCaseExpected, testCase)
		numberOfCases++
	}

	fmt.Println("QuickSort. Number of cases: ", numberOfCases)
}
