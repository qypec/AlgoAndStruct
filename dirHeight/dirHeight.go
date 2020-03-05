package main

import (
	"math"
	"fmt"
)

func dirHeight(arr []int, vertex int) int {
	height := 1
	for i, elem := range arr {
		if elem == vertex {
			height = int(math.Max(float64(height), float64(1 + dirHeight(arr, i))))
		}
	}
	return height
}

func main() {
	var N, root int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
		if arr[i] == -1 {
			root = i
		}
	}
	fmt.Println(dirHeight(arr, root))
}