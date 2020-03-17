package main

import (
	"github.com/AlgoAndStruct/myContainers/heap"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myInt struct {
	x int
}

func (i myInt) Priority() int {
	return i.x
}

func main() {
	var pQueue heap.MinHeap

	pQueue.Init()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	pQueue.Init()
	result := make([]int, 0)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Insert":
			scanner.Scan()
			x, _ := strconv.Atoi(scanner.Text())
			pQueue.Insert(myInt{x})
		case "ExtractMin":
			result = append(result, pQueue.ExtractMin().Priority())
		}
	}

	for _, elem := range(result) {
		fmt.Println(elem)
	}
}

/* Tree */

//type Tree struct {
//	root *TreeElement
//}
//
//type TreeElement struct {
//	parent *TreeElement
//	childLeft *TreeElement
//	childRight *TreeElement
//
//	value interface{}
//}
//
//func (t *TreeElement) print() string {
//	if t == nil {
//		return ""
//	}
//	return fmt.Sprintf("\tparent : %v\n\tchildLeft : %v\n\tchildRight : %v\n \t'%c' %v\n", t.parent, t.childLeft, t.childRight, rune(t.value.(TreeValue).ch), t.value.(TreeValue).frequency)
//}
//
//func (t *TreeElement) Priority() int {
//	return t.value.(TreeValue).Priority()
//}
//
//func (t *TreeElement) MakeParent(child_1 *TreeElement, child_2 *TreeElement) {
//	child_1.parent = t
//	child_2.parent = t
//	t.childLeft = &TreeElement{child_1.parent, child_1.childLeft, child_1.childRight, child_1.value}
//	t.childRight = &TreeElement{child_2.parent, child_2.childLeft, child_2.childRight, child_2.value}
//}
//
//
//type TreeValue struct {
//	frequency int
//	ch rune
//}
//
//func (t TreeValue) Priority() int {
//	return t.frequency
//}
//
//func (t TreeValue) isSymb() bool {
//	if t.ch != '0' {
//		return true
//	}
//	return false
//}
//
//func NewTreeElement(value interface{}) *TreeElement {
//	return &TreeElement{nil, nil, nil, value}
//}
//
///* Haffman */
//
//type Haffman struct {
//	codeStringSize int
//	codeTable map[rune][]rune
//	code []rune
//}
//
//func FillCodeTable(elem *TreeElement, code []rune, codeTable map[rune][]rune) []rune {
//	if elem.value.(TreeValue).isSymb() {
//		codeTable[elem.value.(TreeValue).ch] = append(codeTable[elem.value.(TreeValue).ch], code...)
//		code = code[:len(code) - 1]
//	}
//	if elem.childLeft != nil {
//		code = append(code, rune('0'))
//		code = FillCodeTable(elem.childLeft, code, codeTable)
//	}
//	if elem.childRight != nil {
//		code = append(code, rune('1'))
//		code = FillCodeTable(elem.childRight, code, codeTable)
//	}
//	return code
//}
//
//func HaffmanAlgorithm(str string, frequencyTable map[rune]int) *Haffman {
//	var HaffmanTree Tree
//	var freqPrQueue MinHeap
//	var code []rune
//	var H Haffman
//	var child_1 TreeElement
//
//	freqPrQueue.Init()
//	for ch, frequency := range(frequencyTable) {
//		newElem := NewTreeElement(TreeValue{frequency, ch})
//		freqPrQueue.Insert(newElem.Priority(), *newElem)
//
//	}
//	if freqPrQueue.Size() == 0 {
//		return nil
//	}
//	for i := 1; i < freqPrQueue.Size(); i++ {
//		bla := freqPrQueue.arr[i].value.(TreeElement)
//		fmt.Println(bla.print())
//	}
//	fmt.Println()
//
//	for child_1 = freqPrQueue.ExtractMin().value.(TreeElement); freqPrQueue.Size() != 0; child_1 = freqPrQueue.ExtractMin().value.(TreeElement) {
//		child_2 := freqPrQueue.ExtractMin().value.(TreeElement)
//		parent := NewTreeElement(TreeValue{child_1.Priority() + child_2.Priority(), '0'})
//		fmt.Println(child_1.print())
//		fmt.Println(child_2.print())
//		fmt.Println(parent.print())
//		parent.MakeParent(&child_1, &child_2)
//		freqPrQueue.Insert(parent.Priority(), *parent)
//	}
//	HaffmanTree.root = &child_1
//
//	// H.codeStringSize = HaffmanTree.root.Priority()
//	H.codeTable = make(map[rune][]rune)
//	FillCodeTable(HaffmanTree.root, code, H.codeTable)
//
//	for _, letter := range str {
//		H.code = append(H.code, H.codeTable[rune(letter)]...)
//	}
//	return &H
//}
//
//func main() {
//	// scanner := bufio.NewScanner(os.Stdin)
//	// scanner.Scan()
//
//	str := "beep boop"
//
//	frequencyTable := make(map[rune]int)
//	for _, ch := range(str/*scanner.Text()*/) {
//		frequencyTable[rune(ch)]++
//	}
//
//	H := HaffmanAlgorithm(str/*scanner.Text()*/, frequencyTable)
//	if H == nil {
//		return
//	}
//	H.codeStringSize = len(H.code)
//	fmt.Println(len(frequencyTable), H.codeStringSize)
//	for k, v := range H.codeTable {
//		fmt.Printf("%c: %s\n", k, string(v))
//	}
//	fmt.Println(string(H.code))
//
//}
