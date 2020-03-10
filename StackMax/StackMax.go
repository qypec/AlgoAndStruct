/*
	Формат входа. Первая строка содержит число запросов q. 
		Каждая из последующих q строк задаёт запрос в одном из следующих форматов : push v, pop, max.

	Формат выхода. Для каждого запроса max выведите (в отдельной строке) текущий максимум на стеке.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type StackValue struct {
	value int
	max int
}

type MyStack struct {
	arr []StackValue
	size int
}

func (s *MyStack) Init(N int) {
	s.arr = make([]StackValue, N)
	s.size = 0
}

func (s *MyStack) PushBack(elem int) {
	size := s.size + 1
	if s.GetMax() < elem {
		s.arr[size - 1] = StackValue{elem, elem}
	} else {
		s.arr[size - 1] = StackValue{elem, s.GetMax()}
	}
	s.size++
}

func (s *MyStack) PopBack() {
	if s.size != 0 {
		s.arr[s.size - 1] = StackValue{0, 0}
		s.size--
	}
}

func (s *MyStack) GetMax() int {
	if s.size != 0 {
		return s.arr[s.size - 1].max
	}
	return 0
}

func main() {
	var N, elem int
	var stack MyStack

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())

	result := make([]int, 0)
	stack.Init(N)
	for scanner.Scan() {
		switch scanner.Text() {
		case "pop":
			stack.PopBack()
		case "push":
			scanner.Scan()
			elem, _ = strconv.Atoi(scanner.Text())
			stack.PushBack(elem)
		case "max":
			result = append(result, stack.GetMax())
		}
	}

	for _, elem := range result {
		fmt.Println(elem)
	}
}
