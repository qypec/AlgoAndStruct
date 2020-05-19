package binary_search

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type testBinarySearch struct {
	a      []int
	toFind int
	l, r   int
}

func TestBinarySearch(t *testing.T) {
	numberOfCases := 0

	testsValue := []testBinarySearch{
		/* 0 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 3, 0, 7}, // ascending order
		/* 1 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 3, 0, 7}, // descending order
		/* 2 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 1, 0, 7},
		/* 3 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 15, 0, 7},
		/* 4 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 20, 0, 7}, // greater than anything ascending
		/* 5 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 20, 0, 7}, // greater than anything descending
		/* 6 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 15, 0, 7},
		/* 7 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 1, 0, 7},
		/* 8 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 9},
		/* 9 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 20, 0, 9},
		/* 10 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0, 0, 9},
		/* 11 */ {[]int{1, 3, 5, 7, 9, 13, 13, 15}, 13, 0, 7},
		/* 12 */ {[]int{1, 3, 3, 3, 5, 7, 9, 13, 13, 15}, 3, 0, 9},
		/* 13 */ {[]int{1, 3, 3, 3, 5, 7, 9, 13, 13, 15}, 3, 4, 9},
		/* 14 */ {[]int{1, 3, 3, 3, 5, 7, 9, 13, 13, 15}, 3, 3, 9},
		/* 15 */ {[]int{1, 3, 3, 3, 5, 7, 9, 13, 13, 15}, 3, 2, 9},
		/* 16 */ {[]int{1, 3, 3, 3, 5, 7, 9, 13, 13, 15}, 13, 2, 7},
		/* 17 */ {[]int{1, 3, 3, 3, 5, 7, 9, 13, 13, 15}, 3, 2, 2},
	}

	testsExpected := []int{
		/* 0 */ 1,
		/* 1 */ 6,
		/* 2 */ 0,
		/* 3 */ 0,
		/* 4 */ -1,
		/* 5 */ -1,
		/* 6 */ 7,
		/* 7 */ 7,
		/* 8 */ 0,
		/* 9 */ -1,
		/* 10 */ -1,
		/* 11 */ 5,
		/* 12 */ 1,
		/* 13 */ -1,
		/* 14 */ 3,
		/* 15 */ 2,
		/* 16 */ 7,
		/* 17 */ 2,
	}

	for i := 0; i < len(testsValue); i++ {
		testCase := testsValue[i]
		actual := BinarySearch(testCase.a, testCase.toFind, testCase.l, testCase.r)
		require.Equal(t, testsExpected[i], actual, fmt.Sprintf("case %v", i))
		numberOfCases++
	}

	fmt.Println("UpperBound. Number of cases: ", numberOfCases)
}
