package upper_bound

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type testUppedBound struct {
	a      []int
	toFind int
}

func TestUpperBound(t *testing.T) {
	numberOfCases := 0

	testsValue := []testUppedBound{
		/* 0 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 3}, // ascending order
		/* 1 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 3}, // descending order
		/* 2 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 6},
		/* 3 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 6},
		/* 4 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 20}, // greater than anything ascending
		/* 5 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 20}, // greater than anything descending
		/* 6 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 14},
		/* 7 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 14},
		/* 8 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 14},
		/* 9 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1},
		/* 10 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	testsExpected := []int{
		/* 0 */ 2,
		/* 1 */ 0,
		/* 2 */ 3,
		/* 3 */ 0,
		/* 4 */ -1,
		/* 5 */ -1,
		/* 6 */ 7,
		/* 7 */ 0,
		/* 8 */ -1,
		/* 9 */ -1,
		/* 10 */ 0,
	}

	for i := 0; i < len(testsValue); i++ {
		testCase := testsValue[i]
		actual := UpperBound(testCase.a, testCase.toFind, 0, len(testCase.a)-1)
		require.Equal(t, testsExpected[i], actual, fmt.Sprintf("case %v", i))
		numberOfCases++
	}

	fmt.Println("UpperBound. Number of cases: ", numberOfCases)

}
