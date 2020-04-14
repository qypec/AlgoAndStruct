package main

import (
	"bufio"
	"fmt"
	"os"
)

func polynomialHashFunction(str string) int {
	hash := 0
	for _, ch := range str {
		hash += int(ch)
	}
	return hash
}

func SubStringRabinKarp(text, pattern string) []int {
	if len(text) < len(pattern) {
		return []int{}
	}
	/* calculated a hash for all windows of size len(pattern) */
	pLen := len(pattern)
	hash := make([]int, len(text) - pLen + 1)
	hash[0] = polynomialHashFunction(text[:pLen])
	for i := 1; i < len(text) - pLen + 1; i++ {
		hash[i] = hash[i - 1] - int(text[i - 1]) + int(text[i + pLen - 1])
	}

	patternHash := polynomialHashFunction(pattern)
	occur := make([]int, 0)
	for i, windowHash := range hash {
		if windowHash == patternHash && text[i:i + pLen] == pattern {
			occur = append(occur, i)
		}
	}
	return occur
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	pattern := scanner.Text()
	scanner.Scan()
	text := scanner.Text()

	arr := SubStringRabinKarp(text, pattern)
	for i, v := range arr {
		fmt.Printf("%v", v)
		if i != len(arr) - 1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}