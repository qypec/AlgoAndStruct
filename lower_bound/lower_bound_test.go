package lower_bound

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type testLowerBound struct {
	a      []int
	toFind int
}

func TestLowerBound(t *testing.T) {
	numberOfCases := 0

	testsValue := []testLowerBound{
		/* 0 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 3}, // ascending order
		/* 1 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 3}, // descending order
		/* 2 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 6},
		/* 3 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 6},
		/* 4 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 0}, // greater than anything ascending
		/* 5 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 0}, // greater than anything descending
		/* 6 */ {[]int{1, 3, 5, 7, 9, 11, 13, 15}, 14},
		/* 7 */ {[]int{15, 13, 11, 9, 7, 5, 3, 1}, 14},
		/* 8 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 14},
		/* 9 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1},
		/* 10 */ {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	testsExpected := []int{
		/* 0 */ 0,
		/* 1 */ 7,
		/* 2 */ 0,
		/* 3 */ 5,
		/* 4 */ -1,
		/* 5 */ -1,
		/* 6 */ 0,
		/* 7 */ 1,
		/* 8 */ 0,
		/* 9 */ -1,
		/* 10 */ -1,
	}

	for i := 0; i < len(testsValue); i++ {
		testCase := testsValue[i]
		actual := LowerBound(testCase.a, testCase.toFind, 0, len(testCase.a)-1)
		require.Equal(t, testsExpected[i], actual, fmt.Sprintf("case %v", i))
		numberOfCases++
	}

	fmt.Println("LowerBound. Number of cases: ", numberOfCases)

}
