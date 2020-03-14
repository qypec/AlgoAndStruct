package main

import (
	"fmt"
	"bufio"
	"os"
)

/* MinHeap */

type MinHeap struct {
	arr []interface{}
	size int
}

type arrayElement struct {
	value interface{}
	index int
}

func (p *MinHeap) Init() {
	p.arr = make([]interface{}, 0)
	p.size = 0
}

func (p *MinHeap) Size() int { return p.size }

func (p MinHeap) Back() *arrayElement {
	if p.Size() == 0 {
		return nil
	}
	return &arrayElement{p.arr[p.Size() - 1], p.Size() - 1}
}

func (p MinHeap) GetParent(child *arrayElement) *arrayElement {
	// if child.index < 1 || child.index > p.Len() {
	// 	panic(fmt.Sprintf("p.arr : child.index - %v out of range", child.index))
	// }
	// if child.value != p.arr[child.index] {
	// 	panic(fmt.Sprintf("p.arr[child.index] != child.value\n%v != %v", p.arr[child.index], child.value))
	// }
	parentIndex := int(child.index / 2)
	if parentIndex != 0 {
		return &arrayElement{p.arr[parentIndex], parentIndex}
	}
	return nil
}

func (p *MinHeap) siftingUp() {
	child := p.Back()
	if child == nil {
		return
	}
	for parent := p.GetParent(child); parent != nil; parent = p.GetParent(child) {
		if child.value.Priority() > parent.value.Priority() {
			swap(&p.arr[child.index], &p.arr[parent.index])
			swap(&child.index, &parent.index)
		} else { break }
	}
}

func (p *MinHeap) Insert(x interface{}) {
	p.arr = append(p.arr, x)
	p.size++
	p.siftingUp()
}

/* Tree */

type TreeElement struct {
	parent TreeElement
	childLeft TreeElement
	childRight TreeElement
	value interface{}
}

type Tree struct {
	root TreeElement
}

func NewTreeElement(value interface{}) TreeElement {
	return &TreeElement{nil, nil, nil, value}
}

func (t *TreeElement) MakeParent(child_1 *TreeElement, child_2 *TreeElement) {
	child_1.parent = t
	child_2.parent = t
	t.childLeft = child_1
	t.childRight = child_2
}

func (t *TreeElement) Priority() int {
	return t.value.(TreeValue).frequency
}

/* Tree Element value */

type TreeValue struct {
	frequency int
	ch rune
}

func (t *TreeValue) isSymb() bool {
	if t.ch != '0' {
		return true
	}
	return false
}

/* Haffman */

type Haffman struct {
	codeStringSize int
	codeTable map[rune]string
	code string
}

func HaffmanAlgorithm(str string, frequencyTable map[rune]int) {
	var HaffmanTree Tree
	var freqPrQueue MinHeap

	freqPrQueue.Init()
	for ch, frequency := range(frequencyTable) {
		freqPrQueue.Insert(NewTreeElement(&TreeValue{frequency, ch}))
	}

	for child_1 := freqPrQueue.ExtractMin(); freqPrQueue.Size() != 0; child_1 := freqPrQueue.ExtractMin() {
		child_2 := freqPrQueue.ExtractMin()
		parent := NewTreeElement(&TreeElement{child_1.value.(TreeValue).frequency + child_2.value.(TreeValue).frequency, '0'})
		parent.MakeParent(child_1, child_2)
		freqPrQueue.Insert(parent)
	}
	HaffmanTree.root = freqPrQueue.ExtractMin()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	frequencyTable := make(map[rune]int)
	for _, ch := range(scanner.Text()) {
		frequencyTable[rune(ch)]++
	}

	HaffmanAlgorithm(scanner.Text(), frequencyTable)

	// for k, v := range frequencyTable {
	// 	fmt.Printf("key: %c, value: %d\n", k, v)
	// }

}

/*
package main

import "fmt"


type Priorier interface {
	Priority() int
}

type HeapElem struct {
	p int
}

type Heap struct {
	root HeapElem
}

func (h Heap) Priority() int {
	return h.root.p
}

func main() {
	var h Heap
	
	h.root = HeapElem{4}
	var pr Priorier = h
	fmt.Println(pr.Priority())
}


type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
*/