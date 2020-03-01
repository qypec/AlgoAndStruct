package main

import (
	"container/list"
	"fmt"
)

func isClosingBracket(closer rune, opener rune) bool {
	if closer == ')' && opener == '(' {
		return true
	} else if closer == ']' && opener == '[' {
		return true
	} else if closer == '}' && opener == '{' {
		return true
	}
	return false
}

func isBracket(elem rune) bool {
	brackets := "()[]{}"
	for _, item := range brackets {
		if elem == item {
			return true
		}
	}
	return false
}

func main() {
	var inputCode string
	fmt.Scan(&inputCode)

	errorCounter := 0
	stopCounter := 0
	lst := list.New()
	for _, codeCh := range inputCode {
		if isBracket(codeCh) {
			errorCounter += stopCounter + 1
			stopCounter = 0
			if (lst.Len() != 0) && isClosingBracket(codeCh, lst.Back().Value.(rune)) {
				lst.Remove(lst.Back())
				errorCounter -= 2
			} else {
				lst.PushBack(codeCh)
			}
		} else {
			stopCounter++
		}
	}
	if lst.Len() == 0 {
		fmt.Println("Success")
	} else {
		fmt.Println(errorCounter)
	}
}