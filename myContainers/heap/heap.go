package heap

func outOfRange(x, l, r int) bool {
	if x >= l && x <= r {
		return false
	}
	return true
}

/* Min-Heap implementation */
type MinHeap struct {
	arr []HeapElement
	size int
}

type HeapElement struct {
	index int
	priority int

	value interface{}
}

func (p *MinHeap) Init() {
	p.arr = make([]HeapElement, 1)
	p.arr[0] = HeapElement{0, 0,nil}
	p.size = 0
}

func (p MinHeap) Size() int { return p.size }


func (p MinHeap) Back() *HeapElement {
	if p.Size() != 0 {
		return &p.arr[p.Size()]
	}
	return nil
}

func (p MinHeap) Front() *HeapElement {
	if p.Size() != 0 {
		return &p.arr[1]
	}
	return nil
}

func (p MinHeap) getParent(child *HeapElement) *HeapElement {
	parentIndex := int(child.index / 2)
	if parentIndex != 0 {
		return &p.arr[parentIndex]
	}
	return nil
}

func (p *MinHeap) swap(child, parent **HeapElement) {
	childIndex := (*child).index
	parentIndex := (*parent).index

	p.arr[childIndex], p.arr[parentIndex] = p.arr[parentIndex], p.arr[childIndex]
	p.arr[parentIndex].index, p.arr[childIndex].index = parentIndex, childIndex

	*child = &p.arr[parentIndex]
	*parent = &p.arr[childIndex]
}

func (p *MinHeap) siftingUp() {

	child := p.Back()
	if child == nil {
		return
	}
	for parent := p.getParent(child); parent != nil; parent = p.getParent(child) {
		if child.priority < parent.priority {
			p.swap(&child, &parent)
		} else { break }
	}
}

func (p *MinHeap) Insert(priority int, value interface{}) {
	p.arr = append(p.arr, HeapElement{p.size + 1, priority, value})
	p.size++
	p.siftingUp()
}

func (p MinHeap) getChild(parent *HeapElement) *HeapElement {
	childIndexLeft, childIndexRight := int(parent.index * 2), int(parent.index * 2 + 1)
	if outOfRange(childIndexLeft, 1, p.Size()) && outOfRange(childIndexRight, 1, p.Size()) {
		return nil
	} else if outOfRange(childIndexRight, 1, p.Size()) {
		return &p.arr[childIndexLeft]
	} else {
		if p.arr[childIndexLeft].priority < p.arr[childIndexRight].priority {
			return &p.arr[childIndexLeft]
		} else {
			return &p.arr[childIndexRight]
		}
	}
}

func (p *MinHeap) siftingDown() {

	parent := p.Front()
	if parent == nil {
		return
	}
	for child := p.getChild(parent); child != nil; child = p.getChild(parent) {
		if child.priority < parent.priority {
			p.swap(&child, &parent)
		} else { break }
	}
}

func (p *MinHeap) ExtractMin() interface{} {
	if frontElem := p.Front(); frontElem == nil {
		return nil
	}
	frontElem := p.Front()
	backElem := p.Back()
	minValue := frontElem.value
	p.swap(&frontElem, &backElem)
	p.arr = p.arr[:p.Size()]
	p.size--
	p.siftingDown()
	return minValue
}

func (p *MinHeap) Reset() {
	p.arr = nil
	p.size = 0
	p.Init()
}
