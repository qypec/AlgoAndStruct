package binaryTree

/* binaryTree implementation */
type Tree struct {
	root *Element
}

type Element struct {
	parent *Element
	childLeft *Element
	childRight *Element

	Value interface{}
}

func MakeParent(childLeft *Element, childRight *Element) *Element {
	var t Element

	t.parent = nil
	t.Value = nil
	t.childLeft = childLeft
	t.childRight = childRight
	return &t
}

func (t Element) HasChilds() bool {
	return t.childLeft != nil || t.childRight != nil
}

func (t Element) HasParent() bool {
	return t.parent != nil
}

func MakeElement(value interface{}) *Element {
	var t Element

	t.parent = nil
	t.childLeft = nil
	t.childRight = nil
	t.Value = value
	return &t
}

func MakeRoot(element *Element) *Tree {
	var t Tree

	t.root = element
	return &t
}
