package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func OutOfRange(x, l, r int) bool {
	if x >= l && x <= r {
		return false
	}
	return true
}

type Priorier interface {
	Priority() int
}

/* MinHeap */

type MinHeap struct {
	arr []HeapElement
	size int
}

type HeapElement struct {
	priority int
	index int

	value interface{}
}

func (e *HeapElement) Priority() int {
	return e.priority
}

func (p *MinHeap) Init() {
	p.arr = make([]HeapElement, 1)
	p.arr[0] = HeapElement{0, 0, nil}
	p.size = 0
}

func (p *MinHeap) Size() int { return p.size }

func (p *MinHeap) Back() *HeapElement {
	if p.Size() != 0 {
		return &p.arr[p.Size()]
	}
	return nil
}

func (p *MinHeap) Front() *HeapElement {
	if p.Size() != 0 {
		return &p.arr[1]
	}
	return nil
}

func (p MinHeap) GetParent(child *HeapElement) *HeapElement {
	parentIndex := int(child.index / 2)
	if parentIndex != 0 {
		return &p.arr[parentIndex]
	}
	return nil
}

func (p *MinHeap) siftingUp() {
	child := p.Back()
	if child == nil {
		return
	}
	for parent := p.GetParent(child); parent != nil; parent = p.GetParent(child) {
		// fmt.Println("up")
		if child.Priority() < parent.Priority() {
			p.swap(child.index, parent.index)
			// child.index, parent.index = parent.index, child.index
		} else { break }
	}
}

func (p *MinHeap) Insert(elem HeapElement) {
	elem.index = p.size + 1
	p.arr = append(p.arr, elem)
	p.size++
	p.siftingUp()
}

func (p MinHeap) GetChild(parent *HeapElement) *HeapElement {
	childIndexLeft, childIndexRight := int(parent.index * 2), int(parent.index * 2 + 1)
	if OutOfRange(childIndexLeft, 1, p.Size()) && OutOfRange(childIndexRight, 1, p.Size()) {
		return nil
	} else if OutOfRange(childIndexRight, 1, p.Size()) {
		return &p.arr[childIndexLeft]
	} else {
		if p.arr[childIndexLeft].Priority() < p.arr[childIndexRight].Priority() {
			return &p.arr[childIndexLeft]
		} else {
			return &p.arr[childIndexRight]
		}
	}
	return nil
}

func (p *MinHeap) siftingDown() {
	parent := p.Front()
	if parent == nil {
		return
	}
	for child := p.GetChild(parent); child != nil; child = p.GetChild(parent) {
		if child.Priority() < parent.Priority() {
			p.swap(child.index, parent.index)
		} else { break }
	}
}

func (p *MinHeap) swap(index_1, index_2 int) {
	if p.Size() > 1 {
		p.arr[index_1].priority, p.arr[index_2].priority = p.arr[index_2].priority, p.arr[index_1].priority
		p.arr[index_1].value, p.arr[index_2].value = p.arr[index_2].value, p.arr[index_1].value
	}
}

func (p *MinHeap) ExtractMin() *HeapElement {
	if p.Size() == 0 { return nil }
	min := *p.Front()
	p.swap(1, p.Back().index)
	p.arr = p.arr[:p.Size()]
	p.size--
	p.siftingDown()
	return &min
}

// type TreeElement struct {
// 	value interface{}
// }

// func (t *TreeElement) Priority() int {
// 	return t.value.(*TreeValue).Priority()
// }

// type TreeValue struct {
// 	frequency int
// 	ch rune
// }

// func (t *TreeValue) Priority() int {
// 	return t.frequency
// }

func main() {
	var pQueue MinHeap

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
			pQueue.Insert(HeapElement{x, 0, nil})
		case "ExtractMin":
			result = append(result, pQueue.ExtractMin().Priority())
		}
	}

	for _, elem := range(result) {
		fmt.Println(elem)
	}
}





// /* Tree */

// type TreeElement struct {
// 	parent TreeElement
// 	childLeft TreeElement
// 	childRight TreeElement
// 	value interface{}
// }

// type Tree struct {
// 	root TreeElement
// }

// func NewTreeElement(value interface{}) TreeElement {
// 	return &TreeElement{nil, nil, nil, value}
// }

// func (t *TreeElement) MakeParent(child_1 *TreeElement, child_2 *TreeElement) {
// 	child_1.parent = t
// 	child_2.parent = t
// 	t.childLeft = child_1
// 	t.childRight = child_2
// }

// func (t *TreeElement) Priority() int {
// 	return t.value.(TreeValue).frequency
// }

// /* Tree Element value */

// type TreeValue struct {
// 	frequency int
// 	ch rune
// }

// func (t *TreeValue) isSymb() bool {
// 	if t.ch != '0' {
// 		return true
// 	}
// 	return false
// }

// /* Haffman */

// type Haffman struct {
// 	codeStringSize int
// 	codeTable map[rune]string
// 	code string
// }

// func HaffmanAlgorithm(str string, frequencyTable map[rune]int) {
// 	var HaffmanTree Tree
// 	var freqPrQueue MinHeap

// 	freqPrQueue.Init()
// 	for ch, frequency := range(frequencyTable) {
// 		freqPrQueue.Insert(NewTreeElement(&TreeValue{frequency, ch}))
// 	}

// 	for child_1 := freqPrQueue.ExtractMin(); freqPrQueue.Size() != 0; child_1 := freqPrQueue.ExtractMin() {
// 		child_2 := freqPrQueue.ExtractMin()
// 		parent := NewTreeElement(&TreeElement{child_1.value.(TreeValue).frequency + child_2.value.(TreeValue).frequency, '0'})
// 		parent.MakeParent(child_1, child_2)
// 		freqPrQueue.Insert(parent)
// 	}
// 	HaffmanTree.root = freqPrQueue.ExtractMin()
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()

// 	frequencyTable := make(map[rune]int)
// 	for _, ch := range(scanner.Text()) {
// 		frequencyTable[rune(ch)]++
// 	}

// 	HaffmanAlgorithm(scanner.Text(), frequencyTable)

// 	// for k, v := range frequencyTable {
// 	// 	fmt.Printf("key: %c, value: %d\n", k, v)
// 	// }

// }
