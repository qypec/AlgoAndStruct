package sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type countSortStruct struct {
	arr []int
	m int
}

func TestCountSort(t *testing.T) {
	numberOfCases := 0

	testsValue := []countSortStruct {
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1,}, 1}, // equal elements
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5,}, 5}, // sorted slice
		{[]int{5, 5, 4, 4, 3, 3, 2, 2, 1, 1,}, 5}, // reversed slice
		{[]int{1, 1, 5, 3, 5, 5, 2, 2, 2, 3, 2,}, 5},
		{[]int{2, 3, 9, 2, 9,}, 9},
		{[]int{2, 3, 1, 3, 4, 4, 4, 1, 2, 4, 2,}, 9},
	}
	testsExpected := [][]int {
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1,},
		{1, 1, 2, 2, 3, 3, 4, 4, 5, 5,},
		{1, 1, 2, 2, 3, 3, 4, 4, 5, 5,},
		{1, 1, 2, 2, 2, 2, 3, 3, 5, 5, 5,},
		{2, 2, 3, 9, 9,},
		{1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 4,},
	}

	for i := 0; i < len(testsValue); i++ {
		testCase := testsValue[i]

		actualResult := CountSort(testCase.arr, testCase.m)
		require.Equal(t, testsExpected[i], actualResult)
		numberOfCases++
	}

	fmt.Println("CountSort. Number of cases: ", numberOfCases)
}