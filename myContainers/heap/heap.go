package heap

func outOfRange(x, l, r int) bool {
	if x >= l && x <= r {
		return false
	}
	return true
}

type HeapModel interface {
	Priority() int
}

/* Min-Heap implementation */
type MinHeap struct {
	arr []HeapElement
	size int
}

type HeapElement struct {
	index int
	value HeapModel
}

func (h HeapElement) priority() int {
	return h.value.Priority()
}

func (p *MinHeap) Init() {
	p.arr = make([]HeapElement, 1)
	p.arr[0] = HeapElement{0, nil}
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

func (p MinHeap) getParent(child *HeapElement) *HeapElement {
	parentIndex := int(child.index / 2)
	if parentIndex != 0 {
		return &p.arr[parentIndex]
	}
	return nil
}

func (p *MinHeap) swap(childIndex, parentIndex int) {
	p.arr[childIndex], p.arr[parentIndex] = p.arr[parentIndex], p.arr[childIndex]
	p.arr[childIndex].index, p.arr[parentIndex].index = p.arr[parentIndex].index, p.arr[childIndex].index
}

func (p *MinHeap) siftingUp() {
	child := p.Back()
	if child == nil {
		return
	}
	for parent := p.getParent(child); parent != nil; parent = p.getParent(child) {
		if child.priority() < parent.priority() {
			p.swap(child.index, parent.index)
			child = parent
		} else { break }
	}
}

func (p *MinHeap) Insert(value HeapModel) {
	elem := HeapElement{p.size + 1, value}
	p.arr = append(p.arr, elem)
	p.size++
	p.siftingUp()
}

func (p MinHeap) getChild(parent *HeapElement) *HeapElement {
	childIndexLeft, childIndexRight := int(parent.index*2), int(parent.index*2+1)
	if outOfRange(childIndexLeft, 1, p.Size()) && outOfRange(childIndexRight, 1, p.Size()) {
		return nil
	} else if outOfRange(childIndexRight, 1, p.Size()) {
		return &p.arr[childIndexLeft]
	} else {
		if p.arr[childIndexLeft].priority() < p.arr[childIndexRight].priority() {
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
		if child.priority() < parent.priority() {
			p.swap(child.index, parent.index)
			parent = child
		} else { break }
	}
}

func (p *MinHeap) ExtractMin() HeapModel {
	if frontElem := p.Front(); frontElem == nil {
		return nil
	}
	min := p.Front().value
	p.swap(1, p.Back().index)
	p.arr = p.arr[:p.Size()]
	p.size--
	p.siftingDown()
	return min
}
